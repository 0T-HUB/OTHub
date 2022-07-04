import { get } from 'svelte/store'
import { notifications } from '$stores/notifications'

const getJSONFromLocalStorageByKey = (key) => JSON.parse(localStorage.getItem(key))
const putJSONToLocalStorageByKey = (item, key) => localStorage.setItem(key, JSON.stringify(item))

export const updateNotificationsInLocalStorageFromStore = () => putJSONToLocalStorageByKey(get(notifications), 'notifications')
export const updateNotificationsInStoreFromLocalStorage = () => notifications.set(getJSONFromLocalStorageByKey('notifications'))

export const generatePendingNotificationIfItDoesNotExist = () => {
  const pendingNotificationDoesNotExist = !checkIfPendingNotificationExists()
  if (pendingNotificationDoesNotExist) createPendingNotification()
  launchPendingNotificationTimer()
}

const checkIfPendingNotificationExists = () => Boolean(getJSONFromLocalStorageByKey('pendingNotification'))
const createPendingNotification = () => putJSONToLocalStorageByKey({ timestamp: createNotificationTimestamp(), text: generateNotificationText(), unseen: true }, 'pendingNotification')

const launchPendingNotificationTimer = () => {
  const pendingNotification = getJSONFromLocalStorageByKey('pendingNotification')
  const currentTimestamp = new Date().getTime()
  const notificationIntervalInMiliseconds = pendingNotification.timestamp - currentTimestamp
  setTimeout(emitPendingNotification, notificationIntervalInMiliseconds)
}

const emitPendingNotification = () => {
  const pendingNotification = getJSONFromLocalStorageByKey('pendingNotification')
  const currentNotifications = get(notifications) || []
  const updatedNotifications = [...currentNotifications, pendingNotification]
  localStorage.removeItem('pendingNotification')
  notifications.set(updatedNotifications)
  setTimeout(() => {
    updateNotificationsInLocalStorageFromStore()
    generatePendingNotificationIfItDoesNotExist()
  }, 100)
  
}

const generateNotificationText = () => `New job found for nodes #${generateRandomIntegerBetweenOneAndSix()} for ${generateRandomDecimalBetweenOneTenthAndThree()} trac (${generateRandomHoldingTimeBetweenSixMonthsAndThreeYears()} holding time)`
const generateRandomNotificationIntervalInMilisecondsFromMinutes = (notificationIntervalInMinutes) => generateRandomNumberOfSecondsFromMinuteInterval(notificationIntervalInMinutes) * 1000
const createNotificationTimestamp = () => new Date().getTime() + generateRandomNotificationIntervalInMilisecondsFromMinutes(1)
const generateRandomNumberOfSecondsFromMinuteInterval = (minuteInterval) => Math.ceil(Math.random() * minuteInterval * 60)
const generateRandomIntegerBetweenOneAndSix = () => Math.ceil(Math.random() * 6)
const generateRandomDecimalBetweenOneTenthAndThree = () => Number((Math.random() * 3).toFixed(1))
const generateRandomNumberBetweenSixAndThirtySix = () => Math.ceil(Math.random() * 36)

const generateRandomHoldingTimeBetweenSixMonthsAndThreeYears = () => {
  const numberOfMonths = generateRandomNumberBetweenSixAndThirtySix()
  const numberOfWholeYears = Math.floor(numberOfMonths / 12)
  const numberOfRemainingMonths = numberOfMonths % 12
  const yearsString = numberOfWholeYears ? `${numberOfWholeYears} year${numberOfWholeYears > 1 ? 's' : ''}` : ''
  const monthsString = numberOfRemainingMonths ? `${numberOfRemainingMonths} month${numberOfRemainingMonths > 1 ? 's' : ''}` : ''
  const stringsArray = [yearsString, monthsString].filter(Boolean) 
  const randomHoldingTimeBetweenSixMonthsAndThreeYears = stringsArray.join(' ')
  return randomHoldingTimeBetweenSixMonthsAndThreeYears
}