<!--Error page 401-->
<script>
  import { onMount } from 'svelte'

  let auth0Client
  onMount(async () => auth0Client = await auth.createClient())

  import auth from '../../authService.js'
  const login = () => auth.loginWithPopup(auth0Client)

</script>

<main class="h-full pb-16 overflow-y-auto flex items-center">
  <div class="container flex flex-col items-center px-6 mx-auto flex gap-3 flex-col">
    <svg class="w-12 h-12 mt-8 text-purple-200" fill="currentColor" viewBox="0 0 20 20">
      <path
        fill-rule="evenodd"
        d="M13.477 14.89A6 6 0 015.11 6.524l8.367 8.368zm1.414-1.414L6.524 5.11a6 6 0 018.367 8.367zM18 10a8 8 0 11-16 0 8 8 0 0116 0z"
        clip-rule="evenodd"
      />
    </svg>
    <h1 class="text-6xl font-semibold text-gray-700 dark:text-gray-200">401</h1>
    <p class="text-gray-700 dark:text-gray-300">
      You are not authorized. Please, <button class="text-purple-600 hover:underline dark:text-purple-300" on:click={login}>login</button> to view this page.
    </p>
  </div>
</main>