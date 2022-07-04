import { config } from 'dotenv'
const  ACCESS_TOKEN  = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Im82eGRreUNvWmVQMjROZWtRd2lpMyJ9.eyJpc3MiOiJodHRwczovL290aHViLmV1LmF1dGgwLmNvbS8iLCJzdWIiOiJvYXV0aDJ8ZGlzY29yZHwyODk4ODQzNzI5OTQwOTcxNTUiLCJhdWQiOlsiaHR0cHM6Ly9vdGh1YmFwaSIsImh0dHBzOi8vb3RodWIuZXUuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTY1NjkzNjMwOSwiZXhwIjoxNjU3MDIyNzA5LCJhenAiOiJZeDM4NFdleERRajl4ejhEQks2Mm1kVXc3NEc1NGYyQiIsInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwgb2ZmbGluZV9hY2Nlc3MifQ.uUX1ZjhKYMW9UHnpk7iVsTMsVZzatIqz2w0AcEB7_s6sTKQNNfCNk4FRdpGbpx_tRdIc5wzuf5c2orPDMKfczx2A-hilP_1iRpuyxOPYQ6-DLDE1DuweXdBKkWudroqpHCqW_ay2h24KPEtnUovOccHYhz8DF8xP35e9imOoqQAGVybfmzR5ooElg2KAYOM5PFQinKQFQHxFy7xyumG7amPokGqwsEZteDpFNEAW4hGB-g4yvYCo5pG48aKKFx_dgJKYr7Qh2yR0hvV0u1VFIhz52vMbiE8NvpKJzcHE4JSCXqHE68pBOep1OaUYKobwDZKOmOXIxj3tRY9qguRLPw"
console.log(ACCESS_TOKEN)

export const get = async (args) => await proxyRequest(args)
export const post = async (args) => await proxyRequest(args)
export const del = async (args) => await proxyRequest(args)

const proxyRequest = async ({ url, request }) => {
  try {
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
