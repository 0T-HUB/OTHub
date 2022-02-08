﻿using System;
using System.Collections.Generic;
using System.Text;
using Dapper;
using MySql.Data.MySqlClient;

namespace OTHelperNetStandard.Models.Database
{
    public class OTContract_Litigation_LitigationInitiated
    {
        public string TransactionHash { get; set; }
        public UInt64 BlockNumber { get; set; }
        public DateTime Timestamp { get; set; }
        public String OfferId { get; set; }
        public String HolderIdentity { get; set; }
        public UInt64 RequestedDataIndex { get; set; }
        public ulong GasPrice { get; set; }
        public ulong GasUsed { get; set; }

        public static void InsertIfNotExist(MySqlConnection connection, OTContract_Litigation_LitigationInitiated model)
        {
            var count = connection.QueryFirstOrDefault<Int32>("SELECT COUNT(*) FROM OTContract_Litigation_LitigationInitiated WHERE TransactionHash = @hash", new
            {
                hash = model.TransactionHash
            });

            if (count == 0)
            {
                connection.Execute(
                    @"INSERT INTO OTContract_Litigation_LitigationInitiated
(TransactionHash, BlockNumber, Timestamp, OfferId, HolderIdentity, RequestedDataIndex, GasPrice, GasUsed)
VALUES(@TransactionHash, @BlockNumber, @Timestamp, @OfferId, @HolderIdentity, @RequestedDataIndex, @GasPrice, @GasUsed)",
                    new
                    {
                        model.TransactionHash,
                        model.BlockNumber,
                        model.Timestamp,
                        model.OfferId,
                        model.HolderIdentity,
                        model.RequestedDataIndex,
                        model.GasPrice,
                        model.GasUsed
                    });

                OTOfferHolder.UpdateLitigationStatusesForOffer(connection, model.OfferId);
            }
        }
    }
}