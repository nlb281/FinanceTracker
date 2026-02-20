<script setup>
import { computed } from 'vue'
const props = defineProps({
  mode: {
    type: String,
    required: true,
  },
  items: {
    type: Array,
    default: () => [],
  },
})

const content = computed(() =>
  props.mode === 'categories'
    ? {
        title: 'No categories yet',
        text: 'It is empty here for now: add your first category and start tracking by category.',
        button: 'Add category',
      }
    : {
        title: 'No transactions yet',
        text: 'It is empty here for now: add your first transaction and start tracking.',
        button: 'Add transaction',
      },
)

const emit = defineEmits(['add'])
</script>

<template>
  <section class="empty-state">
    <div class="empty-state__content">
      <h3 class="empty-state__title">{{ content.title }}</h3>
      <p class="empty-state__text">
        {{ content.text }}
      </p>
    </div>
    <button class="empty-state__button" @click="emit('add')">{{ content.button }}</button>
  </section>
</template>

<style lang="scss" scoped>
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  row-gap: 20px;
  margin-top: 20px;
  padding: 24px 16px;
  border: 1px solid $line-color;
  border-radius: 14px;
  background-color: rgba($white, 0.02);

  &__content {
    display: flex;
    flex-direction: column;
    row-gap: 8px;
    text-align: center;
  }

  &__title {
    font-size: 20px;
    font-weight: 700;
    color: $white;
  }

  &__text {
    color: $gray;
  }

  &__button {
    display: inline-block;
    width: fit-content;
    color: $white;
    font-weight: 600;
    border-radius: 50px;
    border: 1px solid $line-color;
    background: transparent;
    padding: 12px 38px;
    transition-duration: $transition-duration;

    &:hover {
      background-color: $white;
      color: $background;
      border-color: $white;
    }
  }
}
</style>
