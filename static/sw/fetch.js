// https://127.0.0.1:8080/
let PROTECTION_SCOPE = []

const init = async () => {
    const resp = await fetch("/ochanoco/config.json")
    const scope = await resp.json()
    PROTECTION_SCOPE = scope
}

const channel = new BroadcastChannel('sw-messages')


const customFetch = async (input, init = {}) => {
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


    const isProxyOrigin = url.host === location.host
    const isInProtectionScope = PROTECTION_SCOPE.includes(url.host)

    console.log("protection_scope: ", PROTECTION_SCOPE)
    console.log("host: ", url.host)

    if (!isProxyOrigin && !isInProtectionScope || input.mode === "navigate") {
        console.log("unmodify to proxy")
        return fetch(input, init)
    }

    if (!isProxyOrigin) {
        console.log("modify to proxy")

        url.pathname = `/ochanoco/redirect/${url.host}/${url.pathname}`
        url.host = location.host
        input.url = url

    }

    input = new Request(input.url, { ...input })


    let fetching = fetch(input, init)
    const resp = await fetching

    if (resp.status === 500) {
        channel.postMessage({ title: 'Hello from SW' })
    }

    return fetching
}


self.addEventListener("fetch", async event => {
    event.respondWith(
        customFetch(event.request)
    );
});

init()