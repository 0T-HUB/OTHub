import createAuth0Client from '@auth0/auth0-spa-js'
import { user, isAuthenticated, popupOpen } from './stores/auth.js'
import config from './authConfig.js'

const createClient = async () => await createAuth0Client(config)

const loginWithPopup = async (client, options) => {
  popupOpen.set(true)
  try {
    await client.loginWithPopup(options)
    user.set(await client.getUser())
    isAuthenticated.set(true)
  } 
  catch (e) {
    // eslint-disable-next-line
    console.error(e)
  } 
  finally {
    popupOpen.set(false)
  }
}

const logout = (client) => client.logout()

const auth = {
  createClient,
  loginWithPopup,
  logout
}

export default auth