<script setup>
defineProps({
  title: {
    type: String,
    default: 'Something went wrong',
  },
  message: {
    type: String,
    default: '',
  },
  closable: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(['close'])
</script>

<template>
  <section v-if="message" class="ui-error">
    <div class="ui-error__icon">!</div>

    <div class="ui-error__content">
      <p class="ui-error__title">{{ title }}</p>
      <p class="ui-error__message">{{ message }}</p>

      <div class="ui-error__actions">
        <slot name="actions" />
        <button v-if="closable" type="button" class="ui-error__button" @click="emit('close')">
          Close
        </button>
      </div>
    </div>
  </section>
</template>

<style lang="scss" scoped>
.ui-error {
  display: flex;
  gap: 12px;
  align-items: flex-start;
  margin-block: 16px;
  padding: 16px;
  border: 1px solid rgba($expense-color, 0.45);
  border-radius: 14px;
  background: rgba($expense-color, 0.08);

  &__icon {
    flex-shrink: 0;
    width: 22px;
    height: 22px;
    border: 1px solid rgba($expense-color, 0.55);
    border-radius: 50%;
    display: grid;
    place-items: center;
    color: $expense-color;
    font-weight: 700;
  }

  &__content {
    flex: 1;
    display: grid;
    gap: 6px;
  }

  &__title {
    color: $expense-color;
    font-size: 15px;
    font-weight: 700;
  }

  &__message {
    color: $white;
  }

  &__details {
    color: $gray;
    font-size: 13px;
  }

  &__actions {
    display: flex;
    gap: 8px;
    flex-wrap: wrap;
    margin-top: 2px;
  }

  &__button {
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
