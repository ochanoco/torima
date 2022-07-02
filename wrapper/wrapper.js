const PROXY_ADDR = "http://localhost:9000";

const __originadFetch = window.fetch;
const __originalXhr = window.XMLHttpRequest;

function customFetch(input, init = {}) {
  console.log("fetch called");

  let newAddr = "";

  // input is type of either string or Request
  newAddr = typeof input === "string" ? input : input.url;

  if (!newAddr.startsWith("http://") && !newAddr.startsWith("/")) {
    // if it is relative path, first make it absolute

    // hoge/fuga/piyo
    //          ^ Get this index
    const slashIndex = newAddr.lastIndexOf("/");
    // hoge/fuga/piyo -> hoge/fuga/
    location.pathname.slice(0, slashIndex + 1);
  }

  // if it is absolute path in the origin, convert it to URL
  if (newAddr.startsWith("/")) {
    newAddr = `${location.origin}${newAddr}`;
  }

  init.mode = "cors";

  const url = PROXY_ADDR;

  // TODO: maybe i need to add CORS-safelisted headers (https://fetch.spec.whatwg.org/#cors-safelisted-request-header)
  return __originadFetch(url, init);
}

class customXhr extends __originalXhr {
  constructor() {
    console.log("xhr created");
    super();
  }
}

window.fetch = customFetch;
window.XMLHttpRequest = customXhr;
