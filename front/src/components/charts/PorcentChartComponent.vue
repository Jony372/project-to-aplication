<template>
  <h4 class="text-2xl mb-3 font-bold text-orange-300">{{title}}</h4>
  <div>
    <label for="gri" class="block mb-2 text-sm font-medium text-gray-900 ">Select a ticker</label>
    <select @change="update($event)" id="countries" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5">
      <option v-for="(item, i) in store.items(growth)" :selected="i==0" :value="i">{{item.ticker}}</option>
    </select>
  </div>
  <apexchart type="radialBar" height="350" :options="chartOptions" :series="series"></apexchart>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useChartDataStore } from '../../stores/chartDataStore';

const props = defineProps<{
  title: string;
  growth: boolean
}>();

const store = useChartDataStore();

onMounted(async() => {
  await store.getItems();
  update()
});

var chartOptions = ref({
  chart: {
    height: 350,
    type: 'radialBar'
  },
  plotOptions: {
    radialBar: {
      hollow: {
        size: '70%',
      }
    },
  },
  colors: ['red'],
  labels: ref([""]),
});
const series = ref([0]);

const update = (i?: any) => {
  const index = i?.target.value || 0;
  const g = props.growth;
  store.selectItem(index, g);
  series.value = [ Number(store.getSelectedItem(g)?.porcent.toFixed(2)) || 0];
  chartOptions.value = {
  chart: {
    height: 350,
    type: 'radialBar'
  },
  plotOptions: {
    radialBar: {
      hollow: {
        size: '70%',
      }
    },
  },
  colors: [g?'#07fa34':'red'],
  labels: ref([store.getSelectedItem(g)?.ticker || ""]),
}
}

</script>