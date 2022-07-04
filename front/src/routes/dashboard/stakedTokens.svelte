<script>
  import { onMount } from 'svelte'
  import Badge from './components/badge.svelte'
  import * as API from '$modules/api.module.js'
  import Chart from 'chart.js/auto/auto.js'

  let stakedTokensChartCanvas
  let stakedTokensChartCanvasContext
  let depositsAndWithdrawalsChartCanvas 
  let depositsAndWithdrawalsChartCanvasContext

  let stakedTokens

  onMount(async () => {
    stakedTokensChartCanvasContext = stakedTokensChartCanvas.getContext('2d')
    depositsAndWithdrawalsChartCanvasContext = depositsAndWithdrawalsChartCanvas.getContext('2d')
    const stakedTokensRequest = await API.getStakedTokens()
    stakedTokens = stakedTokensRequest.data
    drawStakedTokensChart()
    drawDepositsAndWithdrawalsChartCanvas()
  })

  const drawStakedTokensChart = () => {
    const labels = stakedTokens.map(stakedToken => new Date(stakedToken.Date).toLocaleString('en-US', { month: 'short', year: 'numeric' }))
    const stakedTokensArray = stakedTokens.map(stakedToken => stakedToken.Staked)
    const chart = new Chart(stakedTokensChartCanvasContext, {
      type: 'line',
      options: {
        responsive: true,
        maintainAspectRatio: false,
      },
      data: {
        labels: labels,
        datasets: [{
          label: 'Staked',
          data: stakedTokensArray,
          fill: false,
          borderColor: 'rgb(4, 116, 129)',
          tension: 0.1
        }
        ]
      }
    })
  }

  const drawDepositsAndWithdrawalsChartCanvas = () => {
    const labels = stakedTokens.map(stakedToken => new Date(stakedToken.Date).toLocaleString('en-US', { month: 'short', year: 'numeric' }))
    const depositsArray = stakedTokens.map(stakedToken => stakedToken.Deposited)
    const withdrawalsArray = stakedTokens.map(stakedToken => stakedToken.Withdrawn)
    const chart = new Chart(depositsAndWithdrawalsChartCanvasContext, {
      type: 'line',
      options: {
        responsive: true,
        maintainAspectRatio: false,
      },
      data: {
        labels: labels,
        datasets: [{
          label: 'Deposited',
          data: depositsArray,
          fill: false,
          borderColor: 'rgb(4, 116, 129)',
          tension: 0.1
        },
        {
          label: 'Withdrawn',
          data: withdrawalsArray,
          fill: false,
          borderColor: 'rgb(126, 58, 242)',
          tension: 0.1
        }
        ]
      }
    })
  }
</script>

<main class='h-full overflow-y-auto'>
  <div class='container px-6 mx-auto grid'>
    <div class="mt-6">
      <Badge />
    </div>
    <h2 class='my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200'>Staked tokens</h2>
    <div class='h-96 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800 relative'>
      <canvas class="h-full w-full" bind:this={stakedTokensChartCanvas} />
    </div>
    <h2 class='my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200'>Deposits & withdrawals</h2>
    <div class='h-96 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800 relative'>
      <canvas class="h-full w-full" bind:this={depositsAndWithdrawalsChartCanvas} />
    </div>
  </div>
</main>