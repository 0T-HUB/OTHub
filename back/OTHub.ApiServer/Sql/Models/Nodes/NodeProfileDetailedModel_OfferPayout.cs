﻿using System;

namespace OTHub.APIServer.Sql.Models.Nodes
{
    public class NodeProfileDetailedModel_OfferPayout
    {
        public string Amount { get; set; }
        public DateTime Timestamp { get; set; }
        public String OfferId { get; set; }
        public String TransactionHash { get; set; }
        public UInt64 GasUsed { get; set; }
        public UInt64 GasPrice { get; set; }
        public string GasTicker { get; set; }
    }
}