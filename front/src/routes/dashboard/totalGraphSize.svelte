<script>
  import Badge from './components/badge.svelte'
</script>

<main class='h-full overflow-y-auto'>
  <div class='container px-6 mx-auto grid'>
    <div class="mt-6">
      <Badge />
    </div>
    <h2 class='my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200'>
      Total Graph Size
    </h2>
    <div class='aspect-video p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
      <iframe width="100%" height="100%" src="https://app.powerbi.com/view?r=eyJrIjoiNzVlNzk3YTctYjE4MS00MGU1LTg3YWQtMWJjODE4MzY5ZjEwIiwidCI6IjJiY2VkZmNiLTQwMzAtNDk3Mi1hZDE2LWZjZjBhMmIwNGEyOCIsImMiOjl9&amp;pageName=ReportSection" frameborder="0" allowfullscreen="true"></iframe> 
    </div>
  </div>
</main>