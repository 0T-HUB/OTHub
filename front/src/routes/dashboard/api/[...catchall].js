import cookie from 'cookie'

export const get = async ({ params, url, request }) => {
  try {
    const cookies = cookie.parse(request.headers.get('cookie'))
    const { accessToken } = cookies
    const endpoint = url.href.split('/api/')[1]
    const APIrequest = await fetch(`https://v5api.othub.info/api/${endpoint}`, { headers: { 'Authorization': `Bearer ${accessToken}` }})
    const response = await APIrequest.json() 
    return {
      status: 200,
      body: response
    }
  }
  catch (e) {
    return {
      status: 500,
      body: null
    }
  }
}