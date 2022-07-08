import { config } from 'dotenv'
const  ACCESS_TOKEN  = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Im82eGRreUNvWmVQMjROZWtRd2lpMyJ9.eyJpc3MiOiJodHRwczovL290aHViLmV1LmF1dGgwLmNvbS8iLCJzdWIiOiJvYXV0aDJ8ZGlzY29yZHwyODk4ODQzNzI5OTQwOTcxNTUiLCJhdWQiOlsiaHR0cHM6Ly9vdGh1YmFwaSIsImh0dHBzOi8vb3RodWIuZXUuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTY1NzI3MjM2NCwiZXhwIjoxNjU3MzU4NzY0LCJhenAiOiJZeDM4NFdleERRajl4ejhEQks2Mm1kVXc3NEc1NGYyQiIsInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwgb2ZmbGluZV9hY2Nlc3MifQ.liP8pkk7dCdX19ECbuEAt8Du3y2CaInCSZEs5Tw0NhQbrrefBmJdCS9ub8YhhqNZyK6g4A88oGLMObgssLe7QEWN9pJmR_MjGgr0sTvUGimCBA31JL9xLWTtmjnLX2alObtwFLQZ9yVGL-n1dIzfSZAH7z254EiAmwUy9SQyXdldxbkxCCp_Tq_kkaOkGda7QodCfVhHNbkcKcwgMcTu_zSM5NyQELTzPigY7FJ_sjmlyAoLgkZPTlMqPhvX4e16F3am8hCLmAKeleQonCDMFWy6kJxVLuerewzDLMmMBXWN6_tbItHqK_Qp7k1PZMU07csbkh7PIkYUs-_mWePX1w"
console.log(ACCESS_TOKEN)

export const get = async (args) => await proxyRequest(args)
export const post = async (args) => await proxyRequest(args)
export const del = async (args) => await proxyRequest(args)


const proxyRequest = async ({ url, request }) => {
  try {
    console.log("url = ",url.pathname)
    const { method } = request
    const endpoint = url.href.split('/api/')[1]
    const APIrequest = await fetch(`https://v5api.othub.info/api/${endpoint}`, {
      method,
      headers: {
        'Authorization': `Bearer ${ACCESS_TOKEN}`,
        ...(method === 'POST') && { 'Content-Type': 'application/json' },
      },
    })
    const response = await APIrequest.json()
    const { headers } = APIrequest
    const totalCount = headers.get('x-total-count')
    return {
      status: 200,
      body: {
        ...(response) && { data: response },
        ...(totalCount) && { totalCount }
      }
    }
  }
  catch (e) {
    console.log(e)
  }
}
