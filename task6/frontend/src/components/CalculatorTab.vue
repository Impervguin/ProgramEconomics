<script setup>
import { ref, computed } from 'vue'
import { LABELS } from '../constants.js'
import { Calculate } from '../../wailsjs/go/main/App.js'

const emit = defineEmits(['calculate'])

const kloc = ref(100)
const budgetPerWork = ref(1)
const mode = ref('basic')
const result = ref(null)
const loading = ref(false)

const attributes = ref({
  program: { rely: 'N', data: 'N', cplx: 'N' },
  computer: { time: 'N', stor: 'N', virt: 'N', turn: 'N' },
  staff: { acap: 'N', aexp: 'N', pcap: 'N', vexp: 'N', lexp: 'N' },
  project: { modp: 'N', tool: 'N', sced: 'N' },
})

const levels = ['VL', 'L', 'N', 'H', 'VH']

const attributeGroups = computed(() => [
  { key: 'program', label: LABELS.attributes.program, attrs: ['rely', 'data', 'cplx'] },
  { key: 'computer', label: LABELS.attributes.computer, attrs: ['time', 'stor', 'virt', 'turn'] },
  { key: 'staff', label: LABELS.attributes.staff, attrs: ['acap', 'aexp', 'pcap', 'vexp', 'lexp'] },
  { key: 'project', label: LABELS.attributes.project, attrs: ['modp', 'tool', 'sced'] },
])

const modes = [
  { value: 'basic', label: LABELS.mode.basic },
  { value: 'intermediate', label: LABELS.mode.intermediate },
  { value: 'embedded', label: LABELS.mode.embedded },
]

async function calculate() {
  loading.value = true
  try {
    const input = {
      mode: mode.value,
      kloc: kloc.value,
      budgetPerWork: budgetPerWork.value,
      computerAttributes: {
        time: attributes.value.computer.time,
        stor: attributes.value.computer.stor,
        virt: attributes.value.computer.virt,
        turn: attributes.value.computer.turn,
      },
      programAttributes: {
        rely: attributes.value.program.rely,
        data: attributes.value.program.data,
        cplx: attributes.value.program.cplx,
      },
      staffAttributes: {
        acap: attributes.value.staff.acap,
        aexp: attributes.value.staff.aexp,
        pcap: attributes.value.staff.pcap,
        vexp: attributes.value.staff.vexp,
        lexp: attributes.value.staff.lexp,
      },
      projectAttributes: {
        modp: attributes.value.project.modp,
        tool: attributes.value.project.tool,
        sced: attributes.value.project.sced,
      },
    }
    result.value = await Calculate(input)
    emit('calculate', result.value)
  } catch (e) {
    console.error(e)
  } finally {
    loading.value = false
  }
}

const workActivities = ['planning', 'design', 'construction', 'codingTesting', 'integrationTesting']
const budgetActivities = ['analysis', 'design', 'programming', 'testPlanning', 'verification', 'office', 'configManagement', 'documentation']

function getWorkRow(activity) {
  const work = result.value?.workChars?.[activity] || 0
  const coeff = result.value?.workCoeffs?.[activity] || 0
  const percent = (coeff * 100).toFixed(1)
  return { percent, work: work.toFixed(2) }
}

function getTimeRow(activity) {
  const time = result.value?.timeChars?.[activity] || 0
  const coeff = result.value?.timeCoeffs?.[activity] || 0
  const percent = (coeff * 100).toFixed(1)
  return { percent, time: time.toFixed(2) }
}

function getBudgetRow(activity) {
  const budget = result.value?.budgetChars?.[activity] || 0
  const coeff = result.value?.budgetCoeffs?.[activity] || 0
  const work = result.value?.budgetWork?.[activity] || 0
  const percent = (coeff * 100).toFixed(1)
  return { percent, work: work.toFixed(2), budget: budget.toFixed(2) }
}

const workTotal = computed(() => {
  if (!result.value) return { percent: '0', work: '0' }
  const sumPercent = workActivities.reduce((acc, a) => acc + (result.value.workCoeffs?.[a] || 0), 0)
  const sumWork = workActivities.reduce((acc, a) => acc + (result.value.workChars?.[a] || 0), 0)
  return { percent: (sumPercent * 100).toFixed(1), work: sumWork.toFixed(2) }
})

const timeTotal = computed(() => {
  if (!result.value) return { percent: '0', time: '0' }
  const sumPercent = workActivities.reduce((acc, a) => acc + (result.value.timeCoeffs?.[a] || 0), 0)
  const sumTime = workActivities.reduce((acc, a) => acc + (result.value.timeChars?.[a] || 0), 0)
  return { percent: (sumPercent * 100).toFixed(1), time: sumTime.toFixed(2) }
})

const budgetTotal = computed(() => {
  if (!result.value) return { percent: '0', work: '0', budget: '0' }
  const sumPercent = budgetActivities.reduce((acc, a) => acc + (result.value.budgetCoeffs?.[a] || 0), 0)
  const sumWork = budgetActivities.reduce((acc, a) => acc + (result.value.budgetWork?.[a] || 0), 0)
  const sumBudget = budgetActivities.reduce((acc, a) => acc + (result.value.budgetChars?.[a] || 0), 0)
  return { percent: (sumPercent * 100).toFixed(1), work: sumWork.toFixed(2), budget: sumBudget.toFixed(2) }
})
</script>

<template>
  <div class="flex-row">
    <div class="input-panel">
      <div class="form-group">
        <div class="row-2">
          <div class="form-field">
            <label>{{ LABELS.input.kloc }}</label>
            <input v-model="kloc" type="number" min="1" step="1" class="input" />
          </div>
          <div class="form-field">
            <label>{{ LABELS.input.budget }}</label>
            <input v-model="budgetPerWork" type="number" min="0.01" step="0.01" class="input" />
          </div>
        </div>

        <div class="form-field">
          <label>{{ LABELS.mode.label }}</label>
          <select v-model="mode" class="input">
            <option v-for="m in modes" :key="m.value" :value="m.value">{{ m.label }}</option>
          </select>
        </div>

        <div v-for="group in attributeGroups" :key="group.key" class="attr-group">
          <h3 class="group-title">{{ group.label }}</h3>
          <div class="attr-grid">
            <div v-for="attr in group.attrs" :key="attr" class="form-field">
              <label class="text-left">{{ LABELS.attributes[attr] }}</label>
              <select v-model="attributes[group.key][attr]" class="input">
                <option v-for="level in levels" :key="level" :value="level">{{ LABELS.levels[level] }}</option>
              </select>
            </div>
          </div>
        </div>

        <button @click="calculate" :disabled="loading" class="calc-btn">
          {{ loading ? '...' : LABELS.input.calculate }}
        </button>
      </div>
    </div>

    <div class="results-panel">
      <h2 class="results-title">{{ LABELS.results.title }}</h2>
      
      <div v-if="!result" class="no-result">
        Нажмите "Рассчитать" для получения результатов
      </div>

      <div v-else class="results-content">
        <div class="result-section">
          <h3 class="section-title">{{ LABELS.results.workAndTime }}</h3>
          <div class="table-wrapper">
            <table class="result-table">
              <thead>
                <tr>
                  <th>{{ LABELS.results.activity }}</th>
                  <th class="text-right">{{ LABELS.results.workPercent }}</th>
                  <th class="text-right">{{ LABELS.results.workMonths }}</th>
                  <th class="text-right">{{ LABELS.results.timePercent }}</th>
                  <th class="text-right">{{ LABELS.results.timeMonths }}</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="activity in workActivities" :key="activity">
                  <td>{{ LABELS.activities[activity] }}</td>
                  <td class="text-right">{{ getWorkRow(activity).percent }}</td>
                  <td class="text-right">{{ getWorkRow(activity).work }}</td>
                  <td class="text-right">{{ getTimeRow(activity).percent }}</td>
                  <td class="text-right">{{ getTimeRow(activity).time }}</td>
                </tr>
                <tr class="total-row">
                  <td>{{ LABELS.results.total }}</td>
                  <td class="text-right">{{ workTotal.percent }}</td>
                  <td class="text-right">{{ workTotal.work }}</td>
                  <td class="text-right">{{ timeTotal.percent }}</td>
                  <td class="text-right">{{ timeTotal.time }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div class="result-section">
          <h3 class="section-title">{{ LABELS.results.budget }}</h3>
          <div class="table-wrapper">
            <table class="result-table">
              <thead>
                <tr>
                  <th>{{ LABELS.results.activity }}</th>
                  <th class="text-right">{{ LABELS.results.budgetPercent }}</th>
                  <th class="text-right">{{ LABELS.results.workMonths }}</th>
                  <th class="text-right">{{ LABELS.results.budgetAmount }}</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="activity in budgetActivities" :key="activity">
                  <td>{{ activity === 'configManagement' ? LABELS.activities.configManagement : LABELS.activities[activity] }}</td>
                  <td class="text-right">{{ getBudgetRow(activity).percent }}</td>
                  <td class="text-right">{{ getBudgetRow(activity).work }}</td>
                  <td class="text-right">{{ getBudgetRow(activity).budget }}</td>
                </tr>
                <tr class="total-row">
                  <td>{{ LABELS.results.total }}</td>
                  <td class="text-right">{{ budgetTotal.percent }}</td>
                  <td class="text-right">{{ budgetTotal.work }}</td>
                  <td class="text-right">{{ budgetTotal.budget }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>