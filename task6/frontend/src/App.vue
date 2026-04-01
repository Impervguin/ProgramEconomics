<script setup>
import { ref, computed } from 'vue'
import { LABELS } from './constants.js'
import CalculatorTab from './components/CalculatorTab.vue'
import DistributionTab from './components/DistributionTab.vue'
import ResearchTab from './components/ResearchTab.vue'

const activeTab = ref('calculator')
const calculationResult = ref(null)

const tabLabels = computed(() => LABELS.tabs)

function handleCalculation(result) {
  calculationResult.value = result
}
</script>

<template>
  <div class="main-container">
    <div class="tabs">
      <button
        v-for="(label, key) in tabLabels"
        :key="key"
        @click="activeTab = key"
        :class="['tab-btn', activeTab === key ? 'tab-active' : 'tab-inactive']"
      >
        {{ label }}
      </button>
    </div>

    <CalculatorTab v-show="activeTab === 'calculator'" @calculate="handleCalculation" />

    <DistributionTab v-show="activeTab === 'distribution'" :calculationResult="calculationResult" />

    <ResearchTab v-show="activeTab === 'research'" />
  </div>
</template>