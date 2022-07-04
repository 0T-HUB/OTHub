<script>
  import { onMount } from 'svelte'
  import Chart from 'chart.js/auto/auto.js'
  import Badge from './components/badge.svelte'
  import Spinner from './components/spinner.svelte'

  import * as API from '$modules/api.module'
  
  let badge
  let dashboardData
  let blockchainDistribution
  let jobs
  let gaugeContext
  let gaugeCanvas
  let lineChartContext
  let lineChartCanvas
  

  onMount(async () => {
    gaugeContext = gaugeCanvas.getContext('2d')
    lineChartContext = lineChartCanvas.getContext('2d')
    getAndSetDashboardData()
    await getAndSetBlockchainDistribution()
    drawMeter()
    await getAndSetJobs()
    drawLineChart()
  })

  const drawLineChart = () => {
    const labels = jobs.map(job => new Date(job.Date).toLocaleString('en-US', { month: 'short', day: 'numeric' }))
    const newJobs = jobs.map(job => job.NewJobs)
    const completedJobs = jobs.map(job => job.CompletedJobs)
    const chart = new Chart(lineChartContext, {
      type: 'line',
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            display: false
          }
        }
      },
      data: {
        labels: labels,
        datasets: [{
          label: 'Started',
          data: newJobs,
          fill: false,
          borderColor: 'rgb(4, 116, 129)',
          tension: 0.1
        },
        {
          label: 'Completed',
          data: completedJobs,
          fill: false,
          borderColor: 'rgb(126, 58, 242)',
          tension: 0.1
        }]
      }
    })
  }

  const drawMeter = () => {
    const data = blockchainDistribution.Blockchains.map(blockchain => blockchain.Jobs)
    const backgroundColor = blockchainDistribution.Blockchains.map(blockchain => `#${blockchain.Color}`)
    const chart = new Chart(gaugeContext, {
      type: 'doughnut',
      data: {
        datasets: [
          {
            data,
            backgroundColor,
            borderWidth: 0
          }
        ]
      },
      options: {
        cutout: '75%',
        maintainAspectRatio: false,
        radius: '75%',
        rotation: -90,
        circumference: 120
      }
    })
  }

  const getAndSetJobs =  async () => {
    const jobsRequest = await API.getJobs()
    jobs = jobsRequest.data
  }


  const getAndSetDashboardData = async () => {
    const dashboardDataRequest = await API.getDashboardData()
    dashboardData = dashboardDataRequest.data
  }

  const getAndSetBlockchainDistribution = async () => {
    const blockchainDistributionRequest = await API.getBlockchainDistribution()
    blockchainDistribution = blockchainDistributionRequest.data
  }
</script>

<svelte:head>
  <title>OND</title>
</svelte:head>

<main class='h-full overflow-y-auto'>
  <div class='container px-6 mx-auto grid'>

    <div class="mt-6">
      <Badge />
    </div>

    <h2 class='my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200'>
      Origintrail Network Dashboard
    </h2>
    <!-- Cards -->

    <div class="grid grid-cols-2 gap-6 mb-6">
      <div class='grid gap-6 grid-cols-2'>
        <!-- Trac price -->
        <div class='flex items-center p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
          <div class='p-3 mr-4 text-orange-500 bg-orange-100 rounded-full dark:text-orange-100 dark:bg-orange-500'>
            <svg class='w-5 h-5' fill='currentColor' viewBox='0 0 20 20'>
              <path d='M12.173,16.086c0.72,0,1.304-0.584,1.304-1.305V6.089c0-0.72-0.584-1.304-1.304-1.304c-0.721,0-1.305,0.584-1.305,1.304v8.692C10.868,15.502,11.452,16.086,12.173,16.086z M11.738,6.089c0-0.24,0.194-0.435,0.435-0.435s0.435,0.194,0.435,0.435v8.692c0,0.24-0.194,0.436-0.435,0.436s-0.435-0.195-0.435-0.436V6.089zM16.52,16.086c0.72,0,1.304-0.584,1.304-1.305v-11.3c0-0.72-0.584-1.304-1.304-1.304c-0.721,0-1.305,0.584-1.305,1.304v11.3C15.215,15.502,15.799,16.086,16.52,16.086z M16.085,3.481c0-0.24,0.194-0.435,0.435-0.435s0.435,0.195,0.435,0.435v11.3c0,0.24-0.194,0.436-0.435,0.436s-0.435-0.195-0.435-0.436V3.481z M3.48,16.086c0.72,0,1.304-0.584,1.304-1.305v-3.477c0-0.72-0.584-1.304-1.304-1.304c-0.72,0-1.304,0.584-1.304,1.304v3.477C2.176,15.502,2.76,16.086,3.48,16.086z M3.045,11.305c0-0.24,0.194-0.435,0.435-0.435c0.24,0,0.435,0.194,0.435,0.435v3.477c0,0.24-0.195,0.436-0.435,0.436c-0.24,0-0.435-0.195-0.435-0.436V11.305z M18.258,16.955H1.741c-0.24,0-0.435,0.194-0.435,0.435s0.194,0.435,0.435,0.435h16.517c0.24,0,0.435-0.194,0.435-0.435S18.498,16.955,18.258,16.955z M7.826,16.086c0.72,0,1.304-0.584,1.304-1.305V8.696c0-0.72-0.584-1.304-1.304-1.304S6.522,7.977,6.522,8.696v6.085C6.522,15.502,7.106,16.086,7.826,16.086z M7.392,8.696c0-0.239,0.194-0.435,0.435-0.435s0.435,0.195,0.435,0.435v6.085c0,0.24-0.194,0.436-0.435,0.436s-0.435-0.195-0.435-0.436V8.696z'/>
            </svg>
          </div>
          <div>
            <p class='mb-2 text-sm font-medium text-gray-600 dark:text-gray-400'>TRAC price</p>
            {#if dashboardData}
              <div class="flex items-center">
                <span class='text-lg font-semibold text-gray-700 dark:text-gray-200'>${dashboardData.PriceUsd.toFixed(4)}</span>
                <span 
                  class="bg-green-100 text-green-800 text-xs font-semibold mr-2 px-2.5 py-0.5 rounded dark:bg-green-200 dark:text-green-900 ml-2" 
                  class:bg-red-100={(dashboardData.PercentChange24H < 0)}
                  class:text-red-800={(dashboardData.PercentChange24H < 0)}
                  class:dark:text-red-900={(dashboardData.PercentChange24H < 0)}
                  class:dark:bg-red-200={(dashboardData.PercentChange24H < 0)}>{dashboardData.PercentChange24H}%</span>
              </div>
            {:else}
              <Spinner class="h-7" />
            {/if}
          </div>
        </div>

        <!-- Card -->
        <div class='flex items-center p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
          <div
            class='p-3 mr-4 text-green-500 bg-green-100 rounded-full dark:text-green-100 dark:bg-green-500'
          >
            <svg class='w-5 h-5' fill='currentColor' viewBox='0 0 20 20'>
              <path
                fill-rule='evenodd'
                d='M4 4a2 2 0 00-2 2v4a2 2 0 002 2V6h10a2 2 0 00-2-2H4zm2 6a2 2 0 012-2h8a2 2 0 012 2v4a2 2 0 01-2 2H8a2 2 0 01-2-2v-4zm6 4a2 2 0 100-4 2 2 0 000 4z'
                clip-rule='evenodd'
              />
            </svg>
          </div>
          <div>
            <p class='mb-2 text-sm font-medium text-gray-600 dark:text-gray-400'>Market Cap</p>
            {#if dashboardData}
              <p class='text-lg font-semibold text-gray-700 dark:text-gray-200'>${dashboardData.MarketCapUsd.toLocaleString('en-US')}</p>
            {:else}
              <Spinner class="h-7" />
            {/if}
          </div>
        </div>
        <!-- Card -->
        <div class='flex items-center p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
          <div
            class='p-3 mr-4 text-orange-500 bg-orange-100 rounded-full dark:text-orange-100 dark:bg-orange-500'
          >
            <svg class='w-5 h-5' fill='currentColor' viewBox='0 0 20 20'>
              <path
                d='M12.173,16.086c0.72,0,1.304-0.584,1.304-1.305V6.089c0-0.72-0.584-1.304-1.304-1.304c-0.721,0-1.305,0.584-1.305,1.304v8.692C10.868,15.502,11.452,16.086,12.173,16.086z M11.738,6.089c0-0.24,0.194-0.435,0.435-0.435s0.435,0.194,0.435,0.435v8.692c0,0.24-0.194,0.436-0.435,0.436s-0.435-0.195-0.435-0.436V6.089zM16.52,16.086c0.72,0,1.304-0.584,1.304-1.305v-11.3c0-0.72-0.584-1.304-1.304-1.304c-0.721,0-1.305,0.584-1.305,1.304v11.3C15.215,15.502,15.799,16.086,16.52,16.086z M16.085,3.481c0-0.24,0.194-0.435,0.435-0.435s0.435,0.195,0.435,0.435v11.3c0,0.24-0.194,0.436-0.435,0.436s-0.435-0.195-0.435-0.436V3.481z M3.48,16.086c0.72,0,1.304-0.584,1.304-1.305v-3.477c0-0.72-0.584-1.304-1.304-1.304c-0.72,0-1.304,0.584-1.304,1.304v3.477C2.176,15.502,2.76,16.086,3.48,16.086z M3.045,11.305c0-0.24,0.194-0.435,0.435-0.435c0.24,0,0.435,0.194,0.435,0.435v3.477c0,0.24-0.195,0.436-0.435,0.436c-0.24,0-0.435-0.195-0.435-0.436V11.305z M18.258,16.955H1.741c-0.24,0-0.435,0.194-0.435,0.435s0.194,0.435,0.435,0.435h16.517c0.24,0,0.435-0.194,0.435-0.435S18.498,16.955,18.258,16.955z M7.826,16.086c0.72,0,1.304-0.584,1.304-1.305V8.696c0-0.72-0.584-1.304-1.304-1.304S6.522,7.977,6.522,8.696v6.085C6.522,15.502,7.106,16.086,7.826,16.086z M7.392,8.696c0-0.239,0.194-0.435,0.435-0.435s0.435,0.195,0.435,0.435v6.085c0,0.24-0.194,0.436-0.435,0.436s-0.435-0.195-0.435-0.436V8.696z'
              />
            </svg>
          </div>
          <div>
            <p class='mb-2 text-sm font-medium text-gray-600 dark:text-gray-400'>Total jobs</p>
            {#if dashboardData}
              <p class='text-lg font-semibold text-gray-700 dark:text-gray-200'>{dashboardData.All.TotalJobs.toLocaleString('en-US')}</p>
            {:else}
              <Spinner class="h-7" />
            {/if}
          </div>
        </div>
        <!-- Card -->
        <div class='flex items-center p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
          <div
            class='p-3 mr-4 text-blue-500 bg-blue-100 rounded-full dark:text-blue-100 dark:bg-blue-500'
          >
            <svg class='w-5 h-5' fill='currentColor' viewBox='0 0 20 20'>
              <path
                d='M4.68,13.716v-0.169H4.554C4.592,13.605,4.639,13.658,4.68,13.716z M11.931,6.465
    c-0.307-0.087-0.623,0.106-0.706,0.432l-1.389,5.484c-0.901,0.084-1.609,0.833-1.609,1.757c0,0.979,0.793,1.773,1.773,1.773
    c0.979,0,1.773-0.794,1.773-1.773c0-0.624-0.324-1.171-0.812-1.486l1.377-5.439C12.422,6.887,12.239,6.552,11.931,6.465z
    M10.591,14.729H9.408v-1.182h1.183V14.729z M15.32,13.716c0.04-0.058,0.087-0.11,0.126-0.169H15.32V13.716z M10,3.497
    c-3.592,0-6.503,2.911-6.503,6.503H4.68c0-2.938,2.382-5.32,5.32-5.32s5.32,2.382,5.32,5.32h1.182
    C16.502,6.408,13.591,3.497,10,3.497z M10,0.542c-5.224,0-9.458,4.234-9.458,9.458c0,5.224,4.234,9.458,9.458,9.458
    c5.224,0,9.458-4.234,9.458-9.458C19.458,4.776,15.224,0.542,10,0.542z M15.32,16.335v0.167h-0.212
    c-1.407,1.107-3.179,1.773-5.108,1.773c-1.93,0-3.701-0.666-5.108-1.773H4.68v-0.167C2.874,14.816,1.724,12.543,1.724,10
    c0-4.571,3.706-8.276,8.276-8.276c4.57,0,8.275,3.706,8.275,8.276C18.275,12.543,17.126,14.816,15.32,16.335z'
              />
            </svg>
          </div>
          
          
          <div>
            <p class='mb-2 text-sm font-medium text-gray-600 dark:text-gray-400'>Jobs</p>
            {#if dashboardData}
              <p class='text-lg font-semibold text-gray-700 dark:text-gray-200'>{dashboardData.All.ActiveJobs.toLocaleString('en-US')}</p>
            {:else}
              <Spinner class="h-7" />
            {/if}
          </div>
        </div>
        <!-- Card -->
        <div class='flex items-center p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
          <div
            class='p-3 mr-4 text-teal-500 bg-teal-100 rounded-full dark:text-teal-100 dark:bg-teal-500'
          >
            <svg class='w-5 h-5' fill='currentColor' viewBox='0 0 20 20'>
              <path
                fill-rule='evenodd'
                d='M15.684,16.959L10.879,8.52c0.886-0.343,1.517-1.193,1.517-2.186c0-1.296-1.076-2.323-2.396-2.323S7.604,5.037,7.604,6.333c0,0.993,0.63,1.843,1.517,2.186l-4.818,8.439c-0.189,0.311,0.038,0.708,0.412,0.708h10.558C15.645,17.667,15.871,17.27,15.684,16.959 M8.562,6.333c0-0.778,0.645-1.382,1.438-1.382s1.438,0.604,1.438,1.382c0,0.779-0.645,1.412-1.438,1.412S8.562,7.113,8.562,6.333 M5.55,16.726L10,8.91l4.435,7.815H5.55z M15.285,9.62c1.26-2.046,1.26-4.525,0-6.572c-0.138-0.223-0.064-0.512,0.162-0.646c0.227-0.134,0.521-0.063,0.658,0.16c1.443,2.346,1.443,5.2,0,7.546c-0.236,0.382-0.641,0.17-0.658,0.159C15.221,10.131,15.147,9.842,15.285,9.62 M13.395,8.008c0.475-1.063,0.475-2.286,0-3.349c-0.106-0.238,0.004-0.515,0.246-0.62c0.242-0.104,0.525,0.004,0.632,0.242c0.583,1.305,0.583,2.801,0,4.106c-0.214,0.479-0.747,0.192-0.632,0.242C13.398,8.523,13.288,8.247,13.395,8.008 M3.895,10.107c-1.444-2.346-1.444-5.2,0-7.546c0.137-0.223,0.431-0.294,0.658-0.16c0.226,0.135,0.299,0.424,0.162,0.646c-1.26,2.047-1.26,4.525,0,6.572c0.137,0.223,0.064,0.512-0.162,0.646C4.535,10.277,4.131,10.489,3.895,10.107 M5.728,8.387c-0.583-1.305-0.583-2.801,0-4.106c0.106-0.238,0.39-0.346,0.631-0.242c0.242,0.105,0.353,0.382,0.247,0.62c-0.475,1.063-0.475,2.286,0,3.349c0.106,0.238-0.004,0.515-0.247,0.62c-0.062,0.027-0.128,0.04-0.192,0.04C5.982,8.668,5.807,8.563,5.728,8.387'
                clip-rule='evenodd'
              />
            </svg>
          </div>
          <div>
            <p class='mb-2 text-sm font-medium text-gray-600 dark:text-gray-400'>Locked tokens</p>
            {#if dashboardData}
              <span class='text-lg font-semibold text-gray-700 dark:text-gray-200'>{dashboardData.All.TokensLocked24H.toLocaleString('en-US')}</span>
              <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">TRAC</span>
            {:else}
              <Spinner class="h-7" />
            {/if}
          </div>
        </div>
        <!-- Card -->
        <div class='flex items-center p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
          <div
            class='p-3 mr-4 text-teal-500 bg-teal-100 rounded-full dark:text-teal-100 dark:bg-teal-500'
          >
            <svg class='w-5 h-5' fill='currentColor' viewBox='0 0 20 20'>
              <path
                fill-rule='evenodd'
                d='M15.684,16.959L10.879,8.52c0.886-0.343,1.517-1.193,1.517-2.186c0-1.296-1.076-2.323-2.396-2.323S7.604,5.037,7.604,6.333c0,0.993,0.63,1.843,1.517,2.186l-4.818,8.439c-0.189,0.311,0.038,0.708,0.412,0.708h10.558C15.645,17.667,15.871,17.27,15.684,16.959 M8.562,6.333c0-0.778,0.645-1.382,1.438-1.382s1.438,0.604,1.438,1.382c0,0.779-0.645,1.412-1.438,1.412S8.562,7.113,8.562,6.333 M5.55,16.726L10,8.91l4.435,7.815H5.55z M15.285,9.62c1.26-2.046,1.26-4.525,0-6.572c-0.138-0.223-0.064-0.512,0.162-0.646c0.227-0.134,0.521-0.063,0.658,0.16c1.443,2.346,1.443,5.2,0,7.546c-0.236,0.382-0.641,0.17-0.658,0.159C15.221,10.131,15.147,9.842,15.285,9.62 M13.395,8.008c0.475-1.063,0.475-2.286,0-3.349c-0.106-0.238,0.004-0.515,0.246-0.62c0.242-0.104,0.525,0.004,0.632,0.242c0.583,1.305,0.583,2.801,0,4.106c-0.214,0.479-0.747,0.192-0.632,0.242C13.398,8.523,13.288,8.247,13.395,8.008 M3.895,10.107c-1.444-2.346-1.444-5.2,0-7.546c0.137-0.223,0.431-0.294,0.658-0.16c0.226,0.135,0.299,0.424,0.162,0.646c-1.26,2.047-1.26,4.525,0,6.572c0.137,0.223,0.064,0.512-0.162,0.646C4.535,10.277,4.131,10.489,3.895,10.107 M5.728,8.387c-0.583-1.305-0.583-2.801,0-4.106c0.106-0.238,0.39-0.346,0.631-0.242c0.242,0.105,0.353,0.382,0.247,0.62c-0.475,1.063-0.475,2.286,0,3.349c0.106,0.238-0.004,0.515-0.247,0.62c-0.062,0.027-0.128,0.04-0.192,0.04C5.982,8.668,5.807,8.563,5.728,8.387'
                clip-rule='evenodd'
              />
            </svg>
          </div>
          <div>
            <p class='mb-2 text-sm font-medium text-gray-600 dark:text-gray-400'>Paidout tokens</p>
            {#if dashboardData}
              <span class='text-lg font-semibold text-gray-700 dark:text-gray-200'>{dashboardData.All.TokensPaidout24H.toLocaleString('en-US')}</span>
              <span class="bg-gray-100 text-gray-800 text-xs font-medium ml-1 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-300">TRAC</span>
            {:else}
              <Spinner class="h-7" />
            {/if}
          </div>
        </div>
      </div>

      <div class='min-w-0 p-4 h-96 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
        <h4 class='mb-4 font-semibold text-gray-800 dark:text-gray-300'>24h Job-O-Meter</h4>
        {#if !blockchainDistribution}
          <div class="w-full h-full flex items-center justify-center">
            <Spinner />
          </div>
        {/if}
        <div class="w-full h-60 relative flex items-center justify-center" class:hidden={!blockchainDistribution}>
          <canvas bind:this={gaugeCanvas} />
          <div class="absolute w-0.5 h-48 bg-black top-6 origin-bottom rotate-[30deg] translate-x-10 opacity-0 transition-opacity"
          class:opacity-100={blockchainDistribution}></div>
        </div>
        <div class='flex justify-center mt-4 space-x-3 text-sm text-gray-600 dark:text-gray-400' class:hidden={!blockchainDistribution}>
          <div class='flex items-center'>
            <span class='inline-block w-3 h-3 mr-1 bg-blue-500 rounded-full' />
            <span>Ethereum</span>
          </div>
          <div class='flex items-center'>
            <span class='inline-block w-3 h-3 mr-1 bg-teal-600 rounded-full' />
            <span>Polygon</span>
          </div>
          <div class='flex items-center'>
            <span class='inline-block w-3 h-3 mr-1 bg-purple-600 rounded-full' />
            <span>Gnosis</span>
          </div>
        </div>
      </div>
    </div>


    <div class='min-w-0 h-96 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
      <h4 class='mb-4 font-semibold text-gray-800 dark:text-gray-300'>Jobs</h4>
      
      {#if !jobs}
        <div class="w-full h-full flex items-center justify-center">
          <Spinner />
        </div>
      {/if}
      
      <div class="h-60 w-full">
        <canvas bind:this={lineChartCanvas} class="h-full w-full" class:hidden={!jobs}/>
      </div>

      <div class='flex justify-center mt-4 space-x-3 text-sm text-gray-600 dark:text-gray-400' class:hidden={!jobs}>
        <div class='flex items-center'>
          <span class='inline-block w-3 h-3 mr-1 bg-teal-600 rounded-full' />
          <span>Started</span>
        </div>
        <div class='flex items-center'>
          <span class='inline-block w-3 h-3 mr-1 bg-purple-600 rounded-full' />
          <span>Completed</span>
        </div>
      </div>
    </div>
  </div>
</main>
