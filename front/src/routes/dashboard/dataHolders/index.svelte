<script>
  import { onMount } from 'svelte'
  import { badge } from '$stores/badge'
  import * as API from '$modules/api.module'
  import Spinner from '../components/spinner.svelte'
  import Pagination from '../components/pagination.svelte'
  import Export from '../components/export.svelte'
  
  let activePageNumber = 1
  let dataHolders 
  let loading = false
  let nameSearchQuery
  let nodeIDSearchQuery
  let numberOfItemsPerPage = 25
  let sortKey = 'WonOffersLast7Days'
  let sortOrder = 'DESC'

  const getAndSetDataHolders = async () => {
    loading = true
    const dataHoldersRequest = await API.getDataHolders({activePageNumber, numberOfItemsPerPage, sortKey, sortOrder, nodeIDSearchQuery, nameSearchQuery})
    dataHolders = dataHoldersRequest?.data || []
    loading = false
  }

  const toggleSortOrder = () => {
    sortOrder = sortOrder === 'DESC' ? 'ASC' : 'DESC'
    getAndSetDataHolders()
  }
  
  const changeSortKeyToNewSortKey = (newSortKey) => {
    sortOrder = 'DESC'
    sortKey = newSortKey
    getAndSetDataHolders()
  }

  $: activePageNumber, numberOfItemsPerPage, getAndSetDataHolders()
  
  onMount(async () => await getAndSetDataHolders())

</script>

<svelte:head>
  <title>Data holders</title>
</svelte:head>

<main class="h-full pb-16 overflow-y-auto">
  <div class="container grid px-6 mx-auto">
    <header class="flex flex-row justify-between py-6 ">
      <h2 class="text-2xl font-semibold text-gray-700 dark:text-gray-200">Data holders</h2>
      {#if dataHolders}
        <Export data={dataHolders} />
      {/if}
    </header>
    <div class="w-full mb-8 overflow-hidden rounded-lg shadow-xs">
      <div class="w-full overflow-x-auto">
        <table class="w-full table-fixed dark:bg-gray-800">
          <thead>
            <tr class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800">
              <th class="w-96 px-4 py-3 whitespace-nowrap">
                <form class="relative" on:submit|preventDefault={getAndSetDataHolders}>
                  <input bind:value={nodeIDSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Node ID">
                  <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                    <button type="submit">
                      <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                    </button>
                  </div>
                </form>
              </th>
              <th class="w-32 px-4 py-3 whitespace-nowrap">
                <form class="relative" on:submit|preventDefault={getAndSetDataHolders}>
                  <input bind:value={nameSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal  block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Name">
                  <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                    <button type="submit">
                      <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                    </button>
                  </div>
                </form>
              </th>
              <th class="w-16 px-4 py-3 whitespace-nowrap">
                <button class="uppercase" on:click={() => changeSortKeyToNewSortKey('TotalWonOffers')}>Jobs</button>
                {#if sortKey === 'TotalWonOffers'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
              </th>
              <th class="w-24 px-4 py-3 whitespace-nowrap">
                <button class="uppercase" on:click={() => changeSortKeyToNewSortKey('WonOffersLast7Days')}>Jobs/week</button>
                {#if sortKey === 'WonOffersLast7Days'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
              </th>
              <th class="w-28 px-4 py-3 whitespace-nowrap">
                <button class="uppercase" on:click={() => changeSortKeyToNewSortKey('ActiveOffers')}>Active jobs</button>
                {#if sortKey === 'ActiveOffers'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
              </th>
              <th class="w-56 px-4 py-3 whitespace-nowrap">
                <button class="uppercase" on:click={() => changeSortKeyToNewSortKey('StakeTokens')}>Staked tokens</button>
                {#if sortKey === 'StakeTokens'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
              </th>
              <th class="w-56 px-4 py-3 whitespace-nowrap">
                <button class="uppercase" on:click={() => changeSortKeyToNewSortKey('StakeReservedTokens')}>Locked tokens</button>
                {#if sortKey === 'StakeReservedTokens'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
              </th>
              <th class="w-56 px-4 py-3 whitespace-nowrap">
                <button class="uppercase" on:click={() => changeSortKeyToNewSortKey('PaidTokens')}>Paidout tokens</button>
                {#if sortKey === 'PaidTokens'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
              </th>
            </tr>
          </thead>
            <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
              {#if loading}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3 text-sm ">
                    <Spinner />
                  </td>
                </tr>
              {:else}
                {#if !dataHolders.length}
                  <tr class="text-gray-700 dark:text-gray-400">
                    <td class="px-4 py-3 text-sm ">
                      ⚠️ No data holders were found.
                    </td>
                  </tr>
                {:else}
                  {#each dataHolders as dataHolder}
                    <tr class="text-gray-700 dark:text-gray-400 text-sm">
                      <td class="px-4 py-3 whitespace-nowrap"><a href="/dataHolders/{dataHolder.NodeId}" class="text-purple-400 hover:opacity-75">{dataHolder.NodeId}</a></td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolder.DisplayName || ''}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolder.TotalWonOffers}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolder.WonOffersLast7Days}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolder.ActiveOffers}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolder.StakeTokens}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolder.StakeReservedTokens}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataHolder.PaidTokens}</td>
                    </tr>
                  {/each}
                {/if}
              {/if}
            </tbody>
        </table>
      </div>
      <div class="px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800">
        <Pagination totalNumberOfItems={$badge.DataHolders} bind:activePageNumber bind:numberOfItemsPerPage />
      </div>
    </div>
  </div>
</main>