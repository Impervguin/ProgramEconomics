<script setup>
import { ref, reactive } from 'vue'
import { CalculateFunctionalPoints } from '../../wailsjs/go/main/App.js'
import { PARAM_LEVELS, PARAM_LEVEL_LABELS, LANGUAGES, ELEMENTS, ELEMENT_LEVELS, ELEMENT_LEVEL_LABELS } from '../constants.js'

const params = reactive({
  DataPassing: PARAM_LEVELS.RandomInfluence,
  DistributedProcessing: PARAM_LEVELS.RandomInfluence,
  Efficiency: PARAM_LEVELS.RandomInfluence,
  ExpluatationRequirements: PARAM_LEVELS.RandomInfluence,
  TransacrionFrequency: PARAM_LEVELS.RandomInfluence,
  FastIO: PARAM_LEVELS.RandomInfluence,
  UserEffectiveness: PARAM_LEVELS.RandomInfluence,
  UpdateEfficiency: PARAM_LEVELS.RandomInfluence,
  ProcessingComplexity: PARAM_LEVELS.RandomInfluence,
  Reusability: PARAM_LEVELS.RandomInfluence,
  InstallationLightness: PARAM_LEVELS.RandomInfluence,
  ExpluatationLightness: PARAM_LEVELS.RandomInfluence,
  PlatformDiversity: PARAM_LEVELS.RandomInfluence,
  Maintainability: PARAM_LEVELS.RandomInfluence,
})

const languages = reactive({})
for (const lang in LANGUAGES) {
  languages[lang] = 0
}

const elements = reactive({
  externalInput: { Low: 0, Normal: 0, High: 0 },
  externalOutput: { Low: 0, Normal: 0, High: 0 },
  externalQuery: { Low: 0, Normal: 0, High: 0 },
  internalLogicalFile: { Low: 0, Normal: 0, High: 0 },
  externalInterfaceFile: { Low: 0, Normal: 0, High: 0 },
})

const result = ref(null)

const calculate = async () => {
  const dto = {
    params,
    languages,
    elements: Object.keys(elements).flatMap(key =>
      Object.keys(elements[key]).map(level => ({
        element: key,
        level: level,
        count: elements[key][level],
      }))
    ).filter(e => e.count > 0),
  }
  result.value = await CalculateFunctionalPoints(dto)
}

const getLanguagePercentSum = () => {
  return Object.values(languages).reduce((sum, val) => sum + parseFloat(val || 0), 0)
}
</script>

<template>
  <div class="fp-layout">
    <!-- LEFT (1/3) -->
    <div class="left-panel">
      <div class="attr-group">
        <div class="group-title">Характеристики проекта</div>
        <div class="attr-grid">
          <div v-for="(level, key) in params" :key="key" class="form-field">
            <label>{{ key }}</label>
            <select v-model="params[key]" class="input">
              <option v-for="(label, val) in PARAM_LEVEL_LABELS" :key="val" :value="val">
                {{ label }}
              </option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <!-- RIGHT (2/3) -->
    <div class="right-panel">
      
      <!-- LANGUAGES -->
      <div class="attr-group">
        <div class="group-title">Процент использования языков</div>
        <div class="lang-grid">
          <div v-for="(lang, key) in LANGUAGES" :key="key" class="form-field">
            <label>{{ lang }}</label>
            <input v-model.number="languages[key]" type="number" class="input" min="0" max="100" step="0.1" />
          </div>
        </div>
        <div class="form-field">
          <label>Сумма: {{ getLanguagePercentSum().toFixed(1) }}%</label>
        </div>
      </div>

      <!-- ELEMENTS -->
      <div class="attr-group">
        <div class="group-title">Элементы ФТ</div>
        <table class="result-table">
          <thead>
            <tr>
              <th>Элемент</th>
              <th v-for="level in ELEMENT_LEVELS" :key="level">
                {{ ELEMENT_LEVEL_LABELS[level] }}
              </th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(elem, key) in ELEMENTS" :key="key">
              <td>{{ elem }}</td>
              <td v-for="level in ELEMENT_LEVELS" :key="level">
                <input v-model.number="elements[key][level]" type="number" class="input" min="0" />
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- RESULT -->
      <div class="attr-group">
        <button @click="calculate"
                :disabled="calculating || getLanguagePercentSum() !== 100"
                class="calc-btn">
          Рассчитать
        </button>

        <div v-if="result" class="results-content">
          <table class="result-table">
            <tbody>
              <tr>
                <td>Функциональные точки</td>
                <td class="text-right">{{ result.points.toFixed(2) }}</td>
              </tr>
              <tr>
                <td>Функциональные точки после корректировки</td>
                <td class="text-right">{{ result.correctedPoints.toFixed(2) }}</td>
              </tr>
              <tr>
                <td>Lines of Code</td>
                <td class="text-right">{{ result.loc.toFixed(2) }}</td>
              </tr>
            </tbody>
          </table>
        </div>

        <div v-else class="no-result">Нет результатов</div>
      </div>

    </div>
  </div>
</template>

<style scoped>
</style>