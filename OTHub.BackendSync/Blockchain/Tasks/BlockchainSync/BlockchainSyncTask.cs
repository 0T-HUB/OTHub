﻿using System;
using System.Threading.Tasks;
using MySqlConnector;
using OTHub.BackendSync.Blockchain.Tasks.BlockchainSync.Children;
using OTHub.BackendSync.Logging;
using OTHub.Settings;

namespace OTHub.BackendSync.Blockchain.Tasks.BlockchainSync
{
    public class BlockchainSyncTask : TaskRunBlockchain
    {
        public BlockchainSyncTask() : base("Blockchain Sync")
        {
            Add(new SyncProfileContractTask());
            Add(new SyncHoldingContractTask());
            Add(new SyncLitigationContractTask());
            Add(new SyncReplacementContractTask());
            Add(new LoadProfileBalancesTask());
            Add(new ProcessJobsTask());
        }

        public override TimeSpan GetExecutingInterval(BlockchainType type)
        {
            if (type == BlockchainType.xDai)
            {
                return TimeSpan.FromSeconds(50);
            }

            return TimeSpan.FromMinutes(5);
        }

        public override async Task Execute(Source source, BlockchainType blockchain, BlockchainNetwork network)
        {
            await using (var connection = new MySqlConnection(OTHubSettings.Instance.MariaDB.ConnectionString))
            {
                int blockchainID = GetBlockchainID(connection, blockchain, network);

                await RunChildren(source, blockchain, network, blockchainID);
            }
        }
    }
}