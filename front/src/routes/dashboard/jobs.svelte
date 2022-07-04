<script>
  import { onMount } from 'svelte'
  import Pagination from './components/pagination.svelte'

  let jobsOnPage = []
  let numberOfJobsPerPage = 20
  let loading = true
  let activePage = 1
  let keyToSortBy = 'FinalizedTimestamp'
  let sortOrder = 'DESC'
  let searchQuery

  onMount(async () => {
		jobsOnPage = await getJobsFromAPI()
	})

  const getJobsFromAPI = async (searchKey) => {
    loading = true
    const requestURL = `/api/jobs/paging?_sort=${keyToSortBy}&_order=${sortOrder}${searchKey && searchQuery ? '&' + searchKey + '_like=' + searchQuery : ''}&_page=${activePage}&_limit=${numberOfJobsPerPage}`
    const jobsRequest = await(await fetch(requestURL)).json()
    const jobs = jobsRequest.data
    loading = false
    return jobs
  }

  const changeKeyToSortByToKey = async (newSortByKey) => {
    keyToSortBy = newSortByKey
    jobsOnPage = await getJobsFromAPI()
  }

  const toggleSortOrder = async () => {
    sortOrder = sortOrder === 'DESC' ? 'ASC' : 'DESC'
    jobsOnPage = await getJobsFromAPI()
  }


  const searchByKey = async (searchKey) => {
    jobsOnPage = await getJobsFromAPI(searchKey)
  }

</script>

<svelte:head>
  <title>Jobs</title>
</svelte:head>

<main class="h-full pb-16 overflow-y-auto">
  <div class="container grid px-6 mx-auto">
    <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Jobs</h2>
    <!-- Jobs data array -->
    <div class="w-full mb-8 overflow-hidden rounded-lg shadow-xs">
      <div class="w-full overflow-x-auto">
        <table class="w-full whitespace-no-wrap">
          <thead>
            <tr
              class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800"
            >
              <th class="px-4 py-3">Blockchain</th>
              <th class="px-4 py-3 flex items-center gap-x-4">
                <span>Offer</span>
                <div class="relative w-full max-w-xl mr-6 focus-within:text-purple-500 inline">
                  <div class="absolute inset-y-0 flex items-center pl-2">
                    <svg class="w-4 h-4" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20">
                      <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd"></path>
                    </svg>
                  </div>
                  <input bind:value={searchQuery} on:input={() => searchByKey('OfferId')} class="w-full pl-8 pr-2 py-2 text-sm text-gray-700 placeholder-gray-600 bg-gray-100 border-0 rounded-md dark:placeholder-gray-500 dark:focus:shadow-outline-gray dark:focus:placeholder-gray-600 dark:bg-gray-700 dark:text-gray-200 focus:placeholder-gray-500 focus:bg-white focus:border-purple-300 focus:outline-none focus:shadow-outline-purple form-input" type="text" placeholder="Search" aria-label="Search"></div>
              </th>
              <th class="px-4 py-3">
                <button class="uppercase" on:click={() => changeKeyToSortByToKey('FinalizedTimestamp')}>Started</button>
                {#if keyToSortBy === 'FinalizedTimestamp'}
                  <button on:click={toggleSortOrder} class:rotate-180={sortOrder === 'ASC'}>
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 inline" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
                    </svg>
                  </button>
                {/if}
              </th>
              <th class="px-4 py-3">
                <button class="uppercase" on:click={() => changeKeyToSortByToKey('DataSetSizeInBytes')}>Data set size</button>
                {#if keyToSortBy === 'DataSetSizeInBytes'}
                  <button on:click={toggleSortOrder} class:rotate-180={sortOrder === 'ASC'}>
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 inline" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
                    </svg>
                  </button>
                {/if}
              </th>
              <th class="px-4 py-3">
                <button class="uppercase" on:click={() => changeKeyToSortByToKey('HoldingTimeInMinutes')}>Holding time</button>
                {#if keyToSortBy === 'HoldingTimeInMinutes'}
                  <button on:click={toggleSortOrder} class:rotate-180={sortOrder === 'ASC'}>
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 inline" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
                    </svg>
                  </button>
                {/if}
              </th>
              <th class="px-4 py-3">
                <button class="uppercase" on:click={() => changeKeyToSortByToKey('TokenAmountPerHolder')}>Token amount</button>
                {#if keyToSortBy === 'TokenAmountPerHolder'}
                  <button on:click={toggleSortOrder} class:rotate-180={sortOrder === 'ASC'}>
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 inline" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
                    </svg>
                  </button>
                {/if}
              </th>
              <th class="px-4 py-3">Status</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
            {#if loading}
            <tr class="text-gray-700 dark:text-gray-400">
              <td class="px-4 py-3" colspan="7">
                <svg role="status" class="relative bottom-0.5 inline w-5 h-5 mr-2 text-gray-200 animate-spin dark:text-gray-600 fill-purple-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
                  <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
                </svg>
                <span>Loading...</span>
              </td>
            </tr>
            {:else}
              {#if jobsOnPage.length === 0}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3" colspan="7">
                    <svg xmlns="http://www.w3.org/2000/svg" class="relative bottom-0.5 inline h-5 w-5 mr-2 text-gray-200 dark:text-gray-600 fill-purple-600" viewBox="0 0 20 20" fill="none">
                      <path fill-rule="evenodd" d="M13.477 14.89A6 6 0 015.11 6.524l8.367 8.368zm1.414-1.414L6.524 5.11a6 6 0 018.367 8.367zM18 10a8 8 0 11-16 0 8 8 0 0116 0z" clip-rule="evenodd" />
                    </svg>
                    <span>No results were found</span>
                  </td>
                </tr>
              {/if}
              {#each jobsOnPage as job}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-4 py-3">
                    <div class="flex items-center text-sm">
                      <div>
                        <p class="font-semibold">{job.BlockchainDisplayName}</p>
                        <!-- <p class="text-xs text-gray-600 dark:text-gray-400">10x Developer</p> -->
                      </div>
                    </div>
                  </td>
                  <td class="px-4 py-3 text-sm">{job.OfferId}</td>
                  <td class="px-4 py-3 text-sm">{new Date(job.CreatedTimestamp).toLocaleDateString('en-GB')}</td>
                  <td class="px-4 py-3 text-sm">{job.DataSetSizeInBytes}</td>
                  <td class="px-4 py-3 text-sm">{job.HoldingTimeInMinutes}</td>
                  <td class="px-4 py-3 text-sm">{job.TokenAmountPerHolder}</td>
                  <!-- <td class="px-4 py-3 text-sm"></td> -->
                  <td class="px-4 py-3 text-sm">
                    <span
                      class="px-2 py-1 font-semibold leading-tight text-green-700 bg-green-100 rounded-full dark:bg-green-700 dark:text-green-100"
                    >
                    {job.Status}
                    </span>
                  </td>
                </tr>
              {/each}
            {/if}
          </tbody>
        </table>
      </div>
      <div class="px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800">
        <Pagination on:pagination={getJobsFromAPI} totalNumberOfItems={jobsOnPage.length} bind:activePageNumber={activePage} bind:numberOfItemsPerPage={numberOfJobsPerPage}/>
      </div>
    </div>
  </div>
</main>