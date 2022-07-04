<script>
  export let data
  let dropdownIsActive = false

  const toggleDropdown = () => dropdownIsActive = !dropdownIsActive

  const exportToJSON = () => {
    const dataString = encodeURIComponent(JSON.stringify(data))
    downloadDataByDataStringAndFileFormat(dataString, 'json')
  }

  const exportToCSV = () => {
    const replacer = (key, value) => value === null ? '' : value
    const header = Object.keys(data[0])
    const dataString = [
      header.join(','),
      ...data.map(row => header.map(fieldName => JSON.stringify(row[fieldName], replacer)).join(','))
    ].join('\r\n')
    downloadDataByDataStringAndFileFormat(dataString, 'csv')
  }

  const downloadDataByDataStringAndFileFormat = (dataString, fileFormat) => {
    const dataTypeByFileFormatMap = new Map([['json', 'application/json'], ['csv', 'text/csv']])
    const dataURI = `data:${dataTypeByFileFormatMap.get(fileFormat)};charset=utf-8,${dataString}`
    const filename = `data.${fileFormat}`
    const linkElement = document.createElement('a')
    linkElement.setAttribute('href', dataURI)
    linkElement.setAttribute('download', filename)
    linkElement.click()
  }

</script>

<div class="relative">
  <button on:click={toggleDropdown} data-dropdown-toggle="dropdown" class="text-white bg-bpurplelue-700 hover:bg-purple-800 focus:ring-4 focus:outline-none focus:ring-purple-300 font-medium rounded-lg text-sm px-4 py-2.5 text-center inline-flex items-center dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-800" type="button">
    <span>Export</span>
    <svg class="w-4 h-4 ml-2" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
    </svg>
  </button>
  <div class="absolute mt-2 invisible z-10 bg-white divide-y divide-gray-100 rounded shadow w-32 dark:bg-gray-700" class:!visible={dropdownIsActive}>
    <ul class="py-1 text-sm text-gray-700 dark:text-gray-200" aria-labelledby="dropdownDefault">
      <li>
        <button on:click={exportToCSV} class="text-left block px-4 py-2 w-full hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">CSV</button>
      </li>
      <li>
        <button on:click={exportToJSON} class="text-left block px-4 py-2 w-full hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">JSON</button>
      </li>
    </ul>
  </div>
</div>