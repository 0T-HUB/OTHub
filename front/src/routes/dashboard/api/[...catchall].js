import { config } from 'dotenv'
const  ACCESS_TOKEN  = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Im82eGRreUNvWmVQMjROZWtRd2lpMyJ9.eyJpc3MiOiJodHRwczovL290aHViLmV1LmF1dGgwLmNvbS8iLCJzdWIiOiJvYXV0aDJ8ZGlzY29yZHwyODk4ODQzNzI5OTQwOTcxNTUiLCJhdWQiOlsiaHR0cHM6Ly9vdGh1YmFwaSIsImh0dHBzOi8vb3RodWIuZXUuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTY1NzEyMDU3OSwiZXhwIjoxNjU3MjA2OTc5LCJhenAiOiJZeDM4NFdleERRajl4ejhEQks2Mm1kVXc3NEc1NGYyQiIsInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwgb2ZmbGluZV9hY2Nlc3MifQ.ZZG8shDamw7GBMQUyMo4jjYBifexotTp3gFSAmPDrecqJAYWZPfwjJU2Fpa2CGP_07Ex-XoCw9SlJezOiZAP5vXM-RFQyw0CDIB-wv53JFSAiCTNVT1DA8llTgxGPeacmoz3I5u7EDHyWWb68JRnJeEA9zKv64NRaLXnTVSsO2ozDHVIDjoU87g9u_WaxuLSP1MV8sy6DSNTN4_Zg8VYOoUOb-sfuwlnWdKT-Hnq_jr9l28DC59dQSn2oX3LSO9QjF0GF8BRsezLwYRhR9vFbt8_S_n6gt0_8s658-htCNBnfajuYX7tChYoKQsJmx0F9Rd1jEP8TG-ocBgqUHL2yA"
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
