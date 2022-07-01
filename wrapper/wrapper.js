const __originadFetch = window.fetch;

function customFetch(...args) {
  console.log("fetch called");
  return __originadFetch(args);
}

window.fetch = customFetch;
