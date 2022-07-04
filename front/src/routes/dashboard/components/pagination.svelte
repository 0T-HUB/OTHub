<script>
  import { createEventDispatcher } from 'svelte'

  const dispatch = createEventDispatcher()
  const dispatchPaginationEvent = () => dispatch('pagination', { activePageNumber, numberOfItemsPerPage })

  const setActivePageNumber = (pageNumber)  => {
    activePageNumber = pageNumber
    dispatchPaginationEvent()
    arrayOfVisiblePaginationButtons = getArrayOfVisiblePaginationButtons()
  }

  const handleChangeOfNumberOfItemsPerPage = () => {
    arrayOfVisiblePaginationButtons = getArrayOfVisiblePaginationButtons()
    if (activePageNumber > getNumberOfPages()) activePageNumber = getNumberOfPages()
    dispatchPaginationEvent()
  }

  const getArrayOfVisiblePaginationButtons = () => {
    const arrayOfNumberOfPages = Array.from({length: getNumberOfPages()}, (_, i) => i + 1)
    const medianOfnumberOfVisiblePaginationButtons = Math.ceil(numberOfVisiblePaginationButtons / 2)
    if (activePageNumber >= getNumberOfPages() - medianOfnumberOfVisiblePaginationButtons) return arrayOfNumberOfPages.slice(-Math.abs(numberOfVisiblePaginationButtons))
    if (activePageNumber <= medianOfnumberOfVisiblePaginationButtons) return Array.from({length: numberOfVisiblePaginationButtons}, (_, i) => i + 1)
    const firstVisiblePageNumberIndex = activePageNumber - medianOfnumberOfVisiblePaginationButtons
    const lastVisiblePageNumberIndex = firstVisiblePageNumberIndex + numberOfVisiblePaginationButtons
    return arrayOfNumberOfPages.slice(firstVisiblePageNumberIndex, lastVisiblePageNumberIndex)
  }

  const getNumberOfPages = () => Math.ceil(totalNumberOfItems / numberOfItemsPerPage)

  export let totalNumberOfItems
  let numberOfVisiblePaginationButtons = 5
  export let activePageNumber
  export let numberOfItemsPerPage
  let possibleNumbersOfItemsPerPage = [10, 25, 50, 100]
  let arrayOfVisiblePaginationButtons = getArrayOfVisiblePaginationButtons()

</script>

<div class="flex items-center gap-8 justify-between h-6">
  {#if totalNumberOfItems}
    <div class="flex items-center gap-8">
      <span class="flex items-center">Showing {numberOfItemsPerPage * (activePageNumber - 1) + 1}-{numberOfItemsPerPage * (activePageNumber)} of {totalNumberOfItems}</span>
      <div class="flex items-center gap-1">
        <span>Show</span>
        <select bind:value={numberOfItemsPerPage} on:change={handleChangeOfNumberOfItemsPerPage} class="block text-xs text-center text-gray-400 bg-transparent border-2 border-transparent !border-b-gray-700 appearance-none focus:outline-none focus:ring-0">
          {#each possibleNumbersOfItemsPerPage as possibleNumberOfItemsPerPage}
            <option value={possibleNumberOfItemsPerPage}>{possibleNumberOfItemsPerPage}</option>
          {/each}
        </select>
        <span>items per page</span>
      </div>
    </div>
    <div class="flex col-span-4 mt-2 sm:mt-auto sm:justify-end">
      <nav aria-label="Table navigation">
        <ul class="inline-flex items-center">
          <li>
            <button on:click={() => setActivePageNumber(1)} class="px-3 py-1 rounded-md rounded-l-lg focus:outline-none focus:shadow-outline-purple disabled:opacity-0" aria-label="First" disabled={activePageNumber === 1}>
              <svg class="w-4 h-4" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path class="fill-gray-600" fill-rule="evenodd" clip-rule="evenodd" d="M14.707 5.29303C14.8945 5.48056 14.9998 5.73487 14.9998 6.00003C14.9998 6.26519 14.8945 6.5195 14.707 6.70703L11.414 10L14.707 13.293C14.8892 13.4816 14.99 13.7342 14.9877 13.9964C14.9854 14.2586 14.8802 14.5094 14.6948 14.6948C14.5094 14.8803 14.2586 14.9854 13.9964 14.9877C13.7342 14.99 13.4816 14.8892 13.293 14.707L9.293 10.707C9.10553 10.5195 9.00021 10.2652 9.00021 10C9.00021 9.73487 9.10553 9.48056 9.293 9.29303L13.293 5.29303C13.4805 5.10556 13.7348 5.00024 14 5.00024C14.2652 5.00024 14.5195 5.10556 14.707 5.29303V5.29303Z"/>
                <path class="fill-gray-600" fill-rule="evenodd" clip-rule="evenodd" d="M10.7068 5.29303C10.8943 5.48056 10.9996 5.73487 10.9996 6.00003C10.9996 6.26519 10.8943 6.5195 10.7068 6.70703L7.41379 10L10.7068 13.293C10.8889 13.4816 10.9897 13.7342 10.9875 13.9964C10.9852 14.2586 10.88 14.5094 10.6946 14.6948C10.5092 14.8803 10.2584 14.9854 9.99619 14.9877C9.73399 14.99 9.48139 14.8892 9.29279 14.707L5.29279 10.707C5.10532 10.5195 5 10.2652 5 10C5 9.73487 5.10532 9.48056 5.29279 9.29303L9.29279 5.29303C9.48031 5.10556 9.73462 5.00024 9.99979 5.00024C10.265 5.00024 10.5193 5.10556 10.7068 5.29303V5.29303Z"/>
              </svg>
            </button>
          </li>
          <li>
            <button on:click={() => setActivePageNumber(activePageNumber - 1)} class="px-3 py-1 rounded-md rounded-l-lg focus:outline-none focus:shadow-outline-purple disabled:opacity-0" aria-label="Previous" disabled={activePageNumber === 1}>
              <svg aria-hidden="true" class="w-4 h-4" viewBox="0 0 20 20">
                <path class="fill-gray-600" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" fill-rule="evenodd" />
              </svg>
            </button>
          </li>
          {#each arrayOfVisiblePaginationButtons as pageNumber}
            <li>
              <button on:click={() => setActivePageNumber(pageNumber)} class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple transition-colors duration-150 { pageNumber === activePageNumber ? 'text-white bg-purple-600 border border-r-0 border-purple-600' : ''}">{pageNumber}</button>
            </li>
          {/each}
          <li>
            <button on:click={() => setActivePageNumber(activePageNumber + 1)} class="px-3 py-1 rounded-md rounded-r-lg focus:outline-none focus:shadow-outline-purple disabled:opacity-0" aria-label="Next" disabled={activePageNumber === getNumberOfPages()}>
              <svg class="w-4 h-4" aria-hidden="true" viewBox="0 0 20 20">
                <path class="fill-gray-600" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" fill-rule="evenodd" />
              </svg>
            </button>
          </li>
          <li>
            <button on:click={() => setActivePageNumber(getNumberOfPages())} class="px-3 py-1 rounded-md rounded-r-lg focus:outline-none focus:shadow-outline-purple disabled:opacity-0" aria-label="Last" disabled={activePageNumber === getNumberOfPages()}>
              <svg class="w-4 h-4" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path class="fill-gray-600" fill-rule="evenodd" clip-rule="evenodd" d="M5.293 14.707C5.10553 14.5194 5.00021 14.2651 5.00021 14C5.00021 13.7348 5.10553 13.4805 5.293 13.293L8.586 9.99997L5.293 6.70697C5.11084 6.51837 5.01005 6.26576 5.01233 6.00357C5.01461 5.74137 5.11977 5.49056 5.30518 5.30515C5.49059 5.11974 5.7414 5.01457 6.0036 5.0123C6.2658 5.01002 6.5184 5.11081 6.707 5.29297L10.707 9.29297C10.8945 9.4805 10.9998 9.73481 10.9998 9.99997C10.9998 10.2651 10.8945 10.5194 10.707 10.707L6.707 14.707C6.51947 14.8944 6.26516 14.9998 6 14.9998C5.73484 14.9998 5.48053 14.8944 5.293 14.707V14.707Z"/>
                <path class="fill-gray-600" fill-rule="evenodd" clip-rule="evenodd" d="M9.29321 14.707C9.10574 14.5194 9.00043 14.2651 9.00043 14C9.00043 13.7348 9.10574 13.4805 9.29321 13.293L12.5862 9.99997L9.29321 6.70697C9.11106 6.51837 9.01026 6.26576 9.01254 6.00357C9.01482 5.74137 9.11999 5.49056 9.3054 5.30515C9.4908 5.11974 9.74162 5.01457 10.0038 5.0123C10.266 5.01002 10.5186 5.11081 10.7072 5.29297L14.7072 9.29297C14.8947 9.4805 15 9.73481 15 9.99997C15 10.2651 14.8947 10.5194 14.7072 10.707L10.7072 14.707C10.5197 14.8944 10.2654 14.9998 10.0002 14.9998C9.73505 14.9998 9.48074 14.8944 9.29321 14.707V14.707Z"/>
              </svg>
            </button>
          </li>
        </ul>
      </nav>
    </div>
  {/if}
</div>