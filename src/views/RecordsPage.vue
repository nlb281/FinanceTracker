<script setup>
import { computed, inject, ref, provide } from 'vue'
import { useRoute } from 'vue-router'
import { deleteTransaction, createTransaction } from '@/api/transactions.api'
import { deleteCategory, createCategories } from '@/api/categories.api'
import RecordList from '@/components/RecordList.vue'
import RecordFilters from '@/components/RecordFilters.vue'
import AddModal from '@/components/AddModal.vue'
import EmptyState from '@/components/EmptyState.vue'
import UiError from '@/components/UiError.vue'

const route = useRoute()
const mode = computed(() => route.meta.mode)
const transactionType = ref('income')
const selectedCategoryName = ref('All categories')
const isAddModalOpen = ref(false)
const categoryDeleteError = ref('')
const { transactions, categories } = inject('recordsStore')

const transactionsItems = computed(() => {
  if (!transactions.value) return []

  return transactions.value
    .filter((t) => t.type === transactionType.value)
    .map((t) => ({
      ...t,
      category_name: categories.value.find((c) => c.id === t.category_id)?.name,
    }))
    .filter(
      (t) =>
        selectedCategoryName.value === 'All categories' ||
        t.category_name === selectedCategoryName.value,
    )
})

const filterCategories = computed(() => [
  { name: 'All categories' },
  ...categories.value
    .filter((c) => c.type === transactionType.value)
    .map((c) => ({ name: c.name })),
])

const categoriesItems = computed(() => {
  const byCategory = {}

  if (categories.value.length === 0) {
    return []
  }

  if (transactions.value.length === 0 && categories.value.length > 0) {
    return categories.value
      .filter((c) => c.type === transactionType.value)
      .map((c) => ({
        ...c,
        total: 0,
      }))
  }

  for (const t of transactions.value) {
    const categoryId = t.category_id
    const amount = Number(t.amount)

    if (!byCategory[categoryId]) {
      byCategory[categoryId] = {
        total: 0,
      }
    }
    byCategory[categoryId].total += amount
  }

  return categories.value
    .filter((c) => c.type === transactionType.value)
    .map((c) => ({
      id: c.id,
      name: c.name,
      type: c.type,
      total: byCategory[c.id]?.total || 0,
    }))
})

const items = computed(() =>
  mode.value === 'categories' ? categoriesItems.value : transactionsItems.value,
)

const title = computed(() => {
  return mode.value === 'categories' ? 'Categories' : 'Transactions'
})

const addButtonText = computed(() => {
  return mode.value === 'categories' ? 'Add category' : 'Add transaction'
})

const openAddModal = () => {
  isAddModalOpen.value = true
}

const createRecord = async (payload, entityType) => {
  try {
    if (entityType === 'categories') {
      const data = await createCategories(payload)
      categories.value.push({
        id: data.id,
        name: payload.name,
        type: payload.type,
      })
    } else {
      const data = await createTransaction(payload)
      transactions.value.push({
        id: data.id,
        amount: payload.amount,
        type: payload.type,
        category_id: payload.category_id,
        date: payload.date,
        description: payload.description,
      })
      transactionType.value = payload.type
    }

    return { status: true }
  } catch (error) {
    return { status: false, message: error.message }
  }
}

// Remove Transaction

const removeTransaction = async (transactionId) => {
  const prevTransactions = [...transactions.value]

  try {
    await deleteTransaction(transactionId)

    transactions.value = transactions.value.filter((t) => t.id !== transactionId)
  } catch (error) {
    transactions.value = prevTransactions
    console.error('Delete transaction error:', error)
  }
}

// Remove Category

const removeCategory = async (categoryId) => {
  categoryDeleteError.value = ''
  const prevCategories = [...categories.value]
  try {
    await deleteCategory(categoryId)

    categories.value = categories.value.filter((c) => c.id !== categoryId)
  } catch (error) {
    categories.value = prevCategories
    categoryDeleteError.value = error?.message
  }
}

provide('items', {
  categoriesItems,
  filterCategories,
  categories,
})

provide('recordsActions', {
  createRecord,
  removeTransaction,
  removeCategory,
})
</script>

<template>
  <section class="records-page container">
    <header class="records-page__header">
      <h2 class="records-page__title">{{ title }}</h2>
      <button class="records-page__add-button" @click="openAddModal">{{ addButtonText }}</button>
    </header>
    <div class="records-page__body">
      <AddModal v-model="isAddModalOpen" :mode="mode" :transaction-type="transactionType" />
      <RecordFilters
        :mode="mode"
        :items="filterCategories"
        :transaction-type="transactionType"
        :selected-category-name="selectedCategoryName"
        @change-category-name="(name) => (selectedCategoryName = name)"
        @change-type="(type) => (transactionType = type)"
      />
      <RecordList v-if="items?.length" :mode="mode" :items="items" />
      <EmptyState v-else :mode="mode" :items="items" @add="openAddModal" />
      <UiError
        v-if="mode === 'categories'"
        title="Delete error"
        :message="categoryDeleteError"
        closable
        @close="categoryDeleteError = ''"
      />
    </div>
  </section>
</template>
