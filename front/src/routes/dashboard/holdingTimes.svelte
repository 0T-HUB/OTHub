<script>
  import { onMount } from 'svelte'
  import Badge from './components/badge.svelte'
  import * as API from '$modules/api.module.js'
  import Chart from 'chart.js/auto/auto.js'

  const monthNamesByMonthNumberMap = new Map([[1, 'Jan'],[2, 'Feb'],[3, 'Mar'],[4, 'Apr'],[5, 'May'],[6, 'Jun'],[7, 'Jul'],[8, 'Aug'],[9, 'Sep'],[10, 'Oct'],[11, 'Nov'],[12, 'Dec']])

  let barChartCanvas
  let barChartContext
  let holdingTimes

  onMount(async () => {
    barChartContext = barChartCanvas.getContext('2d')
    const holdingTimesFromAPIRequest = await API.getHoldingTimes()
    let holdingTimesFromAPI = holdingTimesFromAPIRequest.data
    holdingTimesFromAPI.Data = holdingTimesFromAPI.Data.map(holdingTime => ({ ...holdingTime, label: `${monthNamesByMonthNumberMap.get(holdingTime.Month)} ’${holdingTime.Year.slice(-2)}`}))
    holdingTimes = holdingTimesFromAPI
    drawBarChart()
  })

  const onlyUnique = (value, index, self) => self.indexOf(value) === index

  const drawBarChart = () => {
    const monthYearLabels = holdingTimes.Data.map(holdingTime => holdingTime.label).filter(onlyUnique)
    const { HoldingTimesAvailable: arrayOfPossibleHoldingTimeValues } = holdingTimes
    const backgroundColorByHoldingTimeValueMap = new Map([[0, '#f05252'],[1, '#3f83f8'],[3, '#e74694'],[4, '#0e9f6e'],[5, '#ff5a1f'],[6, '#6875f5'],[12, '#c27803'],[24, '#0694a2'],[36, '#9061f9'],[54, '#8b5cf6'],[60, '#84cc16']])
    const datasets = arrayOfPossibleHoldingTimeValues.map(holdingTimeValue => {
      const holdingTimesWithSameTimespan = holdingTimes.Data.filter(holdingTime => holdingTime.HoldingTimeInMonths === holdingTimeValue)
      const monthlyHoldingTimeData = monthYearLabels.map(monthYearLabel => holdingTimesWithSameTimespan.find(holdingTime => holdingTime.label === monthYearLabel)?.Count || 0)
      return {
        label: `${holdingTimeValue} months`,
        data: monthlyHoldingTimeData,
        backgroundColor: backgroundColorByHoldingTimeValueMap.get(holdingTimeValue)
      }
    })
    const chart = new Chart(barChartContext, {
      type: 'bar',
      data: {
        labels: monthYearLabels,
        datasets
      },
      options: {
        responsive: true,
        scales: {
          x: {
            stacked: true,
          },
          y: {
            stacked: true
          }
        }
      }
    })
  }
</script>

<main class='h-full overflow-y-auto'>
  <div class='container px-6 mx-auto grid'>

    <div class="mt-6">
      <Badge />
    </div>

    <h2 class='my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200'>
      Job holding times
    </h2>

    <div class='aspect-video p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800 relative'>
      <canvas class="h-full w-full" bind:this={barChartCanvas} />
    </div>

  </div>
</main>