﻿using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Dapper;
using MySqlConnector;
using OTHub.BackendSync.Database.Models;
using OTHub.BackendSync.Logging;
using OTHub.Settings;

namespace OTHub.BackendSync
{
    public class TaskController
    {
        private readonly Source _source;
        private readonly ConcurrentBag<TaskControllerItem> _items = new ConcurrentBag<TaskControllerItem>();

        private class TaskControllerItem
        {
            private readonly BlockchainType _blockchain;
            private readonly BlockchainNetwork _network;
            private readonly Source _source;
            private readonly TaskRunBase _task;
            private readonly TimeSpan _runEveryTimeSpan;
            private readonly bool _startChildrenConcurrently;
            private DateTime _lastRunDateTime;
            private SystemStatus _systemStatus;

            internal TaskControllerItem(BlockchainType blockchain, BlockchainNetwork network, Source source,
                TaskRunBlockchain task, TimeSpan runEveryTimeSpan, bool startNow, int blockchainID, bool startChildrenConcurrently)
            {
                _blockchain = blockchain;
                _network = network;
                _source = source;
                _task = task;
                _runEveryTimeSpan = runEveryTimeSpan;
                _startChildrenConcurrently = startChildrenConcurrently;
                _lastRunDateTime = startNow ? DateTime.MinValue : DateTime.Now;
                _systemStatus = new SystemStatus(task.Name, blockchainID);

                using (var connection = new MySqlConnection(OTHubSettings.Instance.MariaDB.ConnectionString))
                {
                    _systemStatus.InsertOrUpdate(connection, null, NextRunDate, false, _task.ParentName).GetAwaiter().GetResult();
                }

                task.BlockchainStartup(blockchainID, blockchain, network);
            }

            internal TaskControllerItem(Source source,
                TaskRunGeneric task, TimeSpan runEveryTimeSpan, bool startNow)
            {
                _source = source;
                _task = task;
                _runEveryTimeSpan = runEveryTimeSpan;
                _lastRunDateTime = startNow ? DateTime.MinValue : DateTime.Now;
                _systemStatus = new SystemStatus(task.Name);

                using (var connection = new MySqlConnection(OTHubSettings.Instance.MariaDB.ConnectionString))
                {
                    _systemStatus.InsertOrUpdate(connection, null, NextRunDate, false, _task.ParentName).GetAwaiter().GetResult();
                }
            }

            public bool StartChildrenConcurrently => _startChildrenConcurrently;

            public bool NeedsRunning
            {
                get { return ((DateTime.Now - _lastRunDateTime) > _runEveryTimeSpan); }
            }

            public DateTime? NextRunDate
            {
                get
                {
                    if (_lastRunDateTime == DateTime.MinValue)
                        return null;

                    return _lastRunDateTime + _runEveryTimeSpan;
                }
            }

            public async Task Execute()
            {
                DateTime startTime = DateTime.Now;

                bool success = false;

                try
                {
                    if (_systemStatus.BlockchainID.HasValue)
                    {
                        Logger.WriteLine(_source, "Starting " + _task.Name + " on " + _blockchain + " " + _network);
                    }
                    else
                    {
                        Logger.WriteLine(_source, "Starting " + _task.Name);
                    }

                    await using (var connection = new MySqlConnection(OTHubSettings.Instance.MariaDB.ConnectionString))
                    {
                        await _systemStatus.InsertOrUpdate(connection, true, null, true, _task.ParentName);
                    }

                    if (_task is TaskRunBlockchain taskB)
                    {
                        success = await taskB.Execute(_source, _blockchain, _network);
                    }
                    else if (_task is TaskRunGeneric taskG)
                    {
                        await taskG.Execute(_source);
                        success = true;
                    }
                    else
                    {
                        throw new NotImplementedException("Task of type " + _task.GetType().FullName +
                                                          " is not supported.");
                    }
                }
                catch (Exception ex)
                {
                    Logger.WriteLine(_source, ex.ToString());
                }
                finally
                {
                    _lastRunDateTime = DateTime.Now;
                    await using (var connection = new MySqlConnection(OTHubSettings.Instance.MariaDB.ConnectionString))
                    {
                        await _systemStatus.InsertOrUpdate(connection, success, NextRunDate, false, _task.ParentName);
                    }
                    Logger.WriteLine(_source, "Finished " + _task.Name + " in " + (DateTime.Now - startTime).TotalSeconds + " seconds on " + _blockchain + " " + _network);
                }
            }
        }

        public void Schedule(TaskRunGeneric task, TimeSpan runEveryTimeSpan, bool startNow)
        {
            var item = new TaskControllerItem(_source, task, runEveryTimeSpan, startNow);
            _items.Add(item);
        }

        public void Schedule(TaskRunBlockchain task, bool startNow, bool startChildrenConcurrently)
        {
            using (var connection = new MySqlConnection(OTHubSettings.Instance.MariaDB.ConnectionString))
            {
                var blockchains = connection.Query(@"SELECT * FROM blockchains WHERE enabled = 1").ToArray();

                foreach (var blockchain in blockchains)
                {
                    int id = blockchain.ID;
                    string blockchainName = blockchain.BlockchainName;
                    string networkName = blockchain.NetworkName;

                    Logger.WriteLine(Source.BlockchainSync,
                        $"Setting up processes for {blockchainName} {networkName}. Task is " + task.Name);

                    BlockchainType blockchainEnum = Enum.Parse<BlockchainType>(blockchainName);
                    BlockchainNetwork networkNameEnum = Enum.Parse<BlockchainNetwork>(networkName);

                    TimeSpan interval = task.GetExecutingInterval(blockchainEnum);

                    var item = new TaskControllerItem(blockchainEnum, networkNameEnum, _source, task, interval, startNow, id, startChildrenConcurrently);
                    _items.Add(item);
                }
            }
        }

        private bool _showSleepingLogMessage = true;
        private bool isFirstSync = true;

        public TaskController(Source source)
        {
            _source = source;
        }

        public async Task Start()
        {
            while (true)
            {
                if (isFirstSync)
                {
                    isFirstSync = false;
                }

                var items = _items.Where(i => i.NeedsRunning).Reverse().ToArray();

                TaskControllerItem[] concurrentItems = items.Where(i => i.StartChildrenConcurrently).ToArray();
                TaskControllerItem[] sequentialItems = items.Except(concurrentItems).ToArray();

                if (concurrentItems.Any())
                {
                    List<Task> tasks = new List<Task>();
                    foreach (var taskControllerItem in concurrentItems)
                    {
                        tasks.Add(Task.Run(async () =>
                        {
                            await taskControllerItem.Execute();
                        }));
                    }

                    await Task.WhenAll(tasks);
                }

                if (sequentialItems.Any())
                {
                    foreach (var taskControllerItem in sequentialItems)
                    {
                        await taskControllerItem.Execute();
                    }
                }

                if (!items.Any())
                {
                    if (_showSleepingLogMessage)
                    {
                        _showSleepingLogMessage = false;
                        Logger.WriteLine(_source, "Sleeping...");
                    }

                    await Task.Delay(2000);
                }
                else
                {
                    _showSleepingLogMessage = true;
                }
            }
        }
    }
}