<script>
  import { page } from '$app/stores'
  import { onMount } from 'svelte'
  import * as API from '$modules/api.module'
  import Spinner from '../components/spinner.svelte'
  import Pagination from '../components/pagination.svelte'

  const { dataCreatorID } = $page.params

  let dataCreator
  let dataCreatorJobs
  let dataCreatorJobsCount
  let dataCreatorLitigations
  let dataCreatorLitigationsCount
  let dataCreatorTransfers
  let dataCreatorTransfersCount

  let dataCreatorSettings = {
    activePageNumber: 1,
    loading: false,
    offerIDSearchQuery: '',
    numberOfItemsPerPage: 10,
    sortKey: 'WonOffersLast7Days',
    sortOrder: 'DESC',
    dataCreatorID,
    toggleSortOrder: function () {
      this.sortOrder = this.sortOrder === 'DESC' ? 'ASC' : 'DESC'
      getAndSetDataCreatorJobs()
    },
    setSortKey: function (newSortKey) {
      this.sortKey = newSortKey
      getAndSetDataCreatorJobs()
    }
  }

  let litigationSettings = {
    activePageNumber: 1,
    loading: false,
    offerIDSearchQuery: '',
    dataHolderIDSearchQuery: '',
    numberOfItemsPerPage: 10,
    sortKey: 'Timestamp',
    sortOrder: 'DESC',
    dataCreatorID,
    toggleSortOrder: function () {
      this.sortOrder = this.sortOrder === 'DESC' ? 'ASC' : 'DESC'
      getAndSetDataCreatorLitigations()
    },
    setSortKey: function (newSortKey) {
      this.sortKey = newSortKey
      getAndSetDataCreatorLitigations()
    }
  }

  let transferSettings = {
    activePageNumber: 1,
    loading: false,
    transactionHashSearchQuery: '',
    numberOfItemsPerPage: 10,
    sortKey: 'Timestamp',
    sortOrder: 'DESC',
    dataCreatorID,
    toggleSortOrder: function () {
      this.sortOrder = this.sortOrder === 'DESC' ? 'ASC' : 'DESC'
      getAndSetDataCreatorTransfers()
    },
    setSortKey: function (newSortKey) {
      this.sortKey = newSortKey
      getAndSetDataCreatorTransfers()
    }
  }


  onMount(async () => {
    const dataCreatorRequest = await API.getDataCreatorByID(dataCreatorID)
    dataCreator = dataCreatorRequest.data
    dataCreator.Identities[0].Active = true
    await getAndSetDataCreatorJobs()
    await getAndSetDataCreatorLitigations()
    await getAndSetDataCreatorTransfers()
  })

  const activateTabByIdentityIndex = (selectedIdentityIndex) => {
    dataCreator = {
      ...dataCreator,
      Identities: dataCreator.Identities.map((identity, currentIdentityIndex) => {
        return {
          ...identity,
          Active: false,
          ...(currentIdentityIndex === selectedIdentityIndex) && { Active: true }
        }
      })
    }
  }

  const getAndSetDataCreatorJobs = async () => {
    dataCreatorSettings.loading = true
    const dataCreatorJobsRequest = await API.getDataCreatorJobsWithDataCreatorSettings(dataCreatorSettings)
    dataCreatorJobs = dataCreatorJobsRequest.data
    dataCreatorJobsCount = dataCreatorJobsRequest.totalCount
    dataCreatorSettings.loading = false
  }

  const getAndSetDataCreatorLitigations = async () => {
    litigationSettings.loading = true
    const dataCreatorLitigationsRequest = await API.getDataCreatorLitigationsWithLitigationSettings(litigationSettings)
    dataCreatorLitigations = dataCreatorLitigationsRequest.data
    dataCreatorLitigationsCount = dataCreatorLitigationsRequest.totalCount
    litigationSettings.loading = false
  }

  const getAndSetDataCreatorTransfers = async () => {
    transferSettings.loading = true
    const dataCreatorTransfersRequest = await API.getDataCreatorTransfersWithTransferSettings(transferSettings)
    dataCreatorTransfers = dataCreatorTransfersRequest.data
    dataCreatorTransfersCount = dataCreatorTransfersRequest.totalCount
    transferSettings.loading = false
  }


</script>

<main class="h-full p-16 overflow-y-auto">
  <div class="container flex flex-col gap-6 px-6 mx-auto">
    <div class="flex gap-6">
      <!-- Data creator -->
      <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6 ">
        <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Data creator</h4>
        {#if dataCreator === undefined}
          <div class="flex h-96 w-full items-center justify-center">
            <Spinner/>
          </div>
        {:else}
          <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400 table-fixed">
            <tbody>
              {#if dataCreator.DisplayName}
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                  <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Name</th>
                  <td class="pr-4 py-4">{dataCreator.DisplayName}</td>
                </tr>
              {/if}
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Node ID</th>
                <td class="pr-4 py-4">{dataCreator.NodeId}</td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Staked tokens</th>
                <td class="pr-4 py-4">{dataCreator.StakeTokens}</td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Locked tokens</th>
                <td class="pr-4 py-4">
                  <span>{dataCreator.StakeReservedTokens} </span>
                  <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">TRAC</span>
                  <span> / </span>
                  <span>{dataCreator.StakeReservedTokens * dataCreator.LiveTracUSDPrice}</span>
                  <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">USD</span>
                </td>
              </tr>
            </tbody>
          </table>
        {/if}
      </div>
      
      <!-- Identities -->
      <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6 ">
        <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Identities</h4>
        {#if dataCreator === undefined}
          <div class="flex h-96 w-full items-center justify-center">
            <Spinner/>
          </div>
        {:else}
          <div class="text-sm font-medium text-center text-gray-500 border-b border-gray-200 dark:text-gray-400 dark:border-gray-700">
            <ul class="flex flex-wrap -mb-px">
              {#each dataCreator.Identities as identity, identityIndex}          
                <li class="mr-2">
                  <button class="inline-block p-4 rounded-t-lg border-b-2 border-transparent hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300 {identity.Active ? 'active dark:text-blue-500 dark:border-blue-500' : ''}"
                  on:click={() => activateTabByIdentityIndex(identityIndex)}>
                  {identity.BlockchainName}
                  </button>
                </li>
              {/each}
            </ul>
          </div>
          {#each dataCreator.Identities as identity, identityIndex}
            {#if identity.Active}
              <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400 table-fixed">
                <tbody>
                  <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                    <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Staked tokens</th>
                    <td class="pr-4 py-4">
                      <span>{identity.Stake} </span>
                      <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">TRAC</span>
                      <span> / </span>
                      <span>{identity.Stake * dataCreator.LiveTracUSDPrice}</span>
                      <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">USD</span>
                    </td>
                  </tr>
                  <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                    <th scope="row" class="pr-4 py-4 w-32 font-medium text-gray-900 dark:text-white whitespace-nowrap">Locked tokens</th>
                    <td class="pr-4 py-4">
                      <span>{identity.StakeReserved} </span>
                      <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">TRAC</span>
                      <span> / </span>
                      <span>{identity.StakeReserved * dataCreator.LiveTracUSDPrice}</span>
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

    <!-- Offers -->
    <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6 relative overflow-hidden">
      <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Offers</h4>
      {#if dataCreatorJobs === undefined}
        <div class="flex h-96 w-full items-center justify-center">
          <Spinner/>
        </div>
      {:else if dataCreatorJobs.length}
        <div class="w-full overflow-x-auto">
          <table class="w-full table-fixed dark:bg-gray-800">
            <thead>
              <tr class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800">
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <form class="relative" on:submit|preventDefault={() => getAndSetDataCreatorJobs()}>
                    <input bind:value={dataCreatorSettings.offerIDSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Offer ID">
                    <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                      <button type="submit">
                        <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                      </button>
                    </div>
                  </form>
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">Blockchain</th>
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataCreatorSettings.setSortKey('CreatedTimestamp')}>Created</button>
                  {#if dataCreatorSettings.sortKey === 'CreatedTimestamp'}<button on:click={() => dataCreatorSettings.toggleSortOrder()}>{dataCreatorSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataCreatorSettings.setSortKey('FinalizedTimestamp')}>Start</button>
                  {#if dataCreatorSettings.sortKey === 'FinalizedTimestamp'}<button on:click={() => dataCreatorSettings.toggleSortOrder()}>{dataCreatorSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataCreatorSettings.setSortKey('EndTimestamp')}>End</button>
                  {#if dataCreatorSettings.sortKey === 'EndTimestamp'}<button on:click={() => dataCreatorSettings.toggleSortOrder()}>{dataCreatorSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataCreatorSettings.setSortKey('HoldingTimeInMinutes')}>Total time</button>
                  {#if dataCreatorSettings.sortKey === 'HoldingTimeInMinutes'}<button on:click={() => dataCreatorSettings.toggleSortOrder()}>{dataCreatorSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <span class="uppercase">Remaining</span>
                </th>
                <th class="w-48 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataCreatorSettings.setSortKey('TokenAmountPerHolder')}>Token amount</button>
                  {#if dataCreatorSettings.sortKey === 'TokenAmountPerHolder'}<button on:click={() => dataCreatorSettings.toggleSortOrder()}>{dataCreatorSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-32 px-4 py-3 whitespace-nowrap">
                  <span class="uppercase">Price Factor</span>
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => dataCreatorSettings.setSortKey('Status')}>Status</button>
                  {#if dataCreatorSettings.sortKey === 'Status'}<button on:click={() => dataCreatorSettings.toggleSortOrder()}>{dataCreatorSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
              {#if dataCreatorSettings.loading}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm ">
                    <Spinner />
                  </td>
                </tr>
              {:else}
                {#if !dataCreatorJobs.length}
                  <tr class="text-gray-700 dark:text-gray-400">
                    <td class="px-4 py-3 text-sm ">
                      ⚠️ No data holder jobs were found.
                    </td>
                  </tr>
                {:else}
                  {#each dataCreatorJobs as dataCreatorJob}
                    <tr class="text-gray-700 dark:text-gray-400 text-sm">
                      <td class="px-4 py-3 whitespace-nowrap"><!-- <a href="/dataCreators/{dataCreator.NodeId}" class="text-purple-400 hover:opacity-75">-->{dataCreatorJob.OfferId.substring(0, 6)}...{dataCreatorJob.OfferId.split('').slice(-6).join('')}<!--</a>--></td> 
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreatorJob.BlockchainDisplayName}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{new Date(dataCreatorJob.CreatedTimestamp).toLocaleDateString('en-GB')} {new Date(dataCreatorJob.CreatedTimestamp).toLocaleTimeString('en-GB')}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{new Date(dataCreatorJob.FinalizedTimestamp).toLocaleDateString('en-GB')} {new Date(dataCreatorJob.FinalizedTimestamp).toLocaleTimeString('en-GB')}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{new Date(dataCreatorJob.EndTimestamp).toLocaleDateString('en-GB')} {new Date(dataCreatorJob.EndTimestamp).toLocaleTimeString('en-GB')}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{Math.ceil((new Date(dataCreatorJob.EndTimestamp).getTime() - new Date(dataCreatorJob.FinalizedTimestamp).getTime()) / (1000 * 60 * 60 * 24))} days</td>
                      <td class="px-4 py-3 whitespace-nowrap">{Math.max(0, Math.ceil((new Date().getTime() - new Date(dataCreatorJob.EndTimestamp).getTime()) / (1000 * 60 * 60 * 24)))} days</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreatorJob.TokenAmountPerHolder}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreatorJob.EstimatedLambda || 0} ({dataCreatorJob.EstimatedLambdaConfidence || 0}% match)</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreatorJob.Status}</td>
                    </tr>
                  {/each}
                {/if}
              {/if}
            </tbody>
          </table>
        </div>
        <div class="px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800">
          <Pagination on:pagination={getAndSetDataCreatorJobs} totalNumberOfItems={dataCreatorJobsCount} bind:activePageNumber={dataCreatorSettings.activePageNumber} bind:numberOfItemsPerPage={dataCreatorSettings.numberOfItemsPerPage} />
        </div>
      {/if}
    </div>

    <!-- Litigations -->
    <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6 relative overflow-hidden">
      <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Litigations against node</h4>
      {#if dataCreatorLitigations === undefined}
        <div class="flex h-96 w-full items-center justify-center">
          <Spinner/>
        </div>
      {:else}
        <div class="w-full overflow-x-auto">
          <table class="w-full table-fixed dark:bg-gray-800">
            <thead>
              <tr class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800">
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <form class="relative" on:submit|preventDefault={() => getAndSetDataCreatorLitigations()}>
                    <input bind:value={litigationSettings.offerIDSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Offer ID">
                    <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                      <button type="submit">
                        <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                      </button>
                    </div>
                  </form>
                </th>
                <th class="w-20 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => litigationSettings.setSortKey('Timestamp')}>Timestamp</button>
                  {#if litigationSettings.sortKey === 'Timestamp'}<button on:click={() => litigationSettings.toggleSortOrder()}>{litigationSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-40 px-4 py-3 whitespace-nowrap">
                  <form class="relative" on:submit|preventDefault={() => getAndSetDataCreatorLitigations()}>
                    <input bind:value={litigationSettings.dataHolderIDSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Data holder">
                    <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                      <button type="submit">
                        <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                      </button>
                    </div>
                  </form>
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => litigationSettings.setSortKey('RequestedBlockIndex')}>Requested Block Index</button>
                  {#if litigationSettings.sortKey === 'RequestedBlockIndex'}<button on:click={() => litigationSettings.toggleSortOrder()}>{litigationSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => litigationSettings.setSortKey('RequestedObjectIndex')}>Requested Object Index</button>
                  {#if litigationSettings.sortKey === 'RequestedObjectIndex'}<button on:click={() => litigationSettings.toggleSortOrder()}>{litigationSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
              {#if litigationSettings.loading}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm ">
                    <Spinner />
                  </td>
                </tr>
              {:else if !dataCreatorLitigations.length}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm ">
                    ⚠️ No litigations were found.
                  </td>
                </tr>
              {:else}
                {#each dataCreatorLitigations as dataCreatorLitigation}
                  <tr class="text-gray-700 dark:text-gray-400 text-sm">
                    <td class="px-4 py-3 whitespace-nowrap"><!-- <a href="/dataCreators/{dataCreator.NodeId}" class="text-purple-400 hover:opacity-75">-->{dataCreatorLitigation.OfferId.substring(0, 6)}...{dataCreatorLitigation.OfferId.split('').slice(-6).join('')}<!--</a>--></td> 
                    <td class="px-4 py-3 whitespace-nowrap">{new Date(dataCreatorLitigation.Timestamp).toLocaleDateString('en-GB')} {new Date(dataCreatorLitigation.Timestamp).toLocaleTimeString('en-GB')}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{dataCreatorLitigation.NodeId}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{dataCreatorLitigation.RequestedBlockIndex}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{dataCreatorLitigation.RequestedObjectIndex}</td>
                  </tr>
                {/each}
              {/if}
            </tbody>
          </table>
        </div>
        <div class="px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800">
          <Pagination on:pagination={getAndSetDataCreatorLitigations} totalNumberOfItems={dataCreatorLitigationsCount} bind:activePageNumber={litigationSettings.activePageNumber} bind:numberOfItemsPerPage={litigationSettings.numberOfItemsPerPage} />
        </div>
      {/if}
    </div>

    <!-- Deposits & withdrawals -->
    <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6 relative overflow-hidden">
      <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Deposits & withdrawals</h4>
      {#if dataCreatorTransfers === undefined}
        <div class="flex h-96 w-full items-center justify-center">
          <Spinner/>
        </div>
      {:else}
        <div class="w-full overflow-x-auto">
          <table class="w-full table-fixed dark:bg-gray-800">
            <thead>
              <tr class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800">
                <th class="w-64 px-4 py-3 whitespace-nowrap">
                  <form class="relative" on:submit|preventDefault={() => getAndSetDataCreatorTransfers()}>
                    <input bind:value={transferSettings.transactionHashSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Transaction hash">
                    <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                      <button type="submit">
                        <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                      </button>
                    </div>
                  </form>
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => transferSettings.setSortKey('Timestamp')}>Timestamp</button>
                  {#if transferSettings.sortKey === 'Timestamp'}<button on:click={() => transferSettings.toggleSortOrder()}>{transferSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-12 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase">Type</button>
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => transferSettings.setSortKey('Amount')}>Token Amount</button>
                  {#if transferSettings.sortKey === 'Amount'}<button on:click={() => transferSettings.toggleSortOrder()}>{transferSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </th>
                <th class="w-24 px-4 py-3 whitespace-nowrap">
                  <button class="uppercase" on:click={() => transferSettings.setSortKey('GasUsed')}>Transaction fee</button>
                  {#if transferSettings.sortKey === 'GasUsed'}<button on:click={() => transferSettings.toggleSortOrder()}>{transferSettings.sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
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
              {:else if !dataCreatorTransfers.length}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm" colspan="4">
                    ⚠️ No transfers were found.
                  </td>
                </tr>
              {:else}
                {#each dataCreatorTransfers as dataCreatorTransfer}
                  <tr class="text-gray-700 dark:text-gray-400 text-sm">
                    <td class="px-4 py-3 whitespace-nowrap"><a href="{dataCreatorTransfer.RealTransactionUrl}" class="text-purple-400 hover:opacity-75">{dataCreatorTransfer.TransactionHash}</a></td> 
                    <td class="px-4 py-3 whitespace-nowrap">{new Date(dataCreatorTransfer.Timestamp).toLocaleDateString('en-GB')} {new Date(dataCreatorTransfer.Timestamp).toLocaleTimeString('en-GB')}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{dataCreatorTransfer.Amount < 0 ? 'Withdrawal' : 'Deposit'}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{dataCreatorTransfer.Amount}</td>
                    <td class="px-4 py-3 whitespace-nowrap">{(dataCreatorTransfer.GasUsed / dataCreatorTransfer.GasPrice).toFixed(8)} {dataCreatorTransfer.GasTicker}</td>
                  </tr>
                {/each}
              {/if}
            </tbody>
          </table>
        </div>
        <div class="px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800">
          <Pagination on:pagination={getAndSetDataCreatorTransfers} totalNumberOfItems={dataCreatorTransfersCount} bind:activePageNumber={transferSettings.activePageNumber} bind:numberOfItemsPerPage={transferSettings.numberOfItemsPerPage} />
        </div>
      {/if}
    </div>
  </div>
</main>