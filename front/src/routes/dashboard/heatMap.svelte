<script>
  import { onMount } from 'svelte'
  import Badge from './components/badge.svelte'
  import * as API from '$modules/api.module.js'
  import { Chart, CategoryScale, Tooltip } from 'chart.js'
  import { MatrixController, MatrixElement } from 'chartjs-chart-matrix'
  

  let matrixChartCanvas
  let matrixChartContext
  let jobHeatmap

  onMount(async () => {
    Chart.register(CategoryScale, Tooltip, MatrixController, MatrixElement)
    matrixChartContext = matrixChartCanvas.getContext('2d')
    const jobHeatmapRequest = await API.getJobHeatmap()
    jobHeatmap = jobHeatmapRequest.data
    drawMatrixChart()
  })

  const onlyUnique = (value, index, self) => self.indexOf(value) === index

  const drawMatrixChart = () => {
    const daysOfWeek = jobHeatmap.map(item => item.Day).filter(onlyUnique)
    const hoursOfDay = jobHeatmap.map(item => item.HourText).filter(onlyUnique)
    const data = jobHeatmap.map(item => ({ x: item.Day, y: item.HourText, v: item.Count }))
    const chart = new Chart(matrixChartContext, {
      type: 'matrix',
      data: {
        datasets: [{
          label: 'Jobs',
          data: data,
          backgroundColor(context) {
            const value = context.dataset.data[context.dataIndex].v
            const alpha = (value - 5) / 40
            return `rgb(168, 85, 247, ${alpha})`
          },
          borderColor(context) {
            const value = context.dataset.data[context.dataIndex].v
            const alpha = (value - 5) / 40
            return `rgb(107, 33, 168, ${alpha})`
          },
          borderWidth: 0,
          width: ({chart}) => (chart.chartArea || {}).width / 7 - 1,
          height: ({chart}) =>(chart.chartArea || {}).height / 23 - 1
        }],
      },
      options: {
        responsive: true,
        plugins: {
          tooltip: {
            callbacks: {
              title() {
                return '';
              },
              label(context) {
                const v = context.dataset.data[context.dataIndex];
                return [`${v.v} jobs on ${v.x} at ${v.y}`]
              }
            }
          },
        },
        scales: {
          x: {
            type: 'category',
            ticks: {
              display: true
            },
            grid: {
              display: false
            },
            labels: daysOfWeek,
          },
          y: {
            type: 'category',
            offset: true,
            ticks: {
              display: true
            },
            grid: {
              display: false
            },
            labels: hoursOfDay
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
    <h2 class='my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200'>Job heatmap</h2>
    <div class='aspect-video p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
      <canvas class="h-full w-full" bind:this={matrixChartCanvas} />
    </div>
  </div>
</main>