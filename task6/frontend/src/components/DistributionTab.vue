<script setup>
import { ref, watch, nextTick } from 'vue'
import Plotly from 'plotly.js-dist-min'
import { LABELS } from '../constants.js'
import { CalculateDistribution } from '../../wailsjs/go/main/App.js'

const props = defineProps({
  calculationResult: {
    type: Object,
    default: null
  }
})

const chartContainer = ref(null)
const distributionData = ref(null)

async function loadDistribution() {
  if (!props.calculationResult) return
  
  try {
    const input = {
      work: props.calculationResult.work,
      time: props.calculationResult.time,
      workChars: {
        planning: props.calculationResult.workChars.planning,
        design: props.calculationResult.workChars.design,
        construction: props.calculationResult.workChars.construction,
        codingTesting: props.calculationResult.workChars.codingTesting,
        integrationTesting: props.calculationResult.workChars.integrationTesting,
      },
      timeChars: {
        planning: props.calculationResult.timeChars.planning,
        design: props.calculationResult.timeChars.design,
        construction: props.calculationResult.timeChars.construction,
        codingTesting: props.calculationResult.timeChars.codingTesting,
        integrationTesting: props.calculationResult.timeChars.integrationTesting,
      },
    }
    distributionData.value = await CalculateDistribution(input)
    await nextTick()
    renderChart()
  } catch (e) {
    console.error('Distribution error:', e)
  }
}

function renderChart() {
  if (!chartContainer.value || !distributionData.value) return
  
  const months = distributionData.value.map((_, i) => i + 1)
  const humans = distributionData.value
  
  const data = [{
    x: months,
    y: humans,
    type: 'bar',
    marker: {
      color: '#3b82f6'
    }
  }]
  
  const layout = {
    title: {
      text: LABELS.tabs.distribution,
      font: { size: 16, color: '#fff' }
    },
    paper_bgcolor: '#161b22',
    plot_bgcolor: '#161b22',
    xaxis: {
      title: { text: 'Месяц', color: '#fff' },
      tickfont: { color: '#fff' },
      gridcolor: '#30363d'
    },
    yaxis: {
      title: { text: 'Количество человек', color: '#fff' },
      tickfont: { color: '#fff' },
      gridcolor: '#30363d',
      rangemode: 'tozero'
    },
    margin: { t: 50, r: 30, b: 50, l: 50 }
  }
  
  Plotly.newPlot(chartContainer.value, data, layout, { displayModeBar: false })
}

watch(() => props.calculationResult, (newVal) => {
  if (newVal) {
    loadDistribution()
  }
}, { immediate: true })
</script>

<template>
  <div class="distribution-tab">
    <div v-if="!calculationResult || !distributionData" class="no-data">
      Выполните расчёт в первой вкладке для отображения распределения
    </div>
    <div v-else class="chart-wrapper">
      <div ref="chartContainer" class="chart-container"></div>
    </div>
  </div>
</template>

<style scoped>
.distribution-tab {
  padding: 20px;
  height: 100%;
  box-sizing: border-box;
}

.no-data {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 300px;
  color: #6b7280;
  font-size: 16px;
}

.chart-wrapper {
  width: 100%;
  height: 100%;
}

.chart-container {
  width: 100%;
  height: 400px;
  min-height: 400px;
}
</style>
