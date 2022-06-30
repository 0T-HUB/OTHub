<script>
  import { onMount } from 'svelte'

  const getArrayOfNNumbersFromZero = (quantityOfNumbersToGet) => Array.from(Array(quantityOfNumbersToGet).keys())
  const getArrayOfNNumbersFromNumber = (quantityOfNumbersToGet, startingNumber) => getArrayOfNNumbersFromZero(quantityOfNumbersToGet).map(addendum => startingNumber + addendum)
  const getArrayOfNNumbersToNumber = (quantityOfNumbersToGet, endingNumber) => getArrayOfNNumbersFromZero(quantityOfNumbersToGet).map(addendum => endingNumber - addendum).sort()
  const createPaginationButtonsFromArrayOfPageNumbersAndSetActiveButtonBySelectedPageNumber = (arrayOfPageNumbers, selectedPageNumber) => arrayOfPageNumbers.map(pageNumber => ({ pageNumber, ...(pageNumber === selectedPageNumber) && { active: true }}))

  let jobsOnPage = []
  let numberOfJobsLastTwentyYears
  const numberOfJobsPerPage = 20
  const numberOfVisiblePaginationButtons = 5
  let numberOfPages
  let paginationButtons = createPaginationButtonsFromArrayOfPageNumbersAndSetActiveButtonBySelectedPageNumber(getArrayOfNNumbersFromNumber(numberOfVisiblePaginationButtons, 1), 1)
  let loading = true
  let activePage = 1
  let keyToSortBy = 'FinalizedTimestamp'
  let sortOrder = 'DESC'
  let searchQuery

  onMount(async () => {
    numberOfJobsLastTwentyYears = await(await fetch('https://v5api.othub.info/api/Jobs/jobcreatedcountinperiod?timePeriod=years&time=20')).json()
    numberOfPages = Math.ceil(numberOfJobsLastTwentyYears / numberOfJobsPerPage)
		jobsOnPage = await getJobsFromAPI()
	})

  const getJobsFromAPI = async (searchKey) => {
    loading = true
    const requestURL = `https://v5api.othub.info/api/jobs/paging?_sort=${keyToSortBy}&_order=${sortOrder}${searchKey && searchQuery ? '&' + searchKey + '_like=' + searchQuery : ''}&_page=${activePage}&_limit=${numberOfJobsPerPage}`
    const jobs = await(await fetch(requestURL)).json()
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

  const recalculatePaginationAndSetActivePageByPageNumber = async (selectedPageNumber) => {
    const arrayOfNumbersOfFirstPaginationButtons = getArrayOfNNumbersFromNumber(numberOfVisiblePaginationButtons, 1)
    const arrayOfNumbersOfLastPaginationButtons = getArrayOfNNumbersToNumber(numberOfVisiblePaginationButtons, numberOfPages)
    const medianIndexOfNumberOfVisiblePaginationButtons = (numberOfVisiblePaginationButtons % 2) + 1
    const arrayOfNumbersOfFirstPaginationButtonsToMedian = arrayOfNumbersOfFirstPaginationButtons.slice(0, medianIndexOfNumberOfVisiblePaginationButtons)
    const arrayOfNumbersOfLastPaginationButtonsFromMedian = arrayOfNumbersOfLastPaginationButtons.slice(medianIndexOfNumberOfVisiblePaginationButtons)
    const selectedPageIsAmongFirstPaginationButtonsToMedian = arrayOfNumbersOfFirstPaginationButtonsToMedian.includes(selectedPageNumber)
    const selectedPageIsAmongLastPaginationButtonsFromMedian = arrayOfNumbersOfLastPaginationButtonsFromMedian.includes(selectedPageNumber)
    if (selectedPageIsAmongFirstPaginationButtonsToMedian) paginationButtons = createPaginationButtonsFromArrayOfPageNumbersAndSetActiveButtonBySelectedPageNumber(arrayOfNumbersOfFirstPaginationButtons, selectedPageNumber)
    if (selectedPageIsAmongLastPaginationButtonsFromMedian) paginationButtons = createPaginationButtonsFromArrayOfPageNumbersAndSetActiveButtonBySelectedPageNumber(arrayOfNumbersOfLastPaginationButtons, selectedPageNumber)
    const arrayOfNumbersOfPaginationButtonsAroundSelectedPageNumber = getArrayOfNNumbersFromNumber(numberOfVisiblePaginationButtons, selectedPageNumber - 2)
    if (!selectedPageIsAmongFirstPaginationButtonsToMedian && !selectedPageIsAmongLastPaginationButtonsFromMedian) paginationButtons = createPaginationButtonsFromArrayOfPageNumbersAndSetActiveButtonBySelectedPageNumber(arrayOfNumbersOfPaginationButtonsAroundSelectedPageNumber, selectedPageNumber)
    activePage = selectedPageNumber
    jobsOnPage = await getJobsFromAPI()
  }

  const loadFirstPage = () => {
    const currentActivePageNumber = getCurrentActivePageNumber()
    if (currentActivePageNumber !== 1) recalculatePaginationAndSetActivePageByPageNumber(1)
  }

  const loadLastPage = () => {
    const currentActivePageNumber = getCurrentActivePageNumber()
    if (currentActivePageNumber !== numberOfPages) recalculatePaginationAndSetActivePageByPageNumber(numberOfPages)
  }

  const loadNextPage = () => {
    const currentActivePageNumber = getCurrentActivePageNumber()
    if (currentActivePageNumber !== numberOfPages) recalculatePaginationAndSetActivePageByPageNumber(currentActivePageNumber + 1)
  }

  const loadPreviousPage = () => {
    const currentActivePageNumber = getCurrentActivePageNumber()
    if (currentActivePageNumber !== 1) recalculatePaginationAndSetActivePageByPageNumber(currentActivePageNumber - 1)
  }

  const getCurrentActivePageNumber = () => paginationButtons.find(paginationButton => paginationButton.active).pageNumber

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
              <!-- <th class="px-4 py-3">Price factor</th> -->
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
      <div class="grid px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800 min-h-50">
        <span class="flex items-center col-span-3"> Showing {(numberOfJobsPerPage * (activePage - 1)) + 1}-{numberOfJobsPerPage * activePage} of {numberOfJobsLastTwentyYears}</span>
        <span class="col-span-2">&nbsp;</span>
        <!-- Pagination -->
        <span class="flex col-span-4 mt-2 sm:mt-auto sm:justify-end">
          <nav aria-label="Table navigation">
            <ul class="inline-flex items-center">
              <li>
                <button class="px-3 py-1 rounded-md rounded-l-lg focus:outline-none focus:shadow-outline-purple" aria-label="First" on:click={loadFirstPage}>
                  <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M14.707 5.29303C14.8945 5.48056 14.9998 5.73487 14.9998 6.00003C14.9998 6.26519 14.8945 6.5195 14.707 6.70703L11.414 10L14.707 13.293C14.8892 13.4816 14.99 13.7342 14.9877 13.9964C14.9854 14.2586 14.8802 14.5094 14.6948 14.6948C14.5094 14.8803 14.2586 14.9854 13.9964 14.9877C13.7342 14.99 13.4816 14.8892 13.293 14.707L9.293 10.707C9.10553 10.5195 9.00021 10.2652 9.00021 10C9.00021 9.73487 9.10553 9.48056 9.293 9.29303L13.293 5.29303C13.4805 5.10556 13.7348 5.00024 14 5.00024C14.2652 5.00024 14.5195 5.10556 14.707 5.29303V5.29303Z" fill="black"/>
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M10.7068 5.29303C10.8943 5.48056 10.9996 5.73487 10.9996 6.00003C10.9996 6.26519 10.8943 6.5195 10.7068 6.70703L7.41379 10L10.7068 13.293C10.8889 13.4816 10.9897 13.7342 10.9875 13.9964C10.9852 14.2586 10.88 14.5094 10.6946 14.6948C10.5092 14.8803 10.2584 14.9854 9.99619 14.9877C9.73399 14.99 9.48139 14.8892 9.29279 14.707L5.29279 10.707C5.10532 10.5195 5 10.2652 5 10C5 9.73487 5.10532 9.48056 5.29279 9.29303L9.29279 5.29303C9.48031 5.10556 9.73462 5.00024 9.99979 5.00024C10.265 5.00024 10.5193 5.10556 10.7068 5.29303V5.29303Z" fill="black"/>
                  </svg>
                </button>
              </li>
              <li>
                <button class="px-3 py-1 rounded-md rounded-l-lg focus:outline-none focus:shadow-outline-purple" aria-label="Previous" on:click={loadPreviousPage}>
                  <svg aria-hidden="true" class="w-4 h-4 fill-current" viewBox="0 0 20 20">
                    <path d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" fill-rule="evenodd" />
                  </svg>
                </button>
              </li>                
            {#each paginationButtons as paginationButton}
              <li>
                <button 
                  class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple" 
                  class:text-white={paginationButton.active} 
                  class:transition-colors={paginationButton.active} 
                  class:duration-150={paginationButton.active} 
                  class:bg-purple-600={paginationButton.active} 
                  class:border={paginationButton.active} 
                  class:border-r-0={paginationButton.active} 
                  class:border-purple-600={paginationButton.active}
                  on:click={() => recalculatePaginationAndSetActivePageByPageNumber(paginationButton.pageNumber)}>
                  {paginationButton.pageNumber}
                </button>
              </li>
            {/each}
              <li>
                <button class="px-3 py-1 rounded-md rounded-r-lg focus:outline-none focus:shadow-outline-purple" aria-label="Next" on:click={loadNextPage}>
                  <svg class="w-4 h-4 fill-current" aria-hidden="true" viewBox="0 0 20 20">
                    <path d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" fill-rule="evenodd"/>
                  </svg>
                </button>
              </li>
              <li>
                <button class="px-3 py-1 rounded-md rounded-r-lg focus:outline-none focus:shadow-outline-purple" aria-label="Last" on:click={loadLastPage}>
                  <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M5.293 14.707C5.10553 14.5194 5.00021 14.2651 5.00021 14C5.00021 13.7348 5.10553 13.4805 5.293 13.293L8.586 9.99997L5.293 6.70697C5.11084 6.51837 5.01005 6.26576 5.01233 6.00357C5.01461 5.74137 5.11977 5.49056 5.30518 5.30515C5.49059 5.11974 5.7414 5.01457 6.0036 5.0123C6.2658 5.01002 6.5184 5.11081 6.707 5.29297L10.707 9.29297C10.8945 9.4805 10.9998 9.73481 10.9998 9.99997C10.9998 10.2651 10.8945 10.5194 10.707 10.707L6.707 14.707C6.51947 14.8944 6.26516 14.9998 6 14.9998C5.73484 14.9998 5.48053 14.8944 5.293 14.707V14.707Z" fill="black"/>
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M9.29321 14.707C9.10574 14.5194 9.00043 14.2651 9.00043 14C9.00043 13.7348 9.10574 13.4805 9.29321 13.293L12.5862 9.99997L9.29321 6.70697C9.11106 6.51837 9.01026 6.26576 9.01254 6.00357C9.01482 5.74137 9.11999 5.49056 9.3054 5.30515C9.4908 5.11974 9.74162 5.01457 10.0038 5.0123C10.266 5.01002 10.5186 5.11081 10.7072 5.29297L14.7072 9.29297C14.8947 9.4805 15 9.73481 15 9.99997C15 10.2651 14.8947 10.5194 14.7072 10.707L10.7072 14.707C10.5197 14.8944 10.2654 14.9998 10.0002 14.9998C9.73505 14.9998 9.48074 14.8944 9.29321 14.707V14.707Z" fill="black"/>
                  </svg>
                </button>
              </li>
            </ul>
          </nav>
        </span>
      </div>
    </div>
  </div>
</main>