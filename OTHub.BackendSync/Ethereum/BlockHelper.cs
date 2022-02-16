﻿using System;
using System.Threading.Tasks;
using MySql.Data.MySqlClient;
using Nethereum.Hex.HexTypes;
using Nethereum.Web3;
using OTHub.BackendSync.Database.Models;
using OTHub.Settings.Helpers;

namespace OTHub.BackendSync.Ethereum
{
    public static class BlockHelper
    {
        private static readonly object _getEthBlockLock = new object();
        public static async Task<EthBlock> GetEthBlock(MySqlConnection connection, string blockHash, HexBigInteger blockNumber, Web3 cl)
        {
            var block = EthBlock.GetByNumber(connection, (UInt64)blockNumber.Value);

            if (block == null)
            {
                bool madeApiCall = false;

                lock (_getEthBlockLock)
                {
                    block = EthBlock.GetByNumber(connection, (UInt64)blockNumber.Value);

                    if (block == null)
                    {
                        var apiBlock = cl.Eth.Blocks.GetBlockWithTransactionsHashesByNumber.SendRequestAsync(blockNumber).GetAwaiter().GetResult();
                        madeApiCall = true;

                        block = new EthBlock
                        {
                            BlockHash = blockHash,
                            BlockNumber = (UInt64)blockNumber.Value,
                            Timestamp = TimestampHelper.UnixTimeStampToDateTime((double)apiBlock.Timestamp.Value)
                        };

                        EthBlock.Insert(connection, block);
                    }
                }

                if (madeApiCall)
                {
                    await Task.Delay(200);
                }
            }

            return block;
        }
    }
}
