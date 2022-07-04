import { browser } from '$app/env'

export const getBadge = async () => await (await fetch(`/api/badge`)).json()

export const getDataHolders = async ({ activePageNumber = '1', numberOfItemsPerPage = '25', sortKey = 'WonOffersLast7Days', sortOrder = 'DESC', nodeIDSearchQuery = undefined, nameSearchQuery = undefined }) => {
  if (browser) {
    const request = await fetch(`/api/nodes/dataholders?ercVersion=1&_sort=${sortKey}&_order=${sortOrder}&_page=${activePageNumber}&_limit=${numberOfItemsPerPage}${ nodeIDSearchQuery ? '&NodeId_like=' + nodeIDSearchQuery : '' }${ nameSearchQuery ? '&DisplayName_like=' + nameSearchQuery : '' }`)
    const response = await request.json()
    return response
  }
}

export const getDataCreators = async ({ activePageNumber = '1', numberOfItemsPerPage = '25', sortKey = 'OffersLast7Days', sortOrder = 'DESC', nodeIDSearchQuery = undefined, nameSearchQuery = undefined }) => {
  if (browser) {
    const request = await fetch(`/api/nodes/datacreators?ercVersion=1&_sort=${sortKey}&_order=${sortOrder}&_page=${activePageNumber}&_limit=${numberOfItemsPerPage}${ nodeIDSearchQuery ? '&NodeId_like=' + nodeIDSearchQuery : '' }${ nameSearchQuery ? '&DisplayName_like=' + nameSearchQuery : '' }`)
    const response = await request.json()
    return response
  }
}

export const getDataHolderJobsWithDataHolderSettings = async ({ activePageNumber = '1', numberOfItemsPerPage = '10', sortKey = 'FinalizedTimestamp', sortOrder = 'DESC', offerIDSearchQuery = undefined, dataHolderID }) => {
  if (browser) {
    const request = await fetch(`/api/nodes/dataholder/${dataHolderID}/jobs?_sort=${sortKey}&_order=${sortOrder}&_page=${activePageNumber}&_limit=${numberOfItemsPerPage}${ offerIDSearchQuery ? '&OfferId_like=' + offerIDSearchQuery : '' }`)
    const response = await request.json()
    return response
  }
}

export const getDataHolderLitigationsWithLitigationSettings = async ({ activePageNumber = '1', numberOfItemsPerPage = '10', sortKey = 'Timestamp', sortOrder = 'DESC', offerIDSearchQuery = undefined, dataHolderID }) => {
  if (browser) {
    const request = await fetch(`/api/nodes/dataholder/${dataHolderID}/litigations?_sort=${sortKey}&_order=${sortOrder}&_page=${activePageNumber}&_limit=${numberOfItemsPerPage}${ offerIDSearchQuery ? '&OfferId_like=' + offerIDSearchQuery : '' }`)
    const response = await request.json()
    return response
  }
}

export const getDataHolderPayoutsWithPayoutSettings = async ({ activePageNumber = '1', numberOfItemsPerPage = '10', sortKey = 'Timestamp', sortOrder = 'DESC', offerIDSearchQuery = undefined, transactionHashSearchQuery = undefined, dataHolderID }) => {
  if (browser) {
    const request = await fetch(`/api/nodes/dataholder/${dataHolderID}/payouts?_sort=${sortKey}&_order=${sortOrder}&_page=${activePageNumber}&_limit=${numberOfItemsPerPage}${ offerIDSearchQuery ? '&OfferId_like=' + offerIDSearchQuery : '' }${ transactionHashSearchQuery ? '&TransactionHash_like=' + transactionHashSearchQuery : ''}`)
    const response = await request.json()
    return response
  }
}

export const getTransfersWithTransferSettings = async ({ activePageNumber = '1', numberOfItemsPerPage = '10', sortKey = 'Timestamp', sortOrder = 'DESC', transactionHashSearchQuery = undefined, dataHolderID }) => {
  if (browser) {
    const request = await fetch(`/api/nodes/dataholder/${dataHolderID}/profiletransfers?_sort=${sortKey}&_order=${sortOrder}&_page=${activePageNumber}&_limit=${numberOfItemsPerPage}${ transactionHashSearchQuery ? '&TransactionHash_like=' + transactionHashSearchQuery : ''}`)
    const response = await request.json()
    return response
  }
}


export const getDataCreatorJobsWithDataCreatorSettings = async ({ activePageNumber = '1', numberOfItemsPerPage = '10', sortKey = 'FinalizedTimestamp', sortOrder = 'DESC', offerIDSearchQuery = undefined, dataCreatorID }) => {
  if (browser) {
    const request = await fetch(`/api/nodes/datacreator/${dataCreatorID}/jobs?_sort=${sortKey}&_order=${sortOrder}&_page=${activePageNumber}&_limit=${numberOfItemsPerPage}${ offerIDSearchQuery ? '&OfferId_like=' + offerIDSearchQuery : '' }`)
    const response = await request.json()
    return response
  }
}

export const getDataCreatorLitigationsWithLitigationSettings = async ({ activePageNumber = '1', numberOfItemsPerPage = '10', sortKey = 'Timestamp', sortOrder = 'DESC', offerIDSearchQuery = undefined, dataHolderIDSearchQuery = undefined, dataCreatorID }) => {
  if (browser) {
    const request = await fetch(`/api/nodes/datacreator/${dataCreatorID}/litigations?_sort=${sortKey}&_order=${sortOrder}&_page=${activePageNumber}&_limit=${numberOfItemsPerPage}${ offerIDSearchQuery ? '&OfferId_like=' + offerIDSearchQuery : '' }${ dataHolderIDSearchQuery ? '&NodeId_like=' + dataHolderIDSearchQuery : '' }`)
    const response = await request.json()
    return response
  }
}

export const getDataCreatorTransfersWithTransferSettings = async ({ activePageNumber = '1', numberOfItemsPerPage = '10', sortKey = 'Timestamp', sortOrder = 'DESC', transactionHashSearchQuery = undefined, dataCreatorID }) => {
  if (browser) {
    const request = await fetch(`/api/nodes/datacreator/${dataCreatorID}/profiletransfers?_sort=${sortKey}&_order=${sortOrder}&_page=${activePageNumber}&_limit=${numberOfItemsPerPage}${ transactionHashSearchQuery ? '&TransactionHash_like=' + transactionHashSearchQuery : ''}`)
    const response = await request.json()
    return response
  }
}

export const getHoldingTimes = async () => await (await fetch(`/api/reports/HoldingTimePerMonth`)).json()
export const getJobHeatmap = async (blockchain = '') => await (await fetch(`/api/reports/JobHeatmap?blockchain=${blockchain}`)).json()
export const getStakedTokens = async () => await (await fetch(`/api/reports/StakedTokensByDay`)).json()
export const getDataHolderByID = async (dataHolderID) => await (await fetch(`/api/nodes/dataholder/${dataHolderID}`)).json()
export const getDataCreatorByID = async (dataCreatorID) => await (await fetch(`/api/nodes/datacreator/${dataCreatorID}`)).json()
export const getHoldingTimeByMonth = async (month='') => await (await fetch(`/api/mynodes/GetHoldingTimeByMonth?nodeID=${month}`)).json()
export const getNodeStats = async () => await (await fetch('/api/mynodes/GetNodeStats')).json()
export const getJobsPerMonth = async () => await (await fetch('/api/mynodes/JobsPerMonth')).json()
export const getRecentJobsByNodeID = async (nodeId = '') => await (await fetch(`/api/mynodes/RecentJobs?nodeID=${nodeId}`)).json()


export const getJobs = async () => {
  const request = await fetch(`/api/home/JobsChartDataV3`)
  const jobs = await request.json()
  return jobs
}

export const getBlockchainDistribution = async () => {
  const request = await fetch(`/api/home/24HJobBlockchainDistribution`)
  const blockchainDistribution = await request.json()
  return blockchainDistribution
}

export const getDashboardData = async () => {
  const request = await fetch(`/api/home/HomeV3?excludeBreakdown=true`)
  const dashboardData = await request.json()
  return dashboardData
}