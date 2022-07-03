(function () {
  const PROXY_ORIGIN = "http://localhost:9000";

  const __originadFetch = window.fetch;
  const __originalXhr = window.XMLHttpRequest;

  function toProxiedUrl(input, currentPath, targetOrigin) {
    let path = "";

    // input is type of either string or Request
    path = typeof input === "string" ? input : input.url;

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

    return path;
  }

  function customFetch(input, init = {}) {
    console.log("fetch called");

    init.mode = "cors";

    console.log("original url: " + input);
    const url = toProxiedUrl(input, location.pathname, PROXY_ORIGIN);

    console.log(`proxy to ${url}`);

    return __originadFetch(url, init);
  }

  class customXhr extends __originalXhr {
    constructor() {
      console.log("xhr created");
      super();
    }

    open(method, url, ...args) {
      const proxiedUrl = toProxiedUrl(url, location.pathname, PROXY_ORIGIN);
      return super.open(method, proxiedUrl, ...args);
    }
  }
  window.fetch = customFetch;
  window.XMLHttpRequest = customXhr;
})();
