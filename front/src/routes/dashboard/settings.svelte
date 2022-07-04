<script>
  import Unauthorized from './unauthorized.svelte'
  import { isAuthenticated } from '$stores/auth' 
  import { onMount } from 'svelte'

  let dataHolders = []
  let nodeIsBeingAdded = false
  let priceCalculationMode
  let nodeThatIsBeingAdded = {
    NewDisplayName: '',
    NodeId: ''
  }

  onMount(async () => {
    dataHolders = await getDataHoldersFromAPI()
    priceCalculationMode = await getPriceCalculationModeFromAPI()
  })

  const changePriceCalculationMode = async () => {
    const priceCalculationRequest = await fetch(`/api/mynodes/UpdateMyNodesPriceCalculationMode?mode=${priceCalculationMode}`, { method: 'POST' })
    return priceCalculationRequest.data
  }

  const getPriceCalculationModeFromAPI = async () => {
    const request = await fetch(`/api/mynodes/MyNodesPriceCalculationMode`)
    const priceCalculationModeRequest = await request.json()
    const priceCalculationMode = priceCalculationModeRequest.data
    return priceCalculationMode
  }

  const getDataHoldersFromAPI = async () => {
    const dataHoldersRequest = await (await fetch(`/api/nodes/dataholders?ercVersion=1&restrictToMyNodes=true&_sort=DisplayName&_order=ASC&_page=1&_limit=9999`)).json()
    console.log(dataHoldersRequest)
    let dataHolders = dataHoldersRequest?.data || []
    dataHolders = dataHolders.map(dataHolder => {
      return {
        ...dataHolder,
        NewDisplayName: dataHolder.DisplayName
      }
    })
    return dataHolders
  }

  const discardDataHolderEdits = (updatedDataHolder) => {
    const { NodeId, DisplayName } = updatedDataHolder
    dataHolders = dataHolders.map(dataHolder => {
      return {
        ...dataHolder,
        ...(dataHolder.NodeId === NodeId) && { 
          beingEdited: false,
          NewDisplayName: DisplayName
        },

      }
    })
  }

  const saveOrUpdateNode = async (updatedDataHolder) => {
    const { NodeId, NewDisplayName } = updatedDataHolder
    await fetch(`/api/mynodes/addeditnode?nodeID=${NodeId}&name=${NewDisplayName}`, { method: 'POST' })
    dataHolders = dataHolders.map(dataHolder => {
      return {
        ...dataHolder,
        ...(dataHolder.NodeId === NodeId) && { 
          beingEdited: false,
          DisplayName: NewDisplayName
        }
      }
    })
    dataHolders = await getDataHoldersFromAPI()
    nodeIsBeingAdded = false
  }

  const deleteDataHolder = async (deletedDataHolder) => {
    const { NodeId } = deletedDataHolder
    await fetch(`/api/mynodes/deletenode?nodeID=${NodeId}`, { method: 'DELETE' })
    dataHolders = dataHolders.filter(dataHolder => dataHolder.NodeId !== deletedDataHolder.NodeId)
  }

  const importNodes = async () => {
    const importedNodes = prompt('Please, enter the text that you copied from the old OT Hub website')
    await fetch(`/api/mynodes/importnodes?identities=${importedNodes}`, { method: 'POST' })
    dataHolders = await getDataHoldersFromAPI()
  }

</script>

<svelte:head>
  <title>Settings</title>
</svelte:head>

{#if $isAuthenticated === undefined}
  <main class="h-full flex items-center justify-center">
    <svg role="status" class="inline w-6 h-6 mr-2 text-gray-200 animate-spin dark:text-gray-600 fill-purple-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
      <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
    </svg>
  </main>
{/if}

{#if $isAuthenticated === false}
  <Unauthorized />
{/if}

{#if $isAuthenticated}
  <main class="h-full pb-16 overflow-y-auto">
    <!-- nodes array -->
    <div class="container grid px-6 mx-auto">
      <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">Settings</h2>
      <!-- With actions -->
      
      <div class="flex justify-between items-center mb-4">
        <h4 class="text-lg font-semibold text-gray-600 dark:text-gray-300">My nodes</h4>
        {#if !nodeIsBeingAdded}
          <button on:click={() => nodeIsBeingAdded = true} class="px-5 py-3 font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple">Add a node</button>
        {/if}
      </div>

      <div class="w-full overflow-hidden rounded-lg shadow-xs">
        <div class="w-full overflow-x-auto">
          <table class="w-full whitespace-no-wrap dark:bg-gray-800">
            <thead>
              <tr
                class="text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800"
              >
                <th class="px-2 py-2 text-xs">Actions</th>
                <th class="px-2 py-2 text-xs">Node id</th>
                <th class="px-2 py-2 text-xs">Name</th>
                <th class="px-2 py-2 text-xs">Jobs</th>
                <th class="px-2 py-2 text-xs whitespace-nowrap">Jobs (one week)</th>
                <th class="px-2 py-2 text-xs whitespace-nowrap">Active jobs</th>
                <th class="px-2 py-2 text-xs">Staked tockens</th>
                <th class="px-2 py-2 text-xs">Locked tockens</th>
                <th class="px-2 py-2 text-xs">Paidout tockens</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y dark:divide-gray-700 dark:bg-gray-800">
              {#if !dataHolders.length && !nodeIsBeingAdded}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-2 py-2 text-sm" colspan="9">⚠️ You have no nodes</td>
                </tr>
              {/if}

              {#if nodeIsBeingAdded}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-2 py-2">
                    <div class="flex items-center space-x-1">
                      <button on:click={() => saveOrUpdateNode(nodeThatIsBeingAdded)} class="flex items-center justify-between px-2 py-2 text-sm font-medium leading-5 text-purple-600 rounded-lg dark:text-gray-400 focus:outline-none focus:shadow-outline-gray" aria-label="Edit">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                          <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                        </svg>
                      </button>
                      <button on:click={() => nodeIsBeingAdded = false} class="flex items-center justify-between px-2 py-2 text-sm font-medium leading-5 text-purple-600 rounded-lg dark:text-gray-400 focus:outline-none focus:shadow-outline-gray" aria-label="Delete">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                          <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                        </svg>
                      </button>
                    </div>
                  </td>
                  <td class="px-2 py-2 text-xs">
                    <input bind:value={nodeThatIsBeingAdded.NodeId} class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 read-only:text-gray-500 text-xs" placeholder="Node ID">
                  </td>
                  <td class="px-2 py-2 text-xs">
                    <input bind:value={nodeThatIsBeingAdded.NewDisplayName} class="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 read-only:text-gray-500 text-xs w-20" placeholder="Name">
                  </td>
                </tr>
              {/if}


              {#each dataHolders as dataHolder}
                <tr class="text-gray-700 dark:text-gray-400">
                  <td class="px-2 py-2 text-xs">
                    <div class="flex items-center space-x-1 text-sm">
                    {#if !dataHolder.beingEdited}
                      <button on:click={() => dataHolder.beingEdited = true} class="flex items-center justify-between px-2 py-2 text-sm font-medium leading-5 text-purple-600 rounded-lg dark:text-gray-400 focus:outline-none focus:shadow-outline-gray" aria-label="Edit">
                        <svg class="w-5 h-5" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20">
                          <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
                        </svg>
                      </button>
                      <button on:click={() => deleteDataHolder(dataHolder)} class="flex items-center justify-between px-2 py-2 text-sm font-medium leading-5 text-purple-600 rounded-lg dark:text-gray-400 focus:outline-none focus:shadow-outline-gray" aria-label="Delete">
                        <svg class="w-5 h-5" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z"  clip-rule="evenodd" />
                        </svg>
                      </button>
                    {:else}
                      <button on:click={() => saveOrUpdateNode(dataHolder)} class="flex items-center justify-between px-2 py-2 text-sm font-medium leading-5 text-purple-600 rounded-lg dark:text-gray-400 focus:outline-none focus:shadow-outline-gray" aria-label="Save">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                          <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                        </svg>
                      </button>
                      <button on:click={() => discardDataHolderEdits(dataHolder)} class="flex items-center justify-between px-2 py-2 text-sm font-medium leading-5 text-purple-600 rounded-lg dark:text-gray-400 focus:outline-none focus:shadow-outline-gray" aria-label="Cancel">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                          <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                        </svg>
                      </button>
                    {/if}
                    </div></td>
                  <td class="px-2 py-2 text-xs" title={dataHolder.NodeId}>{dataHolder.NodeId}</td>
                  <td class="px-2 py-2 text-xs whitespace-nowrap">
                    {#if !dataHolder.beingEdited} 
                      {dataHolder.DisplayName}
                    {:else}
                      <input bind:value={dataHolder.NewDisplayName} class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 read-only:text-gray-500 text-xs w-20" placeholder="Name">
                    {/if}
                  </td>
                  <td class="px-2 py-2 text-xs">{dataHolder.TotalWonOffers}</td>
                  <td class="px-2 py-2 text-xs">{dataHolder.WonOffersLast7Days}</td>
                  <td class="px-2 py-2 text-xs">{dataHolder.ActiveOffers}</td>
                  <td class="px-2 py-2 text-xs">{dataHolder.StakeTokens}</td>
                  <td class="px-2 py-2 text-xs">{dataHolder.StakeReservedTokens}</td>
                  <td class="px-2 py-2 text-xs">{dataHolder.PaidTokens}</td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
        <div
          class="grid px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800"
        >
          <!-- <span class="flex items-center col-span-3"> Showing 21-30 of 100 </span>
          <span class="col-span-2" /> -->
          <!-- Pagination -->
          <!-- <span class="flex col-span-4 mt-2 sm:mt-auto sm:justify-end">
            <nav aria-label="Table navigation">
              <ul class="inline-flex items-center">
                <li>
                  <button
                    class="px-3 py-1 rounded-md rounded-l-lg focus:outline-none focus:shadow-outline-purple"
                    aria-label="Previous"
                  >
                    <svg class="w-4 h-4 fill-current" aria-hidden="true" viewBox="0 0 20 20">
                      <path
                        d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
                        clip-rule="evenodd"
                        fill-rule="evenodd"
                      />
                    </svg>
                  </button>
                </li>
                <li>
                  <button class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple">
                    1
                  </button>
                </li>
                <li>
                  <button class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple">
                    2
                  </button>
                </li>
                <li>
                  <button
                    class="px-3 py-1 text-white transition-colors duration-150 bg-purple-600 border border-r-0 border-purple-600 rounded-md focus:outline-none focus:shadow-outline-purple"
                  >
                    3
                  </button>
                </li>
                <li>
                  <button class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple">
                    4
                  </button>
                </li>
                <li>
                  <span class="px-3 py-1">...</span>
                </li>
                <li>
                  <button class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple">
                    8
                  </button>
                </li>
                <li>
                  <button class="px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple">
                    9
                  </button>
                </li>
                <li>
                  <button
                    class="px-3 py-1 rounded-md rounded-r-lg focus:outline-none focus:shadow-outline-purple"
                    aria-label="Next"
                  >
                    <svg class="w-4 h-4 fill-current" aria-hidden="true" viewBox="0 0 20 20">
                      <path
                        d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                        clip-rule="evenodd"
                        fill-rule="evenodd"
                      />
                    </svg>
                  </button>
                </li>
              </ul>
            </nav>
          </span> -->
        </div>
      </div>
      <br />
      <!-- button import nodes -->
      <div>
        <button on:click={importNodes} class="px-5 py-3 font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple">Import nodes</button>
      </div>

      <br />
      <!-- Nodes amount calculation-->
      <div class="px-4 py-3 mb-8 bg-white rounded-lg shadow-md dark:bg-gray-800">
        <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">
          USD Amount Calculation
        </h4>
        <div class="mt-4 text-sm">
          <div class="mt-2">
            <label class="inline-flex items-center text-gray-600 dark:text-gray-400">
              <input
                bind:group={priceCalculationMode}
                value={0}
                on:change={changePriceCalculationMode}
                type="radio"
                class="text-purple-600 form-radio focus:border-purple-400 focus:outline-none focus:shadow-outline-purple dark:focus:shadow-outline-gray"
                name="priceCalculationMode"
              />
              <span class="ml-2">Use TRAC Price at Job Start</span>
            </label>
            <label class="inline-flex items-center ml-6 text-gray-600 dark:text-gray-400">
              <input
                bind:group={priceCalculationMode}
                value={1}
                on:change={changePriceCalculationMode}
                type="radio"
                class="text-purple-600 form-radio focus:border-purple-400 focus:outline-none focus:shadow-outline-purple dark:focus:shadow-outline-gray"
                name="priceCalculationMode"
              />
              <span class="ml-2">Use Live TRAC Price</span>
            </label>
          </div>
        </div>
      </div>

      <!-- link telegram -->
    <a
    class="flex items-center justify-between p-4 mb-8 text-sm font-semibold text-purple-100 bg-purple-600 rounded-lg shadow-md focus:outline-none focus:shadow-outline-purple"
    href=""
  >
    <div class="flex items-center">
      <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 24 24">
        <path
          d="M18.384,22.779c0.322,0.228 0.737,0.285 1.107,0.145c0.37,-0.141 0.642,-0.457 0.724,-0.84c0.869,-4.084 2.977,-14.421 3.768,-18.136c0.06,-0.28 -0.04,-0.571 -0.26,-0.758c-0.22,-0.187 -0.525,-0.241 -0.797,-0.14c-4.193,1.552 -17.106,6.397 -22.384,8.35c-0.335,0.124 -0.553,0.446 -0.542,0.799c0.012,0.354 0.25,0.661 0.593,0.764c2.367,0.708 5.474,1.693 5.474,1.693c0,0 1.452,4.385 2.209,6.615c0.095,0.28 0.314,0.5 0.603,0.576c0.288,0.075 0.596,-0.004 0.811,-0.207c1.216,-1.148 3.096,-2.923 3.096,-2.923c0,0 3.572,2.619 5.598,4.062Zm-11.01,-8.677l1.679,5.538l0.373,-3.507c0,0 6.487,-5.851 10.185,-9.186c0.108,-0.098 0.123,-0.262 0.033,-0.377c-0.089,-0.115 -0.253,-0.142 -0.376,-0.064c-4.286,2.737 -11.894,7.596 -11.894,7.596Z"
        />
      </svg>
      <span
        >By linking your OT Hub account with Telegram you can automatically receive messages from
        our bot when certain events happen on the ODN.</span
      >
    </div>
    <span>Telegram Notif &RightArrow;</span>
  </a>

    </div>

    
  </main>
{/if}