<script>
  import { web3, selectedAccount, chainId, defaultEvmStores, makeContractStore } from 'svelte-web3'

  import Unauthorized from './unauthorized.svelte'
  import { isAuthenticated } from '$stores/auth' 

  import { onMount } from 'svelte'
  import Spinner from './components/spinner.svelte'
  import Checkbox from './components/checkbox.svelte'

  let loadingPayouts = false
  let blockchains
  let payouts = []
  let includeActiveJobs = true
  let includeCompletedJobs = true
  let selectedBlockchainID
  let gasPrice

  let holdingAddress
  let litigationStorageAddress
  let holdingStorageAddress

  onMount(async () => {
		blockchains = await getBlockchainsFromAPI()
	})

  const getHoldingAddressFromAPI = async () => {
    const request = await fetch(`/api/contracts/getholdingaddresses?blockchainID=${selectedBlockchainID}`)
    const holdingAddresses = await request.json()
    const latestHoldingAddress = holdingAddresses.find(holdingAddress => holdingAddress.IsLatest).Address
    return latestHoldingAddress
  }

  const getLitigationStorageAddressFromAPI = async () => {
    const request = await fetch(`/api/contracts/getlitigationstorageaddresses?blockchainID=${selectedBlockchainID}`)
    const litigationStorageAddresses = await request.json()
    const latestLitigationStorageAddress = litigationStorageAddresses.find(litigationAddress => litigationAddress.IsLatest).Address
    return latestLitigationStorageAddress
  }

  const getHoldingStorageAddressFromAPI = async () => {
    const request = await fetch(`/api/contracts/getholdingstorageaddresses?blockchainID=${selectedBlockchainID}`)
    const holdingStorageAddresses = await request.json()
    const latestHoldingStorageAddress = holdingStorageAddresses.find(holdingStorageAddress => holdingStorageAddress.IsLatest).Address
    return latestHoldingStorageAddress
  }

  const connectToMetamask = () => {
    defaultEvmStores.setProvider()
  }

  const getBlockchainsFromAPI = async () => {
    const request = await fetch(`/api/blockchain/GetBlockchains`)
    const blockchains = await request.json()
    return blockchains
  }

  const handleBlockchainSelection = async () => {
    loadingPayouts = true
    payouts = await getPayoutsFromAPI()
    holdingAddress = await getHoldingAddressFromAPI()
    litigationStorageAddress = await getLitigationStorageAddressFromAPI()
    holdingStorageAddress = await getHoldingStorageAddressFromAPI()
    loadingPayouts = false
  }

  const getPayoutsFromAPI = async () => {
    const request = await fetch(`/api/mynodes/PossibleJobPayouts?includeActiveJobs=${includeActiveJobs}&includeCompletedJobs=${includeCompletedJobs}&blockchainID=${selectedBlockchainID}&dateFrom=null&dateTo=null`)
    const payouts = await request.json()
    return payouts
  }

  const verifySelectedPayouts = async () => {
    const selectedPayouts = payouts.filter(payout => payout.selected)
    for (let selectedPayout of selectedPayouts) {
      const verification = await verifySelectedPayout(selectedPayout)
      payouts = payouts.map(payout => {
        return {
          ...payout,
          ...(payout.OfferID === selectedPayout.OfferID) && { verified: verification.CanTryPayout }
        }
      })
    }
  }

  const verifySelectedPayout = async (selectedPayout) => {
    const { NodeID, OfferID, Identity } = selectedPayout
    const request = await fetch(`/api/nodes/dataholder/cantrypayout?nodeID=${NodeID}&offerId=${OfferID}&holdingAddress=${holdingAddress}&holdingStorageAddress=${holdingStorageAddress}&litigationStorageAddress=${litigationStorageAddress}&identity=${Identity}&blockchainID=${selectedBlockchainID}&selectedAddress=${$selectedAccount}`)
    const verification = await request.json()
    return verification
  }

  const handleCheckAll = (event) => {
    const { checked: allPayoutsSelectedStatus } = event.target
    payouts = payouts.map(payout => {
      return {
        ...payout,
        selected: allPayoutsSelectedStatus
      }
    })
  }

  const sendPayouts = async () => {
    const selectedPayouts = payouts.filter(payout => payout.selected)
    for (let selectedPayout of selectedPayouts) {
      await sendPayoutByIdentityAndOfferID(selectedPayout)
    }
  }

  const sendPayoutByIdentityAndOfferID = async ({ Identity, OfferID }) => {
    const contractInfo = [{
      constant: false,
      inputs: [
          {
            name: "identity",
            type: "address"
          },
          {
            name: "offerId",
            type: "uint256"
          }
      ],
      name: "payOut",
      outputs: [],
      payable: false,
      stateMutability: "nonpayable",
      type: "function"
    }]
    const contractInstance = new $web3.eth.Contract(contractInfo, holdingAddress)
    const builder = contractInstance.methods.payOut(Identity, OfferID)
    const response = await builder.send({
      from: $web3.currentProvider.selectedAddress, value: 0, gas: 300000,
      to: holdingAddress,
      gasPrice: $web3.utils.toHex($web3.utils.toWei(gasPrice, 'gwei'))
    })
    console.log(response)
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
          
          <Checkbox bind:checked={includeActiveJobs} on:change={handleBlockchainSelection}>Active jobs</Checkbox>
          <Checkbox bind:checked={includeCompletedJobs} on:change={handleBlockchainSelection}>Completed jobs</Checkbox>

        </div>
      </div>

      <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6">
        
        <div class="flex justify-between items-center">
          <h4 class="text-lg font-semibold text-gray-600 dark:text-gray-300">Payouts</h4>
          {#if !$selectedAccount}
            <button on:click={connectToMetamask} type="button" class="focus:outline-none text-white bg-purple-700 hover:bg-purple-800 focus:ring-4 focus:ring-purple-300 font-medium rounded-lg text-sm px-5 py-2.5 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900">Connect to MetaMask</button>      
          {:else}
            <div class="flex gap-4 items-center">
              <div class="text-sm text-gray-500 dark:text-gray-400 dark:bg-gray-800"><span class="font-semibold">Selected account:</span> {$selectedAccount}</div>
              <div class="text-sm text-gray-500 dark:text-gray-400 dark:bg-gray-800"><span class="font-semibold">Network ID:</span> {$chainId}</div>
              {#if !payouts.find(payout => payout.selected)}
              <div class="text-sm text-gray-500 dark:text-gray-400 dark:bg-gray-800">⚠️ There are no selected payouts.</div>      
              {:else}
                {#if payouts.find(payout => payout.selected && payout.verified === undefined)}
                  <button on:click={verifySelectedPayouts} type="button" class="focus:outline-none text-white bg-purple-700 hover:bg-purple-800 focus:ring-4 focus:ring-purple-300 font-medium rounded-lg text-sm px-5 py-2.5 mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900">Verify selected payouts</button>   
                {/if}
                {#if payouts.find(payout => payout.selected && payout.verified !== undefined)}
                  <input type="text" bind:value={gasPrice} class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-40 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Gas price">
                  <button on:click={sendPayouts} type="button" class="focus:outline-none text-white bg-purple-700 hover:bg-purple-800 focus:ring-4 focus:ring-purple-300 font-medium rounded-lg text-sm px-5 py-2.5 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900">Pay out</button>   
                {/if}

                

              {/if}
            </div>
          {/if}
        </div>

      
        <div class="mt-4 overflow-x-auto shadow-md sm:rounded-lg">
          {#if !selectedBlockchainID}
            <div class="p-4 text-sm text-gray-700 bg-gray-100 rounded-lg dark:bg-gray-700 dark:text-gray-300" role="alert">
              <span class="font-medium">⚠️ Please, select your blockchain</span>
            </div>
          {:else}
          <table class="text-sm text-left text-gray-500 dark:text-gray-400" style="table-layout:fixed;">
            <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
              <tr>
                <th scope="col" class="p-4"><Checkbox on:change={handleCheckAll} /></th>
                <th scope="col" class="p-4 pl-0">Offer ID</th>
                <th scope="col" class="p-4 pl-0">Paid out</th>
                <th scope="col" class="p-4 pl-0">Total reward</th>
                <th scope="col" class="p-4 pl-0">Estimated payout</th>
                <th scope="col" class="p-4 pl-0 whitespace-nowrap">Last payout</th>
                <th scope="col" class="p-4 pl-0 whitespace-nowrap">Verified</th>
              </tr>
            </thead>
            {#if !selectedBlockchainID || loadingPayouts || !payouts.length}
              <tbody>
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                  <td colspan="7" class="w-4 p-4"> 
                    {#if loadingPayouts}<Spinner />
                    {:else if !payouts.length}⚠️ No payouts were found within the selected blockchain
                    {/if}
                  </td>
                </tr>
              </tbody>
            {:else}
              <tbody>
                {#each payouts as payout, i}
                  <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                    <td class="w-4 p-4">
                      <Checkbox bind:checked={payout.selected} />
                    </td>
                    <td scope="row" class="p-4 pl-0 text-sm w-10 truncate">
                      <span title="{payout.OfferID}">{payout.OfferID.substring(0, 10)}…</span>
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

                    <td class="p-4 pl-0 text-sm whitespace-nowrap">
                      {#if payout.verified === undefined} 
                        <span class="bg-gray-100 text-gray-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">Unverified</span>
                      {:else if payout.verified} 
                        <span class="bg-green-100 text-green-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-green-200 dark:text-green-900">Verified</span>
                      {:else}
                        <span class="bg-red-100 text-red-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-red-200 dark:text-red-900">Failed</span>
                      {/if}
                    </td>
                  </tr>
                {/each}
              </tbody>
            {/if}
        </table>
      {/if}
    </div>
  </main>
{/if}