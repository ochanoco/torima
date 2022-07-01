const __originadFetch = window.fetch;
const __originalXhr = window.XMLHttpRequest;

function customFetch(...args) {
  console.log("fetch called");
  return __originadFetch(args);
}

class customXhr extends __originalXhr {
  constructor() {
    console.log("xhr created");
    super();
  }
}

window.fetch = customFetch;
window.XMLHttpRequest = customXhr;
