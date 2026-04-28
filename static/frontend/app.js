const API_BASE = "/re-partners/shipping-service-api/v1";

let localSizes = [];

function showMessage(elementId, text, isError) {
  const el = document.getElementById(elementId);
  el.className = isError ? "msg err" : "msg ok";
  el.textContent = text;
}

function clearMessage(elementId) {
  const el = document.getElementById(elementId);
  el.className = "msg";
  el.textContent = "";
}

function renderTags() {
  const container = document.getElementById("tags");
  container.innerHTML = "";

  localSizes
    .slice()
    .sort((a, b) => a - b)
    .forEach((size) => {
      const tag = document.createElement("span");
      tag.className = "tag";
      tag.innerHTML = `${size} <button onclick="removeSize(${size})" title="Remove">x</button>`;
      container.appendChild(tag);
    });
}

async function boot() {
  try {
    const response = await fetch(`${API_BASE}/packs`);
    if (!response.ok) {
      throw new Error("failed to load pack sizes");
    }

    const data = await response.json();
    localSizes = Array.isArray(data.pack_sizes) ? data.pack_sizes : [];
  } catch (_error) {
    localSizes = [250, 500, 1000, 2000, 5000];
    showMessage("sizeMsg", "Could not load pack sizes from API. Using defaults.", true);
  }

  renderTags();
}

function addSize() {
  const input = document.getElementById("newSize");
  const value = parseInt(input.value, 10);

  if (!Number.isInteger(value) || value <= 0) {
    return;
  }

  if (!localSizes.includes(value)) {
    localSizes.push(value);
    renderTags();
  }

  input.value = "";
  input.focus();
}

function removeSize(size) {
  localSizes = localSizes.filter((current) => current !== size);
  renderTags();
}

async function saveSizes() {
  clearMessage("sizeMsg");

  if (localSizes.length === 0) {
    try {
      const getResponse = await fetch(`${API_BASE}/packs`);
      if (!getResponse.ok) {
        const errorText = await getResponse.text();
        throw new Error(errorText || "failed to reload pack sizes");
      }

      const updatedData = await getResponse.json();
      localSizes = Array.isArray(updatedData.pack_sizes) ? updatedData.pack_sizes : [];
      renderTags();
      showMessage("sizeMsg", "No sizes to save. Reloaded current values.", true);
    } catch (error) {
      showMessage("sizeMsg", error.message || "failed to reload pack sizes", true);
    }
    return;
  }

  try {
    const putResponse = await fetch(`${API_BASE}/packs`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ pack_sizes: localSizes }),
    });

    if (!putResponse.ok) {
      const errorText = await putResponse.text();
      throw new Error(errorText || "failed to save pack sizes");
    }

    const getResponse = await fetch(`${API_BASE}/packs`);
    if (!getResponse.ok) {
      const errorText = await getResponse.text();
      throw new Error(errorText || "failed to reload pack sizes");
    }

    const updatedData = await getResponse.json();
    localSizes = Array.isArray(updatedData.pack_sizes) ? updatedData.pack_sizes : [];
    renderTags();

    showMessage("sizeMsg", "Sizes saved.", false);
    setTimeout(() => clearMessage("sizeMsg"), 2500);
  } catch (error) {
    showMessage("sizeMsg", error.message || "failed to save pack sizes", true);
  }
}

function renderResult(order, packs) {
  const resultArea = document.getElementById("resultArea");

  if (!Array.isArray(packs) || packs.length === 0) {
    resultArea.innerHTML = '<p class="result-empty">No packs returned by API.</p>';
    return;
  }

  const orderedPacks = packs
    .slice()
    .sort((a, b) => Number(b.pack_size) - Number(a.pack_size));

  const totalItems = orderedPacks.reduce(
    (sum, item) => sum + Number(item.pack_size) * Number(item.quantity),
    0,
  );
  const totalPacks = orderedPacks.reduce((sum, item) => sum + Number(item.quantity), 0);
  const overage = totalItems - order;

  const rows = orderedPacks
    .map(
      (item) => `
        <tr>
          <td><span class="pack-badge">${item.pack_size}</span></td>
          <td>x ${item.quantity}</td>
          <td>${(Number(item.pack_size) * Number(item.quantity)).toLocaleString()} items</td>
        </tr>`,
    )
    .join("");

  resultArea.innerHTML = `
    <div class="result-summary">
      <div class="stat"><div class="val">${order.toLocaleString()}</div><div class="lbl">Ordered</div></div>
      <div class="stat"><div class="val">${totalItems.toLocaleString()}</div><div class="lbl">Shipped</div></div>
      <div class="stat"><div class="val">${totalPacks.toLocaleString()}</div><div class="lbl">Packs</div></div>
      <div class="stat"><div class="val">+${overage.toLocaleString()}</div><div class="lbl">Overage</div></div>
    </div>
    <table>
      <thead><tr><th>Pack size</th><th>Qty</th><th>Subtotal</th></tr></thead>
      <tbody>${rows}</tbody>
    </table>`;
}

async function calculate() {
  const calcErr = document.getElementById("calcErr");
  calcErr.textContent = "";

  const order = parseInt(document.getElementById("order").value, 10);
  if (!Number.isInteger(order) || order <= 0) {
    calcErr.textContent = "Please enter a positive integer.";
    return;
  }

  const button = document.getElementById("calcBtn");
  button.disabled = true;
  button.textContent = "Calculating...";

  try {
    const response = await fetch(`${API_BASE}/calculator`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ amount: order }),
    });

    if (!response.ok) {
      const errorText = await response.text();
      calcErr.textContent = errorText || "failed to calculate packs";
      return;
    }

    const data = await response.json();
    renderResult(order, data);
  } catch (error) {
    calcErr.textContent = `Network error: ${error.message}`;
  } finally {
    button.disabled = false;
    button.textContent = "Calculate Packs";
  }
}

window.addSize = addSize;
window.removeSize = removeSize;
window.saveSizes = saveSizes;
window.calculate = calculate;

boot();
