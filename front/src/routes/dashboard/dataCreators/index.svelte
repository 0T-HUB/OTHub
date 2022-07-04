<script>
  import { onMount } from 'svelte'
  import { badge } from '$stores/badge'
  import * as API from '$modules/api.module'
  import Spinner from '../components/spinner.svelte'
  import Pagination from '../components/pagination.svelte'
  import Export from '../components/export.svelte'
  
  let activePageNumber = 1
  let dataCreators 
  let loading = false
  let nameSearchQuery
  let nodeIDSearchQuery
  let numberOfItemsPerPage = 25
  let sortKey = 'OffersLast7Days'
  let sortOrder = 'DESC'

  const getAndSetdataCreators = async () => {
    loading = true
    const dataCreatorsRequest = await API.getDataCreators({activePageNumber, numberOfItemsPerPage, sortKey, sortOrder, nodeIDSearchQuery, nameSearchQuery})
    dataCreators = dataCreatorsRequest.data
    loading = false
  }

  const toggleSortOrder = () => {
    sortOrder = sortOrder === 'DESC' ? 'ASC' : 'DESC'
    getAndSetdataCreators()
  }
  
  const changeSortKeyToNewSortKey = (newSortKey) => {
    sortOrder = 'DESC'
    sortKey = newSortKey
    getAndSetdataCreators()
  }

  $: activePageNumber, numberOfItemsPerPage, getAndSetdataCreators()
  
  onMount(async () => await getAndSetdataCreators())

</script>

<svelte:head>
  <title>Data holders</title>
</svelte:head>

<main class="h-full pb-16 overflow-y-auto">
  <div class="container grid px-6 mx-auto">
    <header class="flex flex-row justify-between py-6 ">
      <h2 class="text-2xl font-semibold text-gray-700 dark:text-gray-200">Data creators</h2>
      {#if dataCreators}
        <Export data={dataCreators} />
      {/if}
    </header>
    <div class="w-full mb-8 overflow-hidden rounded-lg shadow-xs">
      <div class="w-full overflow-x-auto">
        <table class="w-full table-fixed dark:bg-gray-800">
          <thead>
            <tr class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800">
              <th class="w-96 px-4 py-3 whitespace-nowrap">
                <form class="relative" on:submit|preventDefault={getAndSetdataCreators}>
                  <input bind:value={nodeIDSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Node ID">
                  <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                    <button type="submit">
                      <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                    </button>
                  </div>
                </form>
              </th>
              <th class="w-32 px-4 py-3 whitespace-nowrap">
                <form class="relative" on:submit|preventDefault={getAndSetdataCreators}>
                  <input bind:value={nameSearchQuery} class="placeholder:font-semibold placeholder:uppercase font-normal  block p-3 w-full text-xs text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Name">
                  <div class="flex absolute inset-y-0 right-4 items-center pl-3">
                    <button type="submit">
                      <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
                    </button>
                  </div>
                </form>
              </th>
              <th class="w-16 px-4 py-3 whitespace-nowrap">
                <button class="uppercase" on:click={() => changeSortKeyToNewSortKey('OffersTotal')}>Jobs</button>
                {#if sortKey === 'OffersTotal'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
              </th>
              <th class="w-16 px-4 py-3">
                <div class="flex gap-1">
                  <button class="uppercase text-left" on:click={() => changeSortKeyToNewSortKey('OffersLast7Days')}>Last week</button>
                  {#if sortKey === 'OffersLast7Days'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </div>
              </th>
              <th class="w-24 px-4 py-3 whitespace-nowrap">
                <button class="uppercase" on:click={() => changeSortKeyToNewSortKey('LastJob')}>Last job</button>
                {#if sortKey === 'LastJob'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
              </th>
              <th class="w-56 px-4 py-3 whitespace-nowrap">
                <button class="uppercase" on:click={() => changeSortKeyToNewSortKey('StakeTokens')}>Staked tokens</button>
                {#if sortKey === 'StakeTokens'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
              </th>
              <th class="w-56 px-4 py-3 whitespace-nowrap">
                <button class="uppercase" on:click={() => changeSortKeyToNewSortKey('StakeReservedTokens')}>Locked tokens</button>
                {#if sortKey === 'StakeReservedTokens'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
              </th>
              <th class="w-24 px-4 py-3">
                <div class="flex gap-1">
                  <button class="uppercase text-left" on:click={() => changeSortKeyToNewSortKey('AvgDataSetSizeKB')}>Dataset Size</button>
                  {#if sortKey === 'AvgDataSetSizeKB'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </div>
              </th>
              <th class="w-24 px-4 py-3">
                <div class="flex gap-1">
                  <button class="uppercase text-left" on:click={() => changeSortKeyToNewSortKey('AvgHoldingTimeInMinutes')}>Avg Holding Time</button>
                  {#if sortKey === 'AvgHoldingTimeInMinutes'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </div>
              </th>
              <th class="w-24 px-4 py-3">
                <div class="flex gap-1">
                  <button class="uppercase text-left" on:click={() => changeSortKeyToNewSortKey('AvgTokenAmountPerHolder')}>Avg Amount</button>
                  {#if sortKey === 'AvgTokenAmountPerHolder'}<button on:click={toggleSortOrder}>{sortOrder === 'DESC' ? '▾' : '▴'}</button>{/if}
                </div>
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
                {#if !dataCreators.length}
                  <tr class="text-gray-700 dark:text-gray-400">
                    <td class="px-4 py-3 text-sm ">
                      ⚠️ No data holders were found.
                    </td>
                  </tr>
                {:else}
                  {#each dataCreators as dataCreator}
                    <tr class="text-gray-700 dark:text-gray-400 text-sm">
                      <td class="px-4 py-3 whitespace-nowrap"><a href="/dataCreators/{dataCreator.NodeId}" class="text-purple-400 hover:opacity-75">{dataCreator.NodeId}</a></td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreator.DisplayName || ''}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreator.OffersTotal}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreator.OffersLast7Days}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{new Date(dataCreator.LastJob).toLocaleString('en-GB', { dateStyle: 'short' })}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreator.StakeTokens}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreator.StakeReservedTokens}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreator.AvgDataSetSizeKB}</td>
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreator.AvgHoldingTimeInMinutes}</td> <!-- convert to days-->
                      <td class="px-4 py-3 whitespace-nowrap">{dataCreator.AvgTokenAmountPerHolder}</td>
                    </tr>
                  {/each}
                {/if}
              {/if}
            </tbody>
        </table>
      </div>
      <div class="px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800">
        <Pagination totalNumberOfItems={$badge.DataCreators} bind:activePageNumber bind:numberOfItemsPerPage />
      </div>
    </div>
  </div>
</main>