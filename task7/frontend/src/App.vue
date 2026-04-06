<script setup>
import { ref } from 'vue'
import FunctionalPointsTab from './components/FunctionalPointsTab.vue'
import CocomoTab from './components/CocomoTab.vue'

const currentTab = ref('functional')

const setTab = (tab) => {
  currentTab.value = tab
}
</script>

<template>
  <div class="main-container">
    <div class="tabs">
      <button
        :class="currentTab === 'functional' ? 'tab-btn tab-active' : 'tab-btn tab-inactive'"
        @click="setTab('functional')"
      >
        Функциональные точки
      </button>
      <button
        :class="currentTab === 'cocomo' ? 'tab-btn tab-active' : 'tab-btn tab-inactive'"
        @click="setTab('cocomo')"
      >
        COCOMO II
      </button>
    </div>
    <FunctionalPointsTab v-if="currentTab === 'functional'" />
    <CocomoTab v-if="currentTab === 'cocomo'" />
  </div>
</template>

<style scoped>
</style>
