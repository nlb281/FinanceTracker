<script setup>
import { ref, onMounted, provide } from 'vue'
import HeroBanner from '@/components/HeroBanner.vue'
import AppHeader from '@/components/AppHeader.vue'
import UiError from '@/components/UiError.vue'
import { getTransactions } from '@/api/transactions.api'
import { getCategories } from '@/api/categories.api'
import { RouterView } from 'vue-router'

const transactions = ref([])
const categories = ref([])
const loadError = ref('')

const loadData = async () => {
  loadError.value = ''

  try {
    const [transactionsData, categoriesData] = await Promise.all([
      getTransactions(),
      getCategories(),
    ])

    transactions.value = Array.isArray(transactionsData) ? transactionsData : []
    categories.value = Array.isArray(categoriesData) ? categoriesData : []
  } catch (error) {
    console.log('Failed to load data.', error)
    loadError.value = 'Failed to load data'
  }
}

onMounted(loadData)

provide('recordsStore', { transactions, categories })
</script>

<template>
  <div class="wrapper">
    <HeroBanner />
    <AppHeader />

    <div class="container">
      <UiError title="Loading error" :message="loadError">
        <template #actions>
          <button type="button" class="wrapper__retry-button" @click="loadData">Retry</button>
        </template>
      </UiError>
    </div>

    <RouterView />
  </div>
</template>

<style lang="scss" scoped>
.wrapper {
  background-color: $background;

  &__retry-button {
    border: 1px solid $line-color;
    color: $white;
    border-radius: 50px;
    padding: 6px 12px;
    font-size: 13px;
    font-weight: 600;
    transition-duration: $transition-duration;

    &:hover {
      border-color: $white;
    }
  }
}
</style>
