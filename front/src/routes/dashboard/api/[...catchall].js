import cookie from 'cookie'

export const get = async (args) => await proxyRequest(args)
export const post = async (args) => await proxyRequest(args)
export const del = async (args) => await proxyRequest(args)

const proxyRequest = async ({ params, url, request }) => {
  const { method } = request
  const cookies = cookie.parse(request.headers.get('cookie'))
  const { accessToken } = cookies
  const endpoint = url.href.split('/api/')[1]
  const APIrequest = await fetch(`https://v5api.othub.info/api/${endpoint}`, { 
    method,
    headers: { 
      'Authorization': `Bearer ${accessToken}`,
      ...(method === 'POST') && { 'Content-Type': 'application/json' },
    },
  })
  const text = await APIrequest.text()
  return {
    status: 200,
    ...(text) && { body: JSON.parse(text) }
  }
}