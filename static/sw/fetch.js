// https://127.0.0.1:8080/
let PROTECTION_SCOPE = []
let LOCAL_WHITELIST = [
    '/ochanoco/login',
    '/ochanoco/auth/callback'
]

const SLEEP_TIME = 3000

const init = async () => {
    const resp = await fetch(location.origin + `/ochanoco/status`)
    const data = await resp.json()
    PROTECTION_SCOPE = data.protection_scope
}

const channel = new BroadcastChannel('sw-messages')

const normalizeURL = (input, init) => {
    let url;

    if (typeof input === "string") {
        if (input[0] === "/")
            input = `${input.slice(1, input.length)}`;

        if (!input.startsWith("https://")) {
            input = location.host + "/" + input
        }

        url = new URL(input)
        input = new Request(url, init)
    }
    else {
        url = new URL(input.url)
    }

    return { url, input }
}

const customFetch = async (input, init = {}) => {
    console.log("input: ", input)
    var { url, input } = normalizeURL(input, init)


    const isProxyOrigin = url.host === location.host
    const isInProtectionScope = PROTECTION_SCOPE.includes(url.host)
    const isInLocalWhitelist = LOCAL_WHITELIST.includes(url.pathname)
    const isNavigate = input.mode === "navigate"

    console.log("target: ", url)
    console.log("protection_scope: ", PROTECTION_SCOPE, isProxyOrigin)
    console.log("host: ", url.host, isInProtectionScope)
    console.log("mode:", input.mode, input.mode == "navigate")
    console.log(!isProxyOrigin && !isInProtectionScope || isNavigate)

    if (!isProxyOrigin && !isInProtectionScope) {
        console.log("unmodify to proxy because not in protection scope")
        return fetch(input, init)
    }

    if (!isProxyOrigin && isInProtectionScope) {
        console.log("modify to proxy")

        url.pathname = `/ochanoco/redirect/${url.host}/${url.pathname}`
        url.host = location.host
        input.url = url
    }

    if (isInLocalWhitelist) {
        console.log("unmodify to check authorized")
        return fetch(input, init)
    }


    console.log("modify to check authorized")


    // input = new Request(input.url, { ...input })

    let fetching = fetch(input, init)
    const resp = await fetching

    if (resp.status === 401) {
        setTimeout(() => {
            console.log("Authentication needed")
            channel.postMessage({ title: 'Authentication needed' })

        }, SLEEP_TIME)
    }

    return fetching
}


self.addEventListener("fetch", async event => {
    event.respondWith(
        customFetch(event.request)
    );
});

init()