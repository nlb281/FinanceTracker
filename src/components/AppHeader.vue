<script setup>
import { reactive } from 'vue'

const menuItems = reactive([
  {
    page: 'Home',
    path: '/',
  },
  {
    page: 'Transactions',
    path: '/transactions',
  },
  {
    page: 'Categories',
    path: '/categories',
  },
])
</script>

<template>
  <header class="header">
    <div class="header__inner container">
      <img class="header__image" src="/favicons/money.png" alt="Money" width="96" height="96" />

      <h1 class="header__title">Finance Tracker</h1>
      <nav class="menu">
        <ul class="menu__list">
          <li v-for="item in menuItems" :key="item.page" class="menu__item">
            <RouterLink :to="item.path" class="menu__link" active-class="menu__link--active">{{
              item.page
            }}</RouterLink>
          </li>
        </ul>
      </nav>
    </div>
  </header>
</template>

<style lang="scss" scoped>
.header {
  &__inner {
    display: flex;
    flex-direction: column;
    row-gap: 50px;
    padding-block: 50px;
    position: relative;
  }

  &__title {
    font-size: 48px;
    font-weight: 700;
  }

  &__image {
    position: absolute;
    top: -18%;
  }
}

.menu {
  &__list {
    display: flex;
    column-gap: 30px;
  }

  &__link {
    font-size: 20px;
    font-weight: 500;
    color: $gray;
    transition-duration: $transition-duration;
    cursor: pointer;
    position: relative;

    &:not(.menu__link--active)::after {
      position: absolute;
      content: '';
      width: 100%;
      height: 1px;
      bottom: 0;
      left: 0;
      background-color: $gray;
    }

    &--active {
      color: $white;
      opacity: 1;
    }
  }
}
</style>
