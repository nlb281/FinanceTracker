<script setup>
defineProps({
  mode: {
    type: String,
    required: true,
  },
  items: {
    type: Array,
    default: () => [],
  },
  transactionType: {
    type: String,
    default: 'income',
  },
  selectedCategoryName: {
    type: String,
    default: 'All categories',
  },
})

const emit = defineEmits(['change-type', 'change-category-name'])

const setType = (type) => {
  emit('change-type', type)
  emit('change-category-name', 'All categories')
}

const setCategoryName = (name) => {
  emit('change-category-name', name)
}
</script>

<template>
  <div class="transactions-filters">
    <div class="transactions-filters__type">
      <button
        class="transactions-filters__button"
        :class="{ 'transactions-filters__button--active': transactionType === 'income' }"
        @click="setType('income')"
      >
        Incomes
      </button>

      <button
        class="transactions-filters__button"
        :class="{ 'transactions-filters__button--active': transactionType === 'expense' }"
        @click="setType('expense')"
      >
        Expenses
      </button>
    </div>

    <div v-if="mode === 'transactions'" class="transactions-filters__categories">
      <button
        v-for="item in items"
        :key="item.name"
        class="transactions-filters__button"
        :class="{ 'transactions-filters__button--active': item.name === selectedCategoryName }"
        @click="setCategoryName(item.name)"
      >
        {{ item.name }}
      </button>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.transactions-filters {
  display: grid;
  gap: 14px;
  margin-bottom: 20px;

  &__type,
  &__categories {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
  }

  &__button {
    border: 2px solid $gray;
    color: $gray;
    border-radius: 999px;
    padding: 8px 14px;
    font-size: 14px;
    font-weight: 600;
    transition-duration: $transition-duration;

    &--active {
      color: $white;
      border-color: $white;
    }
  }
}
</style>
