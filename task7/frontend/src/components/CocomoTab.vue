<script setup>
import { ref, reactive } from 'vue'
import { CalculateCompositionModel, CalculateArchitectureModel } from '../../wailsjs/go/main/App.js'
import { TEAM_PARAM_LEVELS, TEAM_PARAM_LABELS, PRODUCTIVITY_LEVELS, PRODUCTIVITY_LABELS, ARCHITECTURE_ATTRIBUTES, ATTRIBUTE_LEVELS, ATTRIBUTE_LABELS, VIEW_REPORT_LEVELS, VIEW_REPORT_LABELS } from '../constants.js'

const teamConfig = reactive({
  TeamParameterLevel: 'UnifiedWhole',
  PmatParameterLevel: 'CMMLevel5',
  PrecParameterLevel: 'FullyKnownProject',
  ReslParameterLevel: 'Full100',
  FlexParameterLevel: 'GeneralGoalsOnly',
})

const composition = reactive({
  views: { Simple: 0, Medium: 0, Complex: 0 },
  reports: { Simple: 0, Medium: 0, Complex: 0 },
  modules: 0,
  reuse: 0,
  productivity: 'MediumProductivity',
  averagePay: 0,
})

const architecture = reactive({
  attributes: {
    PERS: 'Nominal',
    RCPX: 'Nominal',
    RUSE: 'Nominal',
    PDIF: 'Nominal',
    PREX: 'Nominal',
    FCIL: 'Nominal',
    SCED: 'Nominal',
  },
  kloc: 0,
  averagePay: 0,
})

const compositionResult = ref(null)
const architectureResult = ref(null)

const calculateComposition = async () => {
  const dto = {
    Elements: [
      ...Object.keys(composition.views).map(level => ({
        element: 'ViewElement',
        level: level,
        count: composition.views[level],
      })),
      ...Object.keys(composition.reports).map(level => ({
        element: 'ReportElement',
        level: level,
        count: composition.reports[level],
      })),
    ].filter(e => e.count > 0),
    ModuleCount: composition.modules,
    ReuseParam: composition.reuse,
    ProductivityLevel: composition.productivity,
    AveragePay: composition.averagePay,
    TeamConfig: teamConfig,
  }
  compositionResult.value = await CalculateCompositionModel(dto)
}

const calculateArchitecture = async () => {
  const dto = {
    ArchitectureAttributes: architecture.attributes,
    TeamConfig: teamConfig,
    Kloc: architecture.kloc,
    AveragePay: architecture.averagePay,
  }
  architectureResult.value = await CalculateArchitectureModel(dto)
}
</script>

<template>
  <div class="form-group">
    <div class="attr-group">
      <div class="group-title">Характеристики команды</div>
      <div class="row-2">
        <div v-for="(levels, param) in TEAM_PARAM_LEVELS" :key="param" class="form-field">
          <label>{{ param }}</label>
          <select v-model="teamConfig[param + 'ParameterLevel']" class="input">
            <option v-for="(level, key) in levels" :key="key" :value="key">
              {{ TEAM_PARAM_LABELS[param][key] }}
            </option>
          </select>
        </div>
      </div>
    </div>
    <div class="models-grid">

      <!-- COMPOSITION -->
      <div class="model">
        <div class="attr-group">
          <div class="group-title">Модель композиции</div>

          <div class="form-field">
            <label class="group-label">Экранные формы</label>
            <div class="adaptive-grid">
              <div v-for="level in VIEW_REPORT_LEVELS" :key="level">
                <label>{{ VIEW_REPORT_LABELS[level] }}</label>
                <input v-model.number="composition.views[level]" type="number" class="input" />
              </div>
            </div>
          </div>

          <div class="form-field">
            <label class="group-label">Отчёты</label>
            <div class="adaptive-grid">
              <div v-for="level in VIEW_REPORT_LEVELS" :key="level">
                <label>{{ VIEW_REPORT_LABELS[level] }}</label>
                <input v-model.number="composition.reports[level]" type="number" class="input" />
              </div>
            </div>
          </div>

          <div class="adaptive-grid">
            <div class="form-field">
              <label>Модули</label>
              <input v-model.number="composition.modules" type="number" class="input" />
            </div>

            <div class="form-field">
              <label>Reuse (%)</label>
              <input v-model.number="composition.reuse" type="number" class="input" />
            </div>

            <div class="form-field">
              <label>Средняя зарплата</label>
              <input v-model.number="composition.averagePay" type="number" class="input" />
            </div>

            <div class="form-field">
              <label>Опытность команды</label>
              <select v-model="composition.productivity" class="input">
                <option v-for="(label, val) in PRODUCTIVITY_LABELS" :key="val" :value="val">
                  {{ label }}
                </option>
              </select>
            </div>
          </div>

          <button @click="calculateComposition" class="calc-btn">Рассчитать</button>
        </div>

       <div v-if="compositionResult" class="results-content">
        <table class="result-table">
          <tbody>
            <tr>
              <td>Трудозатраты (чел.-мес.)</td>
              <td class="text-right">{{ compositionResult.Work.toFixed(2) }}</td>
            </tr>
            <tr>
              <td>Время (мес.)</td>
              <td class="text-right">{{ compositionResult.Time.toFixed(2) }}</td>
            </tr>
            <tr>
              <td>Стоимость</td>
              <td class="text-right">{{ compositionResult.Cost.toFixed(2) }}</td>
            </tr>
            <tr>
              <td>Размер команды</td>
              <td class="text-right">{{ compositionResult.TeamSize }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      </div>

      <!-- ARCH -->
      <div class="model">
        <div class="attr-group">
          <div class="group-title">Архитектура</div>

          <div class="adaptive-grid">
            <div v-for="attr in ARCHITECTURE_ATTRIBUTES" :key="attr">
              <label>{{ attr }}</label>
              <select v-model="architecture.attributes[attr]" class="input">
                <option v-for="(label, val) in ATTRIBUTE_LABELS" :key="val" :value="val">
                  {{ label }}
                </option>
              </select>
            </div>
          </div>

          <div class="adaptive-grid">
            <input v-model.number="architecture.kloc" type="number" class="input" placeholder="KLOC"/>
            <input v-model.number="architecture.averagePay" type="number" class="input" placeholder="ЗП"/>
          </div>

          <button @click="calculateArchitecture" class="calc-btn">Рассчитать</button>
        </div>

        <div v-if="architectureResult" class="results-content">
          <table class="result-table">
            <tbody>
              <tr>
                <td>Трудозатраты (чел.-мес.)</td>
                <td class="text-right">{{ architectureResult.Work.toFixed(2) }}</td>
              </tr>
              <tr>
                <td>Время (мес.)</td>
                <td class="text-right">{{ architectureResult.Time.toFixed(2) }}</td>
              </tr>
              <tr>
                <td>Стоимость</td>
                <td class="text-right">{{ architectureResult.Cost.toFixed(2) }}</td>
              </tr>
              <tr>
                <td>Размер команды</td>
                <td class="text-right">{{ architectureResult.TeamSize }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

    </div>
  </div>
</template>

<style scoped>
</style>