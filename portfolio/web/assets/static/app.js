const form = document.querySelector("#simulation-form");
const pathsInput = document.querySelector("#paths");
const pathsOutput = document.querySelector("#paths-output");
const statusEl = document.querySelector("#status");
const chart = document.querySelector("#paths-chart");
const ctx = chart.getContext("2d");

const fields = {
  title: document.querySelector("#result-title"),
  expectedEnd: document.querySelector("#expectedEnd"),
  p05: document.querySelector("#p05"),
  p50: document.querySelector("#p50"),
  p95: document.querySelector("#p95"),
  count: document.querySelector("#count"),
  start: document.querySelector("#start"),
  generated: document.querySelector("#generated"),
  chartHeading: document.querySelector("#chart-heading"),
};

pathsInput.addEventListener("input", () => {
  pathsOutput.value = pathsInput.value;
});

form.addEventListener("submit", async (event) => {
  event.preventDefault();
  await runSimulation();
});

window.addEventListener("resize", () => {
  if (window.latestSimulation) {
    drawChart(window.latestSimulation);
  }
});

async function runSimulation() {
  const payload = Object.fromEntries(new FormData(form).entries());
  const request = {
    ticker: String(payload.ticker || "AAPL").trim().toUpperCase(),
    startPrice: Number(payload.startPrice),
    drift: Number(payload.drift),
    volatility: Number(payload.volatility),
    days: Number(payload.days),
    paths: Number(payload.paths),
    seed: Number(payload.seed),
  };

  setStatus("Simulando", false);
  form.querySelector("button").disabled = true;

  try {
    const response = await fetch("/api/simulate", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(request),
    });
    const data = await response.json();
    if (!response.ok) {
      throw new Error(data.error || "Não foi possível simular");
    }
    window.latestSimulation = data;
    renderSummary(data);
    drawChart(data);
    setStatus("Atualizado", false);
  } catch (error) {
    setStatus(error.message, true);
  } finally {
    form.querySelector("button").disabled = false;
  }
}

function renderSummary(data) {
  fields.title.textContent = `${data.ticker || "Ativo"} · ${data.days} dias`;
  fields.expectedEnd.textContent = money(data.expectedEnd);
  fields.p05.textContent = money(data.percentile05);
  fields.p50.textContent = money(data.percentile50);
  fields.p95.textContent = money(data.percentile95);
  fields.count.textContent = number(data.count);
  fields.start.textContent = money(data.startPrice);
  fields.generated.textContent = new Date(data.generatedAt).toLocaleString("pt-BR");
  fields.chartHeading.textContent = `${data.paths.length} caminhos exibidos`;
}

function drawChart(data) {
  const rect = chart.getBoundingClientRect();
  const ratio = window.devicePixelRatio || 1;
  chart.width = Math.max(640, Math.floor(rect.width * ratio));
  chart.height = Math.max(320, Math.floor(rect.height * ratio));
  ctx.setTransform(ratio, 0, 0, ratio, 0, 0);

  const width = chart.width / ratio;
  const height = chart.height / ratio;
  const padding = { top: 18, right: 18, bottom: 34, left: 58 };
  const plotWidth = width - padding.left - padding.right;
  const plotHeight = height - padding.top - padding.bottom;
  const allValues = data.paths.flat();
  const min = Math.min(...allValues, data.percentile05) * 0.98;
  const max = Math.max(...allValues, data.percentile95) * 1.02;

  ctx.clearRect(0, 0, width, height);
  ctx.fillStyle = "#ffffff";
  ctx.fillRect(0, 0, width, height);

  drawGrid(width, height, padding, plotWidth, plotHeight, min, max);

  data.paths.forEach((path) => {
    drawLine(path, padding, plotWidth, plotHeight, min, max, {
      stroke: "rgba(44, 111, 183, 0.18)",
      width: 1.2,
    });
  });

  drawLine(medianSeries(data.paths), padding, plotWidth, plotHeight, min, max, {
    stroke: "rgba(196, 122, 18, 0.92)",
    width: 2.6,
  });
}

function drawGrid(width, height, padding, plotWidth, plotHeight, min, max) {
  ctx.strokeStyle = "#dfe6ea";
  ctx.lineWidth = 1;
  ctx.fillStyle = "#65717d";
  ctx.font = "12px Inter, system-ui, sans-serif";

  for (let i = 0; i <= 4; i++) {
    const y = padding.top + (plotHeight / 4) * i;
    ctx.beginPath();
    ctx.moveTo(padding.left, y);
    ctx.lineTo(width - padding.right, y);
    ctx.stroke();

    const value = max - ((max - min) / 4) * i;
    ctx.fillText(money(value), 8, y + 4);
  }

  ctx.fillText("Hoje", padding.left, height - 10);
  ctx.fillText("Fim", width - padding.right - 22, height - 10);
}

function drawLine(values, padding, plotWidth, plotHeight, min, max, style) {
  ctx.beginPath();
  values.forEach((value, index) => {
    const x = padding.left + (plotWidth * index) / Math.max(1, values.length - 1);
    const y = padding.top + plotHeight - ((value - min) / Math.max(1, max - min)) * plotHeight;
    if (index === 0) {
      ctx.moveTo(x, y);
    } else {
      ctx.lineTo(x, y);
    }
  });
  ctx.strokeStyle = style.stroke;
  ctx.lineWidth = style.width;
  ctx.stroke();
}

function medianSeries(paths) {
  if (!paths.length) {
    return [];
  }
  const length = paths[0].length;
  const series = [];
  for (let day = 0; day < length; day++) {
    const values = paths.map((path) => path[day]).sort((a, b) => a - b);
    const middle = Math.floor(values.length / 2);
    series.push(values.length % 2 ? values[middle] : (values[middle - 1] + values[middle]) / 2);
  }
  return series;
}

function setStatus(message, isError) {
  statusEl.textContent = message;
  statusEl.classList.toggle("error", isError);
}

function money(value) {
  return Number(value).toLocaleString("pt-BR", {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2,
  });
}

function number(value) {
  return Number(value).toLocaleString("pt-BR");
}

runSimulation();
