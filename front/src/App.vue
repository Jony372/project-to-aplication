<script setup lang="ts">
  import NavBarComponent from './components/NavBarComponent.vue';
</script>

<template>
  <NavBarComponent/>
  <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
    <router-view/>
  </div>
</template>

<style scoped>
</style>
