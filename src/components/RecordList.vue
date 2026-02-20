<script setup>
import RecordItem from '@/components/RecordItem.vue'

defineProps({
  mode: {
    type: String,
    required: true,
  },
  items: {
    type: Array,
    default: () => [],
  },
})
</script>

<template>
  <div class="record-list">
    <RecordItem v-for="item in items" :key="item.id || item.category" :mode="mode" :item="item" />
  </div>
</template>

<style lang="scss" scoped>
.record-list {
  border-top: 1px solid $line-color;
}
</style>
