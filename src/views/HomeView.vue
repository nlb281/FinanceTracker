<script setup>
import { computed, inject } from 'vue'
import BalanceOverview from '@/components/BalanceOverview.vue'

const balance = computed(() =>
  transactions.value.reduce(
    (sum, t) => sum + (t.type === 'expense' ? -Number(t.amount) : Number(t.amount)),
    0,
  ),
)

const income = computed(() =>
  transactions.value.reduce((sum, t) => sum + (t.type === 'income' ? Number(t.amount) : 0), 0),
)

const expense = computed(() =>
  transactions.value.reduce((sum, t) => sum + (t.type === 'expense' ? -Number(t.amount) : 0), 0),
)

const { transactions } = inject('recordsStore')
</script>

<template>
  <section class="home">
    <div class="home__inner container">
      <div class="home__header">
        <h2 class="home__title">Home</h2>
      </div>
      <div class="home__body">
        <BalanceOverview :balance="balance" :income="income" :expense="expense" />
      </div>
    </div>
  </section>
</template>

<style lang="scss" scoped>
.home {
  min-height: 100vh;
  padding-block: 28px;

  &__header {
    margin-bottom: 20px;
  }

  &__title {
    font-size: 28px;
    font-weight: 700;
  }
}
</style>
