// Load Go WASM runtime and provide a function to run Go code in-browser
// This script assumes wasm_exec.js is available in the same directory or via CDN

// Dynamically load wasm_exec.js if not already loaded
(function loadWasmExec() {
  if (!window.Go) {
    var script = document.createElement('script');
    script.src = 'https://golang.org/dl/go1.21.0.beta1.wasm_exec.js'; // Use a CDN or local copy if needed
    script.onload = function() {
      window.goWasmReady = true;
    };
    document.head.appendChild(script);
  } else {
    window.goWasmReady = true;
  }
})();

// Utility to run Go code in-browser using WASM
function setRunButtonsEnabled(enabled) {
  document.querySelectorAll('.run-btn').forEach(btn => {
    btn.disabled = !enabled;
    btn.textContent = '▶ Run Code';
  });
}
setRunButtonsEnabled(false);

window.addEventListener('DOMContentLoaded', function() {
  if (window.goWasmReady) setRunButtonsEnabled(true);
  else {
    let checkReady = setInterval(() => {
      if (window.goWasmReady) {
        setRunButtonsEnabled(true);
        clearInterval(checkReady);
      }
    }, 100);
  }
});

document.runGoCode = async function(code, outputElemId) {
  const btn = document.activeElement;
  if (btn && btn.classList.contains('run-btn')) {
    btn.textContent = '⏳ Running...';
    btn.disabled = true;
  }
  if (!window.goWasmReady) {
    document.getElementById(outputElemId).textContent = 'Go WASM runtime is still loading. Please wait...';
    if (btn && btn.classList.contains('run-btn')) {
      btn.textContent = '▶ Run Code';
      btn.disabled = false;
    }
    return;
  }
  if (!window.Go) {
    document.getElementById(outputElemId).textContent = 'Go WASM runtime not found.';
    if (btn && btn.classList.contains('run-btn')) {
      btn.textContent = '▶ Run Code';
      btn.disabled = false;
    }
    return;
  }

  // Prepare the Go WASM runner
  const go = new window.Go();
  // Compile Go code to WASM using an online API (e.g., play.golang.org/compile)
  // For privacy and offline, you could use a local WASM compiler, but here we use the public API
  const resp = await fetch('https://play.golang.org/compile', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ version: 2, body: code, withVet: false })
  });
  const data = await resp.json();
  if (!data || !data.Body) {
    document.getElementById(outputElemId).textContent = 'Compilation failed: ' + (data.Errors || 'Unknown error');
    if (btn && btn.classList.contains('run-btn')) {
      btn.textContent = '▶ Run Code';
      btn.disabled = false;
    }
    return;
  }
  // Decode base64 WASM
  const wasmBinary = Uint8Array.from(atob(data.Body), c => c.charCodeAt(0));
  // Run the WASM
  const resultElem = document.getElementById(outputElemId);
  resultElem.textContent = '';
  const importObject = go.importObject;
  try {
    const wasmModule = await WebAssembly.instantiate(wasmBinary, importObject);
    go.argv = [];
    go.env = {};
    go.exit = (code) => {};
    await go.run(wasmModule.instance);
  } catch (err) {
    resultElem.textContent = 'Runtime error: ' + err;
  }
  if (btn && btn.classList.contains('run-btn')) {
    btn.textContent = '▶ Run Code';
    btn.disabled = false;
  }
}; 