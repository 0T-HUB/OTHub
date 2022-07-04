<script>
  import { page } from '$app/stores'
  import { onMount } from 'svelte'
  import * as API from '$modules/api.module'
  import Spinner from '../components/spinner.svelte'
  import Pagination from '../components/pagination.svelte'

  const { dataHolderID } = $page.params

  let dataHolder
  let dataHolderJobs
  let dataHolderLitigations
  let dataHolderPayouts
  let dataHolderTransfers

  let dataHolderSettings = {
    activePageNumber: 1,
    loading: false,
    offerIDSearchQuery: '',
    numberOfItemsPerPage: 10,
    sortKey: 'WonOffersLast7Days',
    sortOrder: 'DESC',
    dataHolderID,
    toggleSortOrder: () => {
      this.sortOrder = this.sortOrder === 'DESC' ? 'ASC' : 'DESC'
      getAndSetDataHolderJobs()
    }
  }

  let litigationSettings = {
    activePageNumber: 1,
    loading: false,
    offerIDSearchQuery: '',
    numberOfItemsPerPage: 10,
    sortKey: 'Timestamp',
    sortOrder: 'DESC',
    dataHolderID,
    toggleSortOrder: function () {
      this.sortOrder = this.sortOrder === 'DESC' ? 'ASC' : 'DESC'
      getAndSetDataHolderLitigations()
    }
  }

  let payoutSettings = {
    activePageNumber: 1,
    loading: false,
    offerIDSearchQuery: '',
    transactionHashSearchQuery: '',
    numberOfItemsPerPage: 10,
    sortKey: 'Timestamp',
    sortOrder: 'DESC',
    dataHolderID,
    toggleSortOrder: function () {
      this.sortOrder = this.sortOrder === 'DESC' ? 'ASC' : 'DESC'
      getAndSetDataHolderPayouts()
    }
  }

  let transferSettings = {
    activePageNumber: 1,
    loading: false,
    transactionHashSearchQuery: '',
    numberOfItemsPerPage: 10,
    sortKey: 'Timestamp',
    sortOrder: 'DESC',
    dataHolderID,
    toggleSortOrder: function () {
      this.sortOrder = this.sortOrder === 'DESC' ? 'ASC' : 'DESC'
      getAndSetDataHolderTransfers()
    }
  }


  onMount(async () => {
    const dataHolderRequest = await API.getDataHolderByID(dataHolderID)
    dataHolder = dataHolderRequest.data
    dataHolder.Identities[0].Active = true
    await getAndSetDataHolderJobs()
    await getAndSetDataHolderLitigations()
    await getAndSetDataHolderPayouts()
    await getAndSetDataHolderTransfers()
  })

  const activateTabByIdentityIndex = (selectedIdentityIndex) => {
    dataHolder = {
      ...dataHolder,
      Identities: dataHolder.Identities.map((identity, currentIdentityIndex) => {
        return {
          ...identity,
          Active: false,
          ...(currentIdentityIndex === selectedIdentityIndex) && { Active: true }
        }
      })
    }
  }

  const getAndSetDataHolderJobs = async () => {
    dataHolderSettings.loading = true
    const dataHolderJobsRequest = await API.getDataHolderJobsWithDataHolderSettings(dataHolderSettings)
    dataHolderJobs = dataHolderJobsRequest.data
    dataHolderSettings.loading = false
  }

  const getAndSetDataHolderLitigations = async () => {
    litigationSettings.loading = true
    const dataHolderLitigationsRequest = await API.getDataHolderLitigationsWithLitigationSettings(litigationSettings)
    dataHolderLitigations = dataHolderLitigationsRequest.data
    litigationSettings.loading = false
  }

  const getAndSetDataHolderPayouts = async () => {
    payoutSettings.loading = true
    const dataHolderPayoutsRequest = await API.getDataHolderPayoutsWithPayoutSettings(payoutSettings)
    dataHolderPayouts = dataHolderPayoutsRequest.data
    payoutSettings.loading = false
  }

  const getAndSetDataHolderTransfers = async () => {
    transferSettings.loading = true
    const dataHolderTransfersRequest = await API.getTransfersWithTransferSettings(transferSettings)
    dataHolderTransfers = dataHolderTransfersRequest.data
    transferSettings.loading = false
  }


</script>

<main class="h-full p-16 overflow-y-auto">
  <div class="container grid gap-6 px-6 mx-auto">
    <div class="grid gap-6 overflow-x-hidden">
      <!-- Data holder -->
      <div class="dark:bg-gray-800 dark:border-gray-700 px-6 py-6 overflow-x-auto">
        <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Data holder</h4>
        {#if dataHolder === undefined}
          <div class="flex h-96 w-full items-center justify-center">
            <Spinner/>
          </div>
        {:else}
          <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400 table-fixed">
            <tbody>
              {#if dataHolder.DisplayName}
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                  <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Name</th>
                  <td class="pr-4 py-4 whitespace-nowrap">{dataHolder.DisplayName}</td>
                </tr>
              {/if}
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Node ID</th>
                <td class="pr-4 py-4 whitespace-nowrap">{dataHolder.NodeId}</td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Staked tokens</th>
                <td class="pr-4 py-4 whitespace-nowrap">{dataHolder.StakeTokens}</td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Locked tokens</th>
                <td class="pr-4 py-4 whitespace-nowrap">
                  <span>{dataHolder.StakeReservedTokens} </span>
                  <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">TRAC</span>
                  <span> / </span>
                  <span>{dataHolder.StakeReservedTokens * dataHolder.LiveTracUSDPrice}</span>
                  <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">USD</span>
                </td>
              </tr>
            </tbody>
          </table>
        {/if}
      </div>
      
      <!-- Identities -->
      <div class="dark:bg-gray-800 dark:border-gray-700 px-6 py-6 overflow-x-auto">
        <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Identities</h4>
        {#if dataHolder === undefined}
          <div class="flex h-96 w-full items-center justify-center">
            <Spinner/>
          </div>
        {:else}
          <div class="text-sm font-medium text-center text-gray-500 border-b border-gray-200 dark:text-gray-400 dark:border-gray-700">
            <ul class="flex flex-wrap -mb-px">
              {#each dataHolder.Identities as identity, identityIndex}          
                <li class="mr-2">
                  <button class="inline-block p-4 rounded-t-lg border-b-2 border-transparent hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300 {identity.Active ? 'active dark:text-blue-500 dark:border-blue-500' : ''}"
                  on:click={() => activateTabByIdentityIndex(identityIndex)}>
                  {identity.BlockchainName}
                  </button>
                </li>
              {/each}
            </ul>
          </div>
          {#each dataHolder.Identities as identity, identityIndex}
            {#if identity.Active}
              <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400 table-fixed">
                <tbody>
                  <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                    <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Staked tokens</th>
                    <td class="pr-4 py-4  whitespace-nowrap">
                      <span>{identity.Stake} </span>
                      <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">TRAC</span>
                      <span> / </span>
                      <span>{identity.Stake * dataHolder.LiveTracUSDPrice}</span>
                      <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">USD</span>
                    </td>
                  </tr>
                  <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                    <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Locked tokens</th>
                    <td class="pr-4 py-4  whitespace-nowrap">
                      <span>{identity.StakeReserved} </span>
                      <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">TRAC</span>
                      <span> / </span>
                      <span>{identity.StakeReserved * dataHolder.LiveTracUSDPrice}</span>
                      <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">USD</span>
                    </td>
                  </tr>
                </tbody>
              </table>
            {/if}
          {/each}
        {/if}
      </div>
    </div>

    <!-- Jobs -->
    <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6 relative overflow-hidden">
      <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Jobs</h4>
      {#if dataHolderJobs === undefined}
        <div class="flex h-96 w-full items-center justify-center">
          <Spinner/>
        </div>
      {:else if dataHolderJobs.length}
        <div class="w-full overflow-x-auto">
          <table class="w-full table-fixed dark:bg-gray-800">
            <thead>
              <tr class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800">
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <form class="relative" on:submit|preventDefault={() => getAndSetDataHolderJobs()}>
                    <input bind:value={dataHolderSettings.offerIDSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Offer ID">
                    <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                      <button type="submit">
                        <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                      </button>
                    </div>
                  </form>
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">Blockchain</th>
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataHolderSettings.sortKey = 'FinalizedTimestamp'}>Start</button>
                  {#if dataHolderSettings.sortKey === 'TotalWonOffers'}<button on:click={dataHolderSettings.toggleSortOrder}>{dataHolderSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataHolderSettings.sortKey = 'EndTimestamp'}>End</button>
                  {#if dataHolderSettings.sortKey === 'EndTimestamp'}<button on:click={dataHolderSettings.toggleSortOrder}>{dataHolderSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataHolderSettings.sortKey = 'HoldingTimeInMinutes'}>Total time</button>
                  {#if dataHolderSettings.sortKey === 'HoldingTimeInMinutes'}<button on:click={dataHolderSettings.toggleSortOrder}>{dataHolderSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase">Remaining</button>
                </th>
                <th class="w-48 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataHolderSettings.sortKey = 'TokenAmountPerHolder'}>Token amount</button>
                  {#if dataHolderSettings.sortKey === 'TokenAmountPerHolder'}<button on:click={dataHolderSettings.toggleSortOrder}>{dataHolderSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-32 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase">Price Factor</button>
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataHolderSettings.sortKey = 'Status'}>Status</button>
                  {#if dataHolderSettings.sortKey === 'Status'}<button on:click={dataHolderSettings.toggleSortOrder}>{dataHolderSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-50 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataHolderSettings.sortKey = 'Paidout'}>Paidout</button>
                  {#if dataHolderSettings.sortKey === 'Paidout'}<button on:click={dataHolderSettings.toggleSortOrder}>{dataHolderSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
              {#if dataHolderSettings.loading}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm ">
                    <Spinner />
                  </td>
                </tr>
              {:else}
                {#if !dataHolderJobs.length}
                  <tr class="text-gray-700 dark:text-gray-400">
                    <td class="px-4 py-3 text-sm ">
                      ⚠️ No data holder jobs were found.
                    </td>
                  </tr>
                {:else}
                  {#each dataHolderJobs as dataHolderJob}
                    <tr class="text-gray-700 dark:text-gray-400 text-sm">
                      <td class="px-4 py-3 whitespace-nowrap"><!-- <a href="/dataHolders/{dataHolder.NodeId}" class="text-purple-400 hover:opacity-75">-->{dataHolderJob.OfferId.substring(0, 6)}...{dataHolderJob.OfferId.split('').slice(-6).join('')}<!--</a>--></td> 
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolderJob.BlockchainName}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{new Date(dataHolderJob.FinalizedTimestamp).toLocaleDateString('en-GB')} {new Date(dataHolderJob.FinalizedTimestamp).toLocaleTimeString('en-GB')}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{new Date(dataHolderJob.EndTimestamp).toLocaleDateString('en-GB')} {new Date(dataHolderJob.EndTimestamp).toLocaleTimeString('en-GB')}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{Math.ceil((new Date(dataHolderJob.EndTimestamp).getTime() - new Date(dataHolderJob.FinalizedTimestamp).getTime()) / (1000 * 60 * 60 * 24))} days</td>
                      <td class="px-4 py-3 whitespace-nowrap">{Math.max(0, Math.ceil((new Date().getTime() - new Date(dataHolderJob.EndTimestamp).getTime()) / (1000 * 60 * 60 * 24)))} days</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolderJob.TokenAmountPerHolder}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolderJob.EstimatedLambda} ({dataHolderJob.EstimatedLambdaConfidence}% match)</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolderJob.Status}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolderJob.PaidoutAmount || 0}</td>
                    </tr>
                  {/each}
                {/if}
              {/if}
            </tbody>
          </table>
        </div>
        <div class="px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800">
          <Pagination on:pagination={getAndSetDataHolderJobs} totalNumberOfItems={dataHolder?.TotalWonOffers} bind:activePageNumber={dataHolderSettings.activePageNumber} bind:numberOfItemsPerPage={dataHolderSettings.numberOfItemsPerPage} />
        </div>
      {/if}
    </div>

    <!-- Litigations -->
    <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6 relative overflow-hidden">
      <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Litigations against node</h4>
      {#if dataHolderLitigations === undefined}
        <div class="flex h-96 w-full items-center justify-center">
          <Spinner/>
        </div>
      {:else}
        <div class="w-full overflow-x-auto">
          <table class="w-full table-fixed dark:bg-gray-800">
            <thead>
              <tr class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800">
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <form class="relative" on:submit|preventDefault={() => getAndSetDataHolderLitigations()}>
                    <input bind:value={litigationSettings.offerIDSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Offer ID">
                    <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                      <button type="submit">
                        <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                      </button>
                    </div>
                  </form>
                </th>
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataHolderSettings.sortKey = 'FinalizedTimestamp'}>Start</button>
                  {#if dataHolderSettings.sortKey === 'TotalWonOffers'}<button on:click={dataHolderSettings.toggleSortOrder}>{dataHolderSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataHolderSettings.sortKey = 'EndTimestamp'}>End</button>
                  {#if dataHolderSettings.sortKey === 'EndTimestamp'}<button on:click={dataHolderSettings.toggleSortOrder}>{dataHolderSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataHolderSettings.sortKey = 'HoldingTimeInMinutes'}>Total time</button>
                  {#if dataHolderSettings.sortKey === 'HoldingTimeInMinutes'}<button on:click={dataHolderSettings.toggleSortOrder}>{dataHolderSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
              {#if dataHolderSettings.loading}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm ">
                    <Spinner />
                  </td>
                </tr>
              {:else if !dataHolderLitigations.length}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm ">
                    ⚠️ No litigations were found.
                  </td>
                </tr>
              {:else}
                {#each dataHolderLitigations as dataHolderLitigation}
                  <tr class="text-gray-700 dark:text-gray-400 text-sm">
                    <td class="px-4 py-3 whitespace-nowrap"><!-- <a href="/dataHolders/{dataHolder.NodeId}" class="text-purple-400 hover:opacity-75">-->{dataHolderLitigation.OfferId.substring(0, 6)}...{dataHolderLitigation.OfferId.split('').slice(-6).join('')}<!--</a>--></td> 
                    <td class="px-4 py-3 whitespace-nowrap">{new Date(dataHolderLitigation.Timestamp).toLocaleDateString('en-GB')} {new Date(dataHolderLitigation.Timestamp).toLocaleTimeString('en-GB')}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{dataHolderLitigation.RequestedBlockIndex || ''}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{dataHolderLitigation.RequestedObjectIndex || ''}</td>
                  </tr>
                {/each}
              {/if}
            </tbody>
          </table>
        </div>
        <div class="px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800">
          <Pagination on:pagination={getAndSetDataHolderLitigations} totalNumberOfItems={dataHolder?.LitigationCount} bind:activePageNumber={litigationSettings.activePageNumber} bind:numberOfItemsPerPage={litigationSettings.numberOfItemsPerPage} />
        </div>
      {/if}
    </div>

    <!-- Payouts -->
    <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6 relative overflow-hidden">
      <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Payouts for jobs</h4>
      {#if dataHolderPayouts === undefined}
        <div class="flex h-96 w-full items-center justify-center">
          <Spinner/>
        </div>
      {:else}
        <div class="w-full overflow-x-auto">
          <table class="w-full table-fixed dark:bg-gray-800">
            <thead>
              <tr class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800">
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <form class="relative" on:submit|preventDefault={() => getAndSetDataHolderPayouts()}>
                    <input bind:value={payoutSettings.offerIDSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Offer ID">
                    <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                      <button type="submit">
                        <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                      </button>
                    </div>
                  </form>
                </th>
                <th class="w-144 px-4 py-3 whitespace-nowrap">
                  <form class="relative" on:submit|preventDefault={() => getAndSetDataHolderPayouts()}>
                    <input bind:value={payoutSettings.transactionHashSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Transaction hash">
                    <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                      <button type="submit">
                        <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                      </button>
                    </div>
                  </form>
                </th>
                <th class="w-44 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => payoutSettings.sortKey = 'EndTimestamp'}>Timestamp</button>
                  {#if payoutSettings.sortKey === 'EndTimestamp'}<button on:click={dataHolderSettings.toggleSortOrder}>{payoutSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-48 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => payoutSettings.sortKey = 'HoldingTimeInMinutes'}>Token Amount</button>
                  {#if payoutSettings.sortKey === 'HoldingTimeInMinutes'}<button on:click={payoutSettings.toggleSortOrder}>{payoutSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => payoutSettings.sortKey = 'HoldingTimeInMinutes'}>Transaction fee</button>
                  {#if payoutSettings.sortKey === 'HoldingTimeInMinutes'}<button on:click={payoutSettings.toggleSortOrder}>{payoutSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
              {#if payoutSettings.loading}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm ">
                    <Spinner />
                  </td>
                </tr>
              {:else if !dataHolderPayouts.length}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm" colspan="4">
                    ⚠️ No payouts were found.
                  </td>
                </tr>
              {:else}
                {#each dataHolderPayouts as dataHolderPayout}
                  <tr class="text-gray-700 dark:text-gray-400 text-sm">
                    <td class="px-4 py-3 whitespace-nowrap"><!-- <a href="/dataHolders/{dataHolder.NodeId}" class="text-purple-400 hover:opacity-75">-->{dataHolderPayout.OfferId.substring(0, 6)}...{dataHolderPayout.OfferId.split('').slice(-6).join('')}<!--</a>--></td> 
                    <td class="px-4 py-3 whitespace-nowrap">{dataHolderPayout.TransactionHash}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{new Date(dataHolderPayout.Timestamp).toLocaleDateString('en-GB')} {new Date(dataHolderPayout.Timestamp).toLocaleTimeString('en-GB')}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{dataHolderPayout.Amount}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{(dataHolderPayout.GasUsed / dataHolderPayout.GasPrice).toFixed(8)} {dataHolderPayout.GasTicker}</td>
                  </tr>
                {/each}
              {/if}
            </tbody>
          </table>
        </div>
        <div class="px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800">
          <Pagination on:pagination={getAndSetDataHolderLitigations} totalNumberOfItems={dataHolder?.LitigationCount} bind:activePageNumber={litigationSettings.activePageNumber} bind:numberOfItemsPerPage={litigationSettings.numberOfItemsPerPage} />
        </div>
      {/if}
    </div>

    <!-- Deposits & withdrawals -->
    <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6 relative overflow-hidden">
      <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Deposits & withdrawals</h4>
      {#if dataHolderTransfers === undefined}
        <div class="flex h-96 w-full items-center justify-center">
          <Spinner/>
        </div>
      {:else}
        <div class="w-full overflow-x-auto">
          <table class="w-full table-fixed dark:bg-gray-800">
            <thead>
              <tr class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800">
                <th class="w-64 px-4 py-3 whitespace-nowrap">
                  <form class="relative" on:submit|preventDefault={() => getAndSetDataHolderTransfers()}>
                    <input bind:value={transferSettings.transactionHashSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Transaction hash">
                    <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                      <button type="submit">
                        <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                      </button>
                    </div>
                  </form>
                </th>
                <th class="w-32 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => transferSettings.sortKey = 'Timestamp'}>Timestamp</button>
                  {#if transferSettings.sortKey === 'Timestamp'}<button on:click={transferSettings.toggleSortOrder}>{transferSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-20 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase">Type</button>
                </th>
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => transferSettings.sortKey = 'Amount'}>Token Amount</button>
                  {#if transferSettings.sortKey === 'Amount'}<button on:click={transferSettings.toggleSortOrder}>{transferSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => transferSettings.sortKey = 'GasUsed'}>Transaction fee</button>
                  {#if transferSettings.sortKey === 'GasUsed'}<button on:click={transferSettings.toggleSortOrder}>{transferSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
              {#if transferSettings.loading}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm ">
                    <Spinner />
                  </td>
                </tr>
              {:else if !dataHolderTransfers.length}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm" colspan="4">
                    ⚠️ No transfers were found.
                  </td>
                </tr>
              {:else}
                {#each dataHolderTransfers as dataHolderTransfer}
                  <tr class="text-gray-700 dark:text-gray-400 text-sm">
                    <td class="w-64 px-4 py-3 whitespace-nowrap"><a href="{dataHolderTransfer.RealTransactionUrl}" class="text-purple-400 hover:opacity-75">{dataHolderTransfer.TransactionHash}</a></td> 
                    <td class="px-4 py-3 whitespace-nowrap">{new Date(dataHolderTransfer.Timestamp).toLocaleDateString('en-GB')} {new Date(dataHolderTransfer.Timestamp).toLocaleTimeString('en-GB')}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{dataHolderTransfer.Amount < 0 ? 'Withdrawal' : 'Deposit'}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{dataHolderTransfer.Amount}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{(dataHolderTransfer.GasUsed / dataHolderTransfer.GasPrice).toFixed(8)} {dataHolderTransfer.GasTicker}</td>
                  </tr>
                {/each}
              {/if}
            </tbody>
          </table>
        </div>
        <div class="px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800">
          <Pagination on:pagination={getAndSetDataHolderPayouts} totalNumberOfItems={dataHolder?.TotalWonOffers} bind:activePageNumber={transferSettings.activePageNumber} bind:numberOfItemsPerPage={transferSettings.numberOfItemsPerPage} />
        </div>
      {/if}
    </div>
  </div>
</main>