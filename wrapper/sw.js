const PROXY_ORIGIN = "http://localhost:9000";

function toProxiedUrl(input, currentPath, targetOrigin) {
  let path = "";

  // input is type of either string or Request
  path = typeof input === "string" ? input : input.url;

  console.log("original url: " + path);

  if (/^https?:\/\//.exec(path).length <= 0 && !path.startsWith("/")) {
    // if it is relative path, first make it absolute

    // /path/to/hoge/foo.html
    //              ^ Get this index
    const slashIndex = currentPath.lastIndexOf("/");
    // fuga/piyo -> /path/to/hoge/fuga/piyo
    path = `${currentPath.slice(0, slashIndex + 1)}${path}`;
  }

  // if it is absolute path in the origin, convert it to URL
  if (path.startsWith("/")) {
    path = `${targetOrigin}${path}`;
  }

  console.log(`proxy to ${path}`);

  return path;
}

async function customFetch(input, init = {}) {
  console.log("fetch called");

  init.mode = "cors";

  const url = toProxiedUrl(input, location.pathname, PROXY_ORIGIN);

  if (input instanceof Request) {
    input.url = url;
  } else {
    input = url;
  }

  console.log({ input, init });
  return fetch(input, init);
}

self.addEventListener("fetch", (event) => {
  event.respondWith(customFetch(event.request));
});
