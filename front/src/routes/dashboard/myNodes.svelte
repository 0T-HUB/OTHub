<script>
  import cookie from 'cookie'
  import { onMount } from 'svelte'
  import Unauthorized from './unauthorized.svelte'
  import { isAuthenticated } from '$stores/auth' 
  import * as API from '$modules/api.module'
  import Chart from 'chart.js/auto/auto.js'
  import Badge from './components/badge.svelte'

  let recentJobs
  let jobsPerMonth
  let holdingTime
  let nodeStats
  let context
  let canvas
  let container

  onMount(async () => {
    canvas = document.createElement('canvas')
    const recentJobsRequest = await API.getRecentJobsByNodeID()
    recentJobs = recentJobsRequest.data
    const jobsPerMonthRequest = await API.getJobsPerMonth()
    jobsPerMonth = jobsPerMonthRequest.data
    const nodeStatsRequest = await API.getNodeStats()
    nodeStats = nodeStatsRequest.data
    const holdingTimeRequest = await API.getHoldingTimeByMonth()
    holdingTime = holdingTimeRequest.data
    updateHoldingTimeChart()
  })

  const updateHoldingTimeChart = () => {
    const labels = holdingTime.map(timeSpan => (`${timeSpan.HoldingTimeInMonths} months`))
    const data = holdingTime.map(timeSpan => timeSpan.Count)
    context = canvas.getContext('2d')
    const chart = new Chart(context, {
      type: 'pie',
      options: {
        elements: {
            arc: {
                borderWidth: 0
            }
        }
      },
      data: {
        labels,
        datasets: [
          {
            label: 'Holding time',
            data,
            backgroundColor: [
              'rgb(133, 105, 241)',
              'rgb(164, 101, 241)',
              'rgb(101, 143, 241)',
              'rgb(88, 101, 241)',
              'rgb(64, 98, 241)',
            ],
            hoverOffset: 4,
          },
        ],
      }
    })
    container.appendChild(canvas)
  }


  const activateTabByRecentJobIndex = (recentJobIndexToBeSetActive) => {
    recentJobs = recentJobs.map((recentJob, recentJobIndex) => {
      return {
        ...recentJob,
        Active: recentJobIndex === recentJobIndexToBeSetActive ? true : false
      }
    })
  }

  const activateYearTabByYearIndex = (yearIndexToBeSetActive) => {
    jobsPerMonth.AllNodes.Years = jobsPerMonth.AllNodes.Years.map((year, yearIndex) => {
      return {
        ...year,
        Active: yearIndex === yearIndexToBeSetActive ? true : false
      }
    })
  }


</script>

{#if $isAuthenticated === undefined}
  <main class="h-full flex items-center justify-center">
    <svg role="status" class="inline w-6 h-6 mr-2 text-gray-200 animate-spin dark:text-gray-600 fill-purple-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
      <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
    </svg>
  </main>
{:else if $isAuthenticated === false}
  <Unauthorized />
{/if}

  <main class="h-full pb-16 pt-16 overflow-y-auto" style="{$isAuthenticated ? '' : 'display: none;'}">
    <div class="container flex flex-col gap-6 px-6 mx-auto">
      <Badge />
      
      <h2 class="my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200">My Nodes</h2>

      <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6">
        <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Latest jobs</h4>
        {#if recentJobs === undefined}
          <div class="flex h-96 w-full items-center justify-center">
            <div class="h-5 w-5 rounded-full border-2 border-t-purple-500 border-b-purple-500 border-l-purple-500 border-r-gray-600 animate-spin"></div>
            <span class="ml-2 text-gray-600">Loading...</span>
          </div>
        {:else if recentJobs === null}
          <div class="flex h-96 w-full items-center justify-center">
            <span class="text-gray-600">⚠️ There was an error fetching this data from the API</span>
          </div>
        {:else if recentJobs.length}
          <div class="text-sm font-medium text-center text-gray-500 border-b border-gray-200 dark:text-gray-400 dark:border-gray-700">
            <ul class="flex flex-wrap -mb-px">
              {#each recentJobs as recentJob, recentJobIndex}          
                <li class="mr-2">
                  <button class="inline-block p-4 rounded-t-lg border-b-2 border-transparent hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300" 
                  class:active={recentJob.Active}
                  class:dark:text-blue-500={recentJob.Active}
                  class:dark:border-blue-500={recentJob.Active}
                  on:click={() => activateTabByRecentJobIndex(recentJobIndex)}>
                  {recentJob.Day}
                  </button>
                </li>
              {/each}
            </ul>
          </div>
          {#each recentJobs as recentJob}
            {#if recentJob.Active }
              <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
                <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                  <tbody>
                    {#each recentJob.Jobs as job}
                      <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                        <td scope="row" class="px-6 pr-4 py-4 w-8 font-medium text-gray-900 dark:text-white whitespace-nowrap">
                          <img class="w-4 h-4 max-w-none" src="https://othub.origin-trail.network/assets/images/{job.BlockchainLogo}" alt="Blockchain {job.BlockchainLogo} logo.">
                        </td>
                        <th class="px-6 pl-0 py-4 w-6 whitespace-nowrap ">
                            {job.DisplayName}
                        </th>
                        <td class="px-6 py-4 pl-0">
                            Chosen for job at {job.FinalizedTimestamp.split('T')[1].slice(0, 5)}
                        </td>
                        <td class="px-6 py-4 text-right">
                            {job.TokenAmountPerHolder.toFixed(4)} <span class="text-xs font-semibold opacity-50">TRAC</span> / {job.USDAmount.toFixed(2)} <span class="text-xs font-semibold opacity-50">USD</span>
                        </td>
                      </tr>
                    {/each}
                  </tbody>
                </table>
              </div>
            {/if}
          {/each}
        {/if}
      </div>

      <div class="flex gap-4 items-start">
        <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6">
          <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Jobs per month</h4>
          {#if jobsPerMonth === undefined}
            <div class="flex h-96 w-full items-center justify-center">
              <div class="h-5 w-5 rounded-full border-2 border-t-purple-500 border-b-purple-500 border-l-purple-500 border-r-gray-600 animate-spin"></div>
              <span class="ml-2 text-gray-600">Loading...</span>
            </div>
          {:else if jobsPerMonth === null}
            <div class="flex h-96 w-full items-center justify-center">
              <span class="text-gray-600">⚠️ There was an error fetching this data from the API</span>
            </div>
          {:else if jobsPerMonth.AllNodes}
            <div class="text-sm font-medium text-center text-gray-500 border-b border-gray-200 dark:text-gray-400 dark:border-gray-700">
              <ul class="flex flex-wrap -mb-px">
                {#each jobsPerMonth.AllNodes.Years as year, yearIndex }          
                  <li class="mr-2">
                    <button class="inline-block p-4 rounded-t-lg border-b-2 border-transparent hover:text-gray-600 hover:border-gray-300 dark:hover:text-gray-300" 
                    class:active={year.Active}
                    class:dark:text-blue-500={year.Active}
                    class:dark:border-blue-500={year.Active}
                    on:click={() => activateYearTabByYearIndex(yearIndex)}>
                    {year.Year}
                    </button>
                  </li>
                {/each}
              </ul>
            </div>

            {#each jobsPerMonth.AllNodes.Years as year }
              {#if year.Active }
                <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
                  <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                    <tbody>
                      {#each year.Months as month}
                        <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                          <td scope="row" class="px-6 pr-1 py-4 w-8 font-medium text-gray-900 dark:text-white whitespace-nowrap">
                            {month.Month}
                          </td>
                          <th class="px-6 pl-0 py-4 w-6 whitespace-nowrap ">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor" class:-scale-100={month.Down} class:fill-red-600={month.Down} class:fill-green-600={!month.Down}>
                              <path fill-rule="evenodd" d="M5.293 7.707a1 1 0 010-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 01-1.414 1.414L11 5.414V17a1 1 0 11-2 0V5.414L6.707 7.707a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                            </svg>
                          </th>
                          <td class="px-6 pl-0 py-4">
                            {month.JobCount} jobs
                          </td>
                          <td class="px-6 py-4 text-right">
                            {month.TokenAmount.toFixed(4)} <span class="text-xs font-semibold opacity-50">TRAC</span> / {month.USDAmount.toFixed(2)} <span class="text-xs font-semibold opacity-50">USD</span>
                          </td>
                        </tr>
                      {/each}
                    </tbody>
                  </table>
                </div>
              {/if}
            {/each}
          {/if}
        </div>
        
        <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6 ">
          <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Total</h4>
          {#if nodeStats === undefined}
            <div class="flex h-96 w-full items-center justify-center">
              <div class="h-5 w-5 rounded-full border-2 border-t-purple-500 border-b-purple-500 border-l-purple-500 border-r-gray-600 animate-spin"></div>
              <span class="ml-2 text-gray-600">Loading...</span>
            </div>
          {:else if nodeStats === null}
            <div class="flex h-96 w-full items-center justify-center">
              <span class="text-gray-600">⚠️ There was an error fetching this data from the API</span>
            </div>
          {:else if nodeStats.TotalJobs.length}
            <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
              <tbody>
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                  <td scope="row" class="px-6 pr-1 py-4 w-8 font-medium text-gray-900 dark:text-white whitespace-nowrap">Jobs</td>
                  <td class="px-6 py-4 text-right">{nodeStats.TotalJobs.Value}</td>
                </tr>
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                  <td scope="row" class="px-6 pr-1 py-4 w-8 font-medium text-gray-900 dark:text-white whitespace-nowrap">Locked</td>
                  <td class="px-6 py-4 text-right">{nodeStats.TotalLocked.TokenAmount.toFixed(4)} <span class="text-xs font-semibold opacity-50">TRAC</span> / {nodeStats.TotalLocked.USDAmount.toFixed(2)} <span class="text-xs font-semibold opacity-50">USD</span></td>
                </tr>
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                  <td scope="row" class="px-6 pr-1 py-4 w-8 font-medium text-gray-900 dark:text-white whitespace-nowrap">Rewards</td>
                  <td class="px-6 py-4 text-right">{nodeStats.TotalRewards.TokenAmount.toFixed(4)} <span class="text-xs font-semibold opacity-50">TRAC</span> / {nodeStats.TotalRewards.USDAmount.toFixed(2)} <span class="text-xs font-semibold opacity-50">USD</span></td>
                </tr>
                <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                  <td scope="row" class="px-6 pr-1 py-4 w-8 font-medium text-gray-900 dark:text-white whitespace-nowrap">Staked</td>
                  <td class="px-6 py-4 text-right">{nodeStats.TotalStaked.TokenAmount.toFixed(4)} <span class="text-xs font-semibold opacity-50">TRAC</span> / {nodeStats.TotalStaked.USDAmount.toFixed(2)} <span class="text-xs font-semibold opacity-50">USD</span></td>
                </tr>
              </tbody>
            </table>
          {/if}
        </div>
        <div class="w-full dark:bg-gray-800 dark:border-gray-700 px-6 py-6 ">
          <h4 class="mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300">Holding Time on Active Jobs</h4>
          {#if holdingTime === undefined}
            <div class="flex h-96 w-full items-center justify-center">
              <div class="h-5 w-5 rounded-full border-2 border-t-purple-500 border-b-purple-500 border-l-purple-500 border-r-gray-600 animate-spin"></div>
              <span class="ml-2 text-gray-600">Loading...</span>
            </div>
          {/if}

          {#if holdingTime === null}
            <div class="flex h-96 w-full items-center justify-center">
              <span class="text-gray-600">⚠️ There was an error fetching this data from the API</span>
            </div>
          {/if}
          <div bind:this={container}/>
        </div>
      </div>
    </div>
  </main>
