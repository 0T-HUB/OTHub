<script>
  export let change
  export let checked

  const timestamp = Date.now()
</script>


<div class="flex gap-2 items-center">
  <div class="w-4 h-4 rounded flex items-center overflow-hidden dark:bg-gray-600 dark:border-gray-500">
    <input on:change bind:checked id={timestamp} type="checkbox" class="peer absolute opacity-0">
    <label for={timestamp} class="block bg-purple-500 opacity-0 peer-checked:opacity-100">
      <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 20 20" fill="white">
        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
      </svg>
    </label>
  </div>
  <span class="text-sm text-gray-500 dark:text-gray-400">
    <slot></slot>
  </span>
</div>