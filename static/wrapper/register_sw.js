const SERVICE_WORKER_PATH = location.origin + '/ochanoco/static/sw/service_worker.js'

const registerServiceWorker = async () => {
  if ("serviceWorker" in navigator) {
    try {
      const registration = await navigator.serviceWorker.register(SERVICE_WORKER_PATH, {
        scope: "/",
      })
      registration.update()

      if (registration.installing) {
        console.log("Service worker installing")
      } else if (registration.waiting) {
        console.log("Service worker installed")
      } else if (registration.active) {
        console.log("Service worker active")
      }
    } catch (error) {
      console.error(`Registration failed with ${error}`)
    }
  }
}

registerServiceWorker()
