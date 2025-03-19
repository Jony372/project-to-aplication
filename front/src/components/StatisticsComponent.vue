<template>
  <h2 class="text-5xl font-bold mb-4 text-orange-300">Ticker's statistics</h2>
  <div class="h-screen md:grid md:grid-cols-2 gap-4 flex flex-col">
    <div class="md:col-span-2 ">
      <BarChartComponent/>
    </div>
    <div class="md:col-span-1 ">
      <PorcentChartComponent title="Growth percentage" :growth="true"/>
    </div>
    <div class="md:col-span-1 ">
      <PorcentChartComponent title="Percentage decrease" :growth="false"/>
    </div>
  </div>
</template>

<script setup lang="ts">
import PorcentChartComponent from './charts/PorcentChartComponent.vue';
import BarChartComponent from './charts/BarChartComponent.vue';
</script>