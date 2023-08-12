const PROXY_ORIGIN = "https://127.0.0.1:8080"

const toProxiedUrl = (input, currentOrigin, currentPath, targetOrigin) => {
    console.log({ currentPath, currentOrigin });
    let path = "";

    // input is type of either string or Request
    path = typeof input === "string" ? input : input.url;

    console.log("original url: " + path);

    if (!path.startsWith(currentOrigin)) {
        console.log(
            "proxying a request to an external origin is not supported yet"
        );
        return path;
    }

    if (/^https?:\/\//.exec(path).length > 0) {
        const url = new URL(path);
        path = url.pathname;
    }

    if (!path.startsWith("/")) {
        // if it is relative path, first make it absolute

        // /path/to/hoge/foo.html
        //              ^ Get this indexproxy.example.com
        const slashIndex = currentPath.lastIndexOf("/");
        // fuga/piyo -> /path/to/hoge/fuga/piyo
        path = `${currentPath.slice(0, slashIndex + 1)}${path}`;
    }

    // if it is absolute path in the origin, convert it to URL
    if (path.startsWith("/"))
        path = `${targetOrigin}${path}`;

    console.log(`proxy to ${path}`);

    return path;
}


const customFetch = async (input, init = {}) => {
    console.log("fetch called");
    console.log(input)

    if (input.mode === "navigate") {
        console.log("navigate mode");
        return fetch(input, init);
    }

    init.mode = "cors";

    const url = toProxiedUrl(
        input,
        location.origin,
        location.pathname,
        PROXY_ORIGIN
    );

    if (input instanceof Request) {
        // a Request's properties are read-only, so we need to create new one
        input = new Request(input.url, { ...input });
    } else {
        input = url;
    }

    const resp = fetch(input, init)
    await resp
    console.log({ resp, input, init });

    // const channel = new BroadcastChannel('sw-messages');
    // channel.postMessage({ title: 'Hello from SW' });

    // await sleep(5000)

    return resp
}


self.addEventListener("fetch", async event => {
    event.respondWith(customFetch(event.request));
});
