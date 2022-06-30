<script>
  import Unauthorized from './unauthorized.svelte'
  import { isAuthenticated } from '$stores/auth' 

  import { onMount } from 'svelte'
  import Spinner from './components/spinner.svelte'

  let loadingPayouts = false
  let blockchains
  let payouts = []
  let includeActiveJobs = true
  let includeCompletedJobs = true
  let selectedBlockchainID

  onMount(async () => {
		blockchains = await getBlockchainsFromAPI()
	})

  const getBlockchainsFromAPI = async () => {
    const request = await fetch(`/api/blockchain/GetBlockchains`)
    const blockchains = await request.json()
    return blockchains
  }

  const handleBlockchainSelection = async () => {
    loadingPayouts = true
    payouts = await getPayoutsFromAPI()
    loadingPayouts = false
  }

  const getPayoutsFromAPI = async () => {
    const request = await fetch(`/api/mynodes/PossibleJobPayouts?includeActiveJobs=${includeActiveJobs}&includeCompletedJobs=${includeCompletedJobs}&blockchainID=${selectedBlockchainID}&dateFrom=null&dateTo=null`)
    const payouts = await request.json()
    return payouts
  }
</script>

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
  <main class="flex h-full justify-center pb-16 pt-16">
    <div class="container flex flex-col gap-6 px-6 mx-auto">
      
      <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6">
        <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">How to use this page?</h4>
        <p class="text-sm text-gray-500 dark:text-gray-400 dark:bg-gray-800">This page allows you to payout jobs easier when you have many jobs or many nodes to manage. Currently you must have all your nodes on the same management wallet, while this page will semi work if you have multiple management wallets you'll experience more validation errors when using this. When you have selected the nodes you wish to payout, OT Hub will check they will succeed and then prompt each transaction to open in MetaMask.Note that this does not reduce the amount of transactions you need to send to the blockchain, you still need one transaction per payout.</p>
      </div>

      <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6">
        <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Filters</h4>
        <div class="flex gap-4">
          <select bind:value={selectedBlockchainID} on:change="{handleBlockchainSelection}" class="w-auto bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-gray-300 dark:focus:ring-blue-500 dark:focus:border-blue-500">
            <option value="" disabled selected>Choose a blockchain</option>
            {#each blockchains as blockchain}
              <option value={blockchain.ID}>{blockchain.BlockchainName}</option>
            {/each}
          </select>
          <div class="flex items-center">
            <input bind:checked={includeActiveJobs} on:change="{handleBlockchainSelection}" id="activeJobsCheckbox" type="checkbox" value="" class="w-4 h-4 text-blue-600 bg-gray-100 rounded border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
            <label for="activeJobsCheckbox" class="ml-2 text-sm text-gray-500 dark:text-gray-400 dark:bg-gray-800">Active jobs</label>
          </div>
          <div class="flex items-center">
            <input bind:checked={includeCompletedJobs} on:change="{handleBlockchainSelection}" id="completedJobsCheckbox" type="checkbox" value="" class="w-4 h-4 text-blue-600 bg-gray-100 rounded border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
            <label for="completedJobsCheckbox" class="ml-2 text-sm text-gray-500 dark:text-gray-400 dark:bg-gray-800">Completed jobs</label>
          </div>
        </div>
      </div>

      <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6">
        <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Payouts</h4>
      
        <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
          <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
              <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                <tr>
                  <th scope="col" class="p-4">
                    <div class="flex items-center">
                      <input id="checkbox-all" type="checkbox" class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
                      <label for="checkbox-all" class="sr-only">checkbox</label>
                    </div>
                  </th>
                  <th scope="col" class="p-4 pl-0">Offer ID</th>
                  <th scope="col" class="p-4 pl-0">Paid out</th>
                  <th scope="col" class="p-4 pl-0">Total reward</th>
                  <th scope="col" class="p-4 pl-0">Estimated payout</th>
                  <th scope="col" class="p-4 pl-0 whitespace-nowrap">Last payout</th>
                </tr>
              </thead>
              {#if !selectedBlockchainID || loadingPayouts || !payouts.length}
                <tbody>
                  <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                    <td colspan="6" class="w-4 p-4">
                      {#if !selectedBlockchainID}⚠️ Please, select your blockchain
                      {:else if loadingPayouts}<Spinner />
                      {:else if !payouts.length}⚠️ No payouts were found within the selected blockchain
                      {/if}
                    </td>
                  </tr>
                </tbody>
              {:else}
                <tbody>
                  {#each payouts as payout}
                    <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                      <td class="w-4 p-4">
                        <div class="flex items-center">
                          <input id="checkbox-table-1" type="checkbox" class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
                          <label for="checkbox-table-1" class="sr-only">checkbox</label>
                        </div>
                      </td>
                      <td scope="row" class="p-4 pl-0 text-sm whitespace-nowrap">{payout.OfferID}
                      </td>
                      <td class="p-4 pl-0 text-sm whitespace-nowrap">
                        <span>{payout.PaidAmount.toFixed(16)}</span>
                        <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">TRAC</span>
                      </td>
                      <td class="p-4 pl-0 text-sm whitespace-nowrap">
                        <span>{payout.TokenAmount.toFixed(16)}</span>
                        <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">TRAC</span>
                      </td>
                      <td class="p-4 pl-0 text-sm whitespace-nowrap">
                        <span>{payout.EstimatedPayout.toFixed(16)}</span>
                        <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">TRAC</span>
                      </td>
                      <td class="p-4 pl-0 text-sm whitespace-nowrap">{new Date(payout.LastPayout).toLocaleString('en-GB', { dateStyle: 'short' })}</td>
                    </tr>
                  {/each}
                </tbody>
              {/if}
          </table>
      </div>
    </div>

    <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6">
      <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Process Payouts via MetaMask</h4>
      <div class="p-4 text-sm text-gray-700 bg-gray-100 rounded-lg dark:bg-gray-700 dark:text-gray-300" role="alert">
        You have the wrong network selected in MetaMask. You need 100 to be selected. See <a href="https://chainid.network/" target="_blank" class="underline">here</a> to translate the network ID numbers into something more useful.
      </div>
    </div>

  </main>
{/if}
