<script setup>
defineProps({
  balance: {
    type: [Number, String],
    default: 0,
  },
  income: {
    type: [Number, String],
    default: 0,
  },
  expense: {
    type: [Number, String],
    default: 0,
  },
})
</script>

<template>
  <section class="balance-overview" aria-label="Financial summary">
    <article class="balance-overview__card balance-overview__card--balance">
      <p class="balance-overview__label">Balance</p>
      <p class="balance-overview__value">{{ balance }}</p>
    </article>

    <article class="balance-overview__card balance-overview__card--income">
      <p class="balance-overview__label">Income</p>
      <p class="balance-overview__value balance-overview__value--income">{{ income }}</p>
    </article>

    <article class="balance-overview__card balance-overview__card--expense">
      <p class="balance-overview__label">Expense</p>
      <p class="balance-overview__value balance-overview__value--expense">{{ expense }}</p>
    </article>
  </section>
</template>

<style lang="scss" scoped>
.balance-overview {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 14px;
  margin-bottom: 24px;

  &__card {
    border: 1px solid $line-color;
    border-radius: 12px;
    background-color: rgba(255, 255, 255, 0.02);
    padding: 16px;
  }

  &__card--income {
    border-color: rgba($income-color, 0.45);
  }

  &__card--expense {
    border-color: rgba($expense-color, 0.45);
  }

  &__label {
    margin-bottom: 8px;
    font-size: 12px;
    text-transform: uppercase;
    color: $gray;
  }

  &__value {
    font-size: 26px;
    font-weight: 700;

    &--income {
      color: $income-color;
    }

    &--expense {
      color: $expense-color;
    }
  }
}
</style>
