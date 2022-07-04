<script lang="ts">
  import { onMount } from 'svelte'
  import * as API from '$modules/api.module'
  import { badge } from '$stores/badge'

  import { closeSideMenu, pageMenus, togglePageMenu } from '$stores/menus'
  import { page } from '$app/stores'
  import { goto } from '$app/navigation'

  onMount(async() => {
    const badgeRequest = await API.getBadge()
    $badge = badgeRequest.data
  })

  $: $badge, updateLinkCounts()

  const updateLinkCounts = () => {
    links = links.map(link => {
      return {
        ...link,
        ...(link.name === 'Jobs') && { count: $badge.TotalJobs },
        ...(link.name === 'All Nodes') && { 
          sublinks: link.sublinks.map(sublink => {
            return {
              ...sublink,
            ...(sublink.name === 'Data holders') && { count: $badge.DataHolders },
            ...(sublink.name === 'Data creators') && { count: $badge.DataCreators }
            }
          })
        }
      }
    })
  }

  const changeUrl = (url: string) => {
    closeSideMenu()
    goto(url)
  }

  let activeMenu = $page.url

  $: if ($page.url) {
    activeMenu = $page.url
  }

  export let withTitle = true
  export let links = [
    {
      name: 'Home',
      url: '/',
      h: '24',
      w: '24',
      svg: [
        'M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6',
      ],
    },
    {
      name: 'Jobs',
      url: '/jobs',
      w: 24,
      h: 24,
      svg: [
        'M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01',
      ],
      count: undefined
    },
    {
      name: 'All Nodes',
      h: 48,
      w: 48,
      svg: [
        'M30.2 42V35.75H22.5V15.25H17.85V21.75H4V6H17.85V12.25H30.2V6H44V21.75H30.2V15.25H25.5V32.75H30.2V26.25H44V42ZM7 9V18.75ZM33.2 29.25V39ZM33.2 9V18.75ZM33.2 18.75H41V9H33.2ZM33.2 39H41V29.25H33.2ZM7 18.75H14.85V9H7Z',
      ],
      sublinks: [
        { name: 'Data holders', count: undefined, url: '/dataHolders' },
        { name: 'Data creators', count: undefined, url: '/dataCreators' },
      ],
    },
    {
      name: 'My Nodes',
      h: 48,
      w: 48,
      svg: [
        'M5.5 40.5Q4.05 40.5 3.025 39.475Q2 38.45 2 37Q2 35.55 3.025 34.525Q4.05 33.5 5.5 33.5Q5.75 33.5 6 33.525Q6.25 33.55 6.65 33.65L16.65 23.65Q16.55 23.25 16.525 23Q16.5 22.75 16.5 22.5Q16.5 21.05 17.525 20.025Q18.55 19 20 19Q21.45 19 22.475 20.025Q23.5 21.05 23.5 22.5Q23.5 22.6 23.35 23.65L28.85 29.15Q29.25 29.05 29.5 29.025Q29.75 29 30 29Q30.25 29 30.5 29.025Q30.75 29.05 31.15 29.15L39.15 21.15Q39.05 20.75 39.025 20.5Q39 20.25 39 20Q39 18.55 40.025 17.525Q41.05 16.5 42.5 16.5Q43.95 16.5 44.975 17.525Q46 18.55 46 20Q46 21.45 44.975 22.475Q43.95 23.5 42.5 23.5Q42.25 23.5 42 23.475Q41.75 23.45 41.35 23.35L33.35 31.35Q33.45 31.75 33.475 32Q33.5 32.25 33.5 32.5Q33.5 33.95 32.475 34.975Q31.45 36 30 36Q28.55 36 27.525 34.975Q26.5 33.95 26.5 32.5Q26.5 32.25 26.525 32Q26.55 31.75 26.65 31.35L21.15 25.85Q20.75 25.95 20.5 25.975Q20.25 26 20 26Q19.9 26 18.85 25.85L8.85 35.85Q8.95 36.25 8.975 36.5Q9 36.75 9 37Q9 38.45 7.975 39.475Q6.95 40.5 5.5 40.5ZM30 15.85 28.45 12.55 25.15 11 28.45 9.45 30 6.15 31.55 9.45 34.85 11 31.55 12.55ZM8 18.4 7 16.2 4.8 15.2 7 14.2 8 12 9 14.2 11.2 15.2 9 16.2Z',
      ],
      sublinks: [
        { name: 'Dashboard', url: '/myNodes' },
        { name: 'Settings', url: '/settings' },
        { name: 'Payouts', url: '/payouts' },
        { name: 'Tax report', url: '/taxReport' },
      ],
    },
    {
      name: 'Reports',
      h: 48,
      w: 48,
      svg: [
        'M9 42Q7.8 42 6.9 41.1Q6 40.2 6 39V9Q6 7.8 6.9 6.9Q7.8 6 9 6H39Q40.2 6 41.1 6.9Q42 7.8 42 9V39Q42 40.2 41.1 41.1Q40.2 42 39 42ZM9 39H39Q39 39 39 39Q39 39 39 39V9Q39 9 39 9Q39 9 39 9H9Q9 9 9 9Q9 9 9 9V39Q9 39 9 39Q9 39 9 39ZM14.2 34.15H17.2V23.9H14.2ZM30.8 34.15H33.8V13.15H30.8ZM22.5 34.15H25.5V28.25H22.5ZM22.5 23.9H25.5V20.9H22.5ZM9 39Q9 39 9 39Q9 39 9 39V9Q9 9 9 9Q9 9 9 9Q9 9 9 9Q9 9 9 9V39Q9 39 9 39Q9 39 9 39Z',
      ],
      sublinks: [
        { name: 'Total Graph Size', url: '/totalGraphSize' },
        { name: 'Job Holding Times', url: '/holdingTimes' },
        { name: 'Job Heatmap', url: '/heatMap' },
        { name: 'Staked tokens', url: '/stakedTokens' },
      ],
    }
  ]
</script>

<div class="py-4 text-gray-500 dark:text-gray-400">
  {#if withTitle}
    <a class="ml-6 text-lg font-bold text-gray-800 dark:text-gray-200" href="/Front/staticstatic"
      >OND</a
    >
  {/if}
  <ul class="mt-6">
    {#each links as link, a}
      <li class="relative px-6 py-3">
        {#if activeMenu == link.url}
          <span
            class="absolute inset-y-0 left-0 w-1 bg-purple-600 rounded-tr-lg rounded-br-lg"
            aria-hidden="true"
          />
        {/if}

        {#if !link.sublinks}
          <a
            class="inline-flex gap-4 items-center w-full text-sm font-semibold transition-colors duration-150 hover:text-gray-800 dark:hover:text-gray-200"
            class:text-gray-800={activeMenu == link.url}
            class:dark:text-gray-100={activeMenu == link.url}
            href={link.url}
            on:click={(e) => {
              e.preventDefault()
              changeUrl(link.url)
            }}
          >
            {#if link.svg}
              <svg
                class="w-5 h-5"
                aria-hidden="true"
                fill="none"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                viewBox="0 0 {link.h} {link.w}"
                stroke="currentColor"
              >
                {#each link.svg as s, b}
                  <path d={s} />
                {/each}
              </svg>
            {/if}
            <div class="flex gap-2">
              <span>{link.name}</span>
              {#if 'count' in link && link.count === undefined} 
                <span class="bg-gray-100 text-gray-800 text-xs mr-2 px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-500 animate-pulse">Loading…</span>
              {:else if link.count}
                <span class="bg-purple-100 text-purple-800 text-xs mr-2 px-2.5 py-0.5 rounded dark:bg-purple-200 dark:text-purple-900">{link.count}</span>
              {/if}
            </div>
            
          </a>
        {:else}
          <button
            on:click={() => togglePageMenu(link.name)}
            class="inline-flex items-center justify-between w-full text-sm font-semibold transition-colors duration-150 hover:text-gray-800 dark:hover:text-gray-200"
            aria-haspopup="true"
          >
            <span class="inline-flex items-center">
              {#if link.svg}
                <svg
                  class="w-5 h-5"
                  aria-hidden="true"
                  fill="none"
                  stroke-width="2"
                  viewBox="0 0 {link.h} {link.w}"
                  stroke="currentColor"
                >
                  {#each link.svg as s, b}
                    <path d={s} />
                  {/each}
                </svg>
              {/if}
              <span class="ml-4">{link.name}</span>
            </span>
            <svg class="w-4 h-4" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20">
              <path
                fill-rule="evenodd"
                d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                clip-rule="evenodd"
              />
            </svg>
          </button>
          {#if $pageMenus[link.name]}
            <ul
              class="p-2 mt-2 space-y-2 overflow-hidden text-sm font-medium text-gray-500 rounded-md shadow-inner bg-gray-50 dark:text-gray-400 dark:bg-gray-900"
              aria-label="submenu"
            >
              {#each link.sublinks as sublink, c}
                <li
                  class="px-2 py-1 transition-colors duration-150 hover:text-gray-800 dark:hover:text-gray-200"
                >
                  
                  <div class="flex gap-2">
                    <a href={sublink.url}>{sublink.name}</a>
                    {#if 'count' in sublink && sublink.count === undefined} 
                      <span class="bg-gray-100 text-gray-800 text-xs px-2.5 py-0.5 rounded dark:bg-gray-700 dark:text-gray-500 animate-pulse">Loading…</span>
                    {:else if sublink.count}
                      <span class="bg-purple-100 text-purple-800 text-xs px-2.5 py-0.5 rounded dark:bg-purple-200 dark:text-purple-900">{sublink.count}</span>
                    {/if}
                  </div>


                </li>
              {/each}
            </ul>
          {/if}
        {/if}
      </li>
    {/each}
  </ul>
</div>
