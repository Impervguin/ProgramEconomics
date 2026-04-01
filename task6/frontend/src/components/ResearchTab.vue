<script setup>
import { ref, nextTick } from 'vue'
import Plotly from 'plotly.js-dist-min'
import { Calculate } from '../../wailsjs/go/main/App.js'
import { LABELS } from '../constants.js'

const kloc = ref(100)
const sced = ref('N')
const loading = ref(false)

const relyPmRef = ref(null)
const relyTmRef = ref(null)
const dataPmRef = ref(null)
const dataTmRef = ref(null)
const cplxPmRef = ref(null)
const cplxTmRef = ref(null)

const levels = ['VL', 'L', 'N', 'H', 'VH']
const modes = ['basic', 'intermediate', 'embedded']

async function calculateResearch() {
  loading.value = true

  try {
    const baseInput = () => ({
      mode: 'basic',
      kloc: kloc.value,
      budgetPerWork: 1,
      computerAttributes: { time: 'N', stor: 'N', virt: 'N', turn: 'N' },
      programAttributes: { rely: 'N', data: 'N', cplx: 'N' },
      staffAttributes: { acap: 'N', aexp: 'N', pcap: 'N', vexp: 'N', lexp: 'N' },
      projectAttributes: { modp: 'N', tool: 'N', sced: sced.value },
    })

    // result[mode][attr] = { pm: [], tm: [] }
    const result = {}

    for (let mode of modes) {
      result[mode] = {}

      for (let attr of ['rely', 'data', 'cplx']) {
        const pm = []
        const tm = []

        for (let lvl of levels) {
          const input = baseInput()
          input.mode = mode
          input.programAttributes[attr] = lvl

          const res = await Calculate(input)

          pm.push(res.work)
          tm.push(res.time)
        }

        result[mode][attr] = { pm, tm }
      }
    }

    await nextTick()

    function plot(refEl, title, yLabel, modeData, type) {
      const traces = [
        {
          x: levels,
          y: modeData.rely[type],
          name: 'RELY',
          type: 'scatter',
        },
        {
          x: levels,
          y: modeData.data[type],
          name: 'DATA',
          type: 'scatter',
        },
        {
          x: levels,
          y: modeData.cplx[type],
          name: 'CPLX',
          type: 'scatter',
        },
      ]

      const layout = {
        title: { text: title, font: { color: '#fff' } },
        paper_bgcolor: '#161b22',
        plot_bgcolor: '#161b22',
        xaxis: { title: 'Уровень', color: '#fff' },
        yaxis: { title: yLabel, color: '#fff' },
        margin: { t: 40, l: 50, r: 20, b: 40 },
      }

      Plotly.newPlot(refEl.value, traces, layout, { displayModeBar: false })
    }

    // PM
    plot(relyPmRef, 'PM (basic)', 'Чел.-мес', result.basic, 'pm')
    plot(dataPmRef, 'PM (intermediate)', 'Чел.-мес', result.intermediate, 'pm')
    plot(cplxPmRef, 'PM (embedded)', 'Чел.-мес', result.embedded, 'pm')

    // TM
    plot(relyTmRef, 'TM (basic)', 'Месяцы', result.basic, 'tm')
    plot(dataTmRef, 'TM (intermediate)', 'Месяцы', result.intermediate, 'tm')
    plot(cplxTmRef, 'TM (embedded)', 'Месяцы', result.embedded, 'tm')

  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="p-5 space-y-4">
    
    <!-- controls -->
    <div class="flex gap-4 items-end">
      <div>
        <label class="block text-sm text-gray-400">KLOC</label>
        <input v-model="kloc" type="number" class="input" />
      </div>

      <div>
        <label class="block text-sm text-gray-400">SCED</label>
        <select v-model="sced" class="input">
          <option v-for="l in levels" :key="l" :value="l">
            {{ LABELS.levels[l] }}
          </option>
        </select>
      </div>

      <button @click="calculateResearch" class="calc-btn">
        {{ loading ? '...' : 'Исследовать' }}
      </button>
    </div>

    <!-- charts -->
    <div class="grid grid-cols-2 gap-4">
      <div ref="relyPmRef" class="chart"></div>
<div ref="relyTmRef" class="chart"></div>

<div ref="dataPmRef" class="chart"></div>
<div ref="dataTmRef" class="chart"></div>

<div ref="cplxPmRef" class="chart"></div>
<div ref="cplxTmRef" class="chart"></div>
    </div>
  </div>
</template>

<style scoped>
.chart {
  height: 300px;
}
</style>