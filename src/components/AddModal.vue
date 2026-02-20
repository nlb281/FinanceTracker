<script setup>
import { reactive, computed, inject, ref, watch } from 'vue'
import UiError from '@/components/UiError.vue'
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },
  mode: {
    type: String,
    required: true,
  },
  transactionType: {
    type: String,
    default: 'income',
  },
})

const transactionForm = reactive({
  amount: null,
  type: 'income',
  category_id: null,
  date: '',
  description: '',
})

const categoryForm = reactive({
  name: '',
  type: 'income',
})

watch(
  () => props.transactionType,
  (type) => {
    transactionForm.type = type
    categoryForm.type = type
  },
  { immediate: true },
)

const payload = computed(() => {
  return props.mode === 'categories' ? categoryForm : transactionForm
})

const title = computed(() => {
  return props.mode === 'categories' ? 'Add category' : 'Add transaction'
})

const emit = defineEmits(['update:modelValue'])
const errorMessage = ref('')

const closeModal = () => {
  errorMessage.value = ''
  emit('update:modelValue', false)
}

const handleCreateRecord = async () => {
  errorMessage.value = ''
  const result = await createRecord(payload.value, props.mode)

  if (result.status) {
    closeModal()
    return
  }

  errorMessage.value = result.message
}

const categoryOptions = computed(() =>
  categories.value
    .filter((c) => c.type === props.transactionType)
    .map((c) => ({ id: c.id, name: c.name })),
)

const { createRecord } = inject('recordsActions')
const { categories } = inject('items')
</script>

<template>
  <div v-if="modelValue" class="modal" @click.self="closeModal">
    <div class="modal__content">
      <div class="modal__header">
        <h3 class="modal__title">{{ title }}</h3>
        <button type="button" class="modal__close" @click="closeModal"></button>
      </div>
      <UiError title="Create error" :message="errorMessage" @close="errorMessage = ''" />

      <form v-if="mode === 'categories'" class="modal__form" @submit.prevent="handleCreateRecord">
        <input
          name="name"
          class="modal__input"
          type="text"
          placeholder="Name"
          v-model="categoryForm.name"
          required
        />
        <select
          name="type"
          class="modal__input modal__select"
          v-model="categoryForm.type"
          disabled
          required
        >
          <option value="income">Income</option>
          <option value="expense">Expense</option>
        </select>

        <div class="modal__actions">
          <button type="button" class="modal__button" @click="closeModal">Cancel</button>
          <button type="submit" class="modal__button">Add</button>
        </div>
      </form>

      <form v-else class="modal__form" @submit.prevent="handleCreateRecord">
        <select
          name="type"
          class="modal__input modal__select"
          v-model.number="transactionForm.category_id"
          required
        >
          <option v-for="category in categoryOptions" :key="category.id" :value="category.id">
            {{ category.name }}
          </option>
        </select>
        <input
          name="description"
          class="modal__input"
          type="text"
          placeholder="Description"
          v-model="transactionForm.description"
        />
        <select
          name="type"
          class="modal__input modal__select"
          v-model="transactionForm.type"
          disabled
          required
        >
          <option value="income">Income</option>
          <option value="expense">Expense</option>
        </select>
        <input
          name="amount"
          class="modal__input"
          type="number"
          placeholder="Amount"
          v-model.number="transactionForm.amount"
          required
        />
        <input
          name="date"
          class="modal__input"
          type="date"
          v-model="transactionForm.date"
          required
        />

        <div class="modal__actions">
          <button type="button" class="modal__button" @click="closeModal">Cancel</button>
          <button type="submit" class="modal__button">Add</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style lang="scss" scoped>
.modal {
  position: fixed;
  inset: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 16px;
  background-color: rgba(0, 0, 0, 0.55);

  &__content {
    width: 440px;
    padding: 24px 16px;
    border: 1px solid $line-color;
    border-radius: 14px;
    background-color: $background;
  }

  &__header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 16px;
  }

  &__title {
    font-size: 20px;
    font-weight: 600;
  }

  &__close {
    position: relative;
    width: 30px;
    height: 30px;
    border: 1px solid $line-color;
    border-radius: 50%;
    color: $white;
    font-size: 18px;
    transition-duration: $transition-duration;

    &::before {
      content: '';
      position: absolute;
      width: 50%;
      height: 2px;
      top: 50%;
      left: 50%;
      translate: -50% -50%;
      rotate: 45deg;
      background-color: $white;
    }

    &::after {
      content: '';
      position: absolute;
      width: 50%;
      height: 2px;
      top: 50%;
      left: 50%;
      translate: -50% -50%;
      rotate: -45deg;
      background-color: $white;
    }

    &:hover {
      border-color: $white;
    }
  }

  &__form {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  &__input {
    padding: 12px 10px;
    border: 1px solid $line-color;
    border-radius: 10px;
    outline: none;
    background-color: $background;
    color: $white;
  }

  &__select {
    appearance: none;
  }

  &__actions {
    margin-top: 4px;
    display: flex;
    justify-content: space-between;
    gap: 8px;
  }

  &__button {
    width: 100%;
    padding: 8px 10px;
    border: 1px solid $line-color;
    border-radius: 50px;
    color: $white;
    transition-duration: $transition-duration;

    &:hover {
      border-color: $white;
    }
  }
}
</style>
