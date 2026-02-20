<script setup>
import { computed, inject } from 'vue'
const props = defineProps({
  mode: {
    type: String,
    required: true,
  },
  item: {
    type: Object,
    required: true,
  },
})

const amountClass = computed(() => {
  if (props.item.total === 0) return 'record-item__amount--empty'
  return props.item.type === 'income'
    ? 'record-item__amount--income'
    : 'record-item__amount--expense'
})

const amountSign = computed(() => {
  if (props.item.total === 0) return ''
  return props.item.type === 'income' ? '+' : '-'
})

const handleRemoveTransaction = async (id) => {
  await removeTransaction(id)
}

const handleRemoveCategory = async (id) => {
  await removeCategory(id)
}

const { removeTransaction, removeCategory } = inject('recordsActions')
</script>

<template>
  <div class="record-item">
    <template v-if="mode === 'transactions'">
      <div class="record-item__meta">
        <p v-if="item.description" class="record-item__description">{{ item.description }}</p>
        <p class="record-item__details">{{ item.category_name }} | {{ item.date }}</p>
      </div>

      <div class="record-item__actions">
        <p
          class="record-item__amount"
          :class="{
            'record-item__amount--income': item.type === 'income',
            'record-item__amount--expense': item.type === 'expense',
          }"
        >
          {{ item.type === 'income' ? '+' : '-' }}{{ Number(item.amount).toLocaleString() }} RUB
        </p>
        <button type="button" class="record-item__delete" @click="handleRemoveTransaction(item.id)">
          Delete
        </button>
      </div>
    </template>

    <template v-else>
      <div class="record-item__meta">
        <p class="record-item__name">{{ item.name }}</p>
        <p class="record-item__details">Total amount by category</p>
      </div>

      <div class="record-item__actions">
        <p class="record-item__amount" :class="amountClass">
          {{ amountSign }}{{ Number(item.total).toLocaleString() }} RUB
        </p>
        <button type="button" class="record-item__delete" @click="handleRemoveCategory(item.id)">
          Delete
        </button>
      </div>
    </template>
  </div>
</template>

<style lang="scss" scoped>
.record-item {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  padding-block: 16px;
  border-bottom: 1px solid $line-color;

  &__meta {
    display: flex;
    flex-direction: column;
    justify-content: center;
    flex: 1;
  }

  &__name {
    font-weight: 600;
    margin-bottom: 4px;
  }

  &__description,
  &__details {
    color: $gray;
    font-size: 14px;
  }

  &__description {
    margin-bottom: 4px;
  }

  &__actions {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  &__amount {
    font-weight: 700;

    &--income {
      color: $income-color;
    }

    &--expense {
      color: $expense-color;
    }

    &--empty {
      color: $white;
    }
  }

  &__delete {
    border: 1px solid $line-color;
    color: $expense-color;
    border-radius: 8px;
    padding: 6px 10px;
    font-size: 14px;
    font-weight: 600;
  }
}
</style>
