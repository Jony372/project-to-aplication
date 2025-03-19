<template>
  <div class="relative overflow-x-auto rounded-2xl">
    <table class="w-full text-sm text-left rtl:text-right text-gray-500">
      <thead class="text-xs text-gray-700 uppercase bg-gray-50">
        <tr>
          <th scope="col" class="px-6 py-3 whitespace-nowrap">
            TICKER
          </th>
          <th scope="col" class="px-6 py-3 whitespace-nowrap">
            COMPANY
          </th>
          <th scope="col" class="px-6 py-3 whitespace-nowrap">
            BROKERAGE
          </th>
          <th scope="col" class="px-6 py-3 whitespace-nowrap">
            ACTION
          </th>
          <th scope="col" class="px-6 py-3 whitespace-nowrap">
            RATING FROM
          </th>
          <th scope="col" class="px-6 py-3 whitespace-nowrap">
            RATING TO
          </th>
          <th scope="col" class="px-6 py-3 whitespace-nowrap">
            TARGET FROM
          </th>
          <th scope="col" class="px-6 py-3 whitespace-nowrap">
            TARGET TO
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="i in store.items" class="bg-white border-b  border-gray-200">
          <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap ">
            {{i.ticker}}
          </th>
          <td class="px-6 py-4 whitespace-nowrap">
            {{i.company}}
          </td>
          <td class="px-6 py-4 whitespace-nowrap">
            {{i.brokerage}}
          </td>
          <td class="px-6 py-4 whitespace-nowrap">
            {{i.action}}
          </td>
          <td class="px-6 py-4 whitespace-nowrap">
            {{i.rating_from}}
          </td>
          <td class="px-6 py-4 whitespace-nowrap">
            {{i.rating_to}}
          </td>
          <td class="px-6 py-4 whitespace-nowrap">
            {{i.target_from}}
          </td>
          <td class="px-6 py-4 whitespace-nowrap">
            {{i.target_to}}
          </td>
        </tr>
      </tbody>
    </table>
    <div class="relative">
      <p class="absolute left-0 pt-4 text-gray-400">Pagina: {{ store.pages }}</p>
      <nav class="flex justify-center items-center pt-4" aria-label="Table navigation">
        <button v-on:click="store.fetchItems(true)" :disabled="store.pages===1" class="disabled:cursor-auto cursor-pointer flex items-center justify-center px-3 h-8 ms-0 leading-tight text-gray-500 bg-white border border-gray-300 rounded-s-lg hover:bg-gray-100 hover:text-gray-700 ">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="M15.75 19.5 8.25 12l7.5-7.5" />
          </svg>
        </button>
        <button v-on:click="store.fetchItems()" class="cursor-pointer flex items-center justify-center px-3 h-8 leading-tight text-gray-500 bg-white border border-gray-300 rounded-e-lg hover:bg-gray-100 hover:text-gray-700 ">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
            <path stroke-linecap="round" stroke-linejoin="round" d="m8.25 4.5 7.5 7.5-7.5 7.5" />
          </svg>
        </button>
      </nav>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useItemStore } from '../stores/itemStore';

const store = useItemStore();

onMounted(() => {
  store.fetchItems();
});
// import { ref, onMounted } from "vue";
// import axios from "axios";
// import type { ApiDataInterface } from "../assets/interfaces";

// const data = ref<ApiDataInterface[]>([]);

// onMounted(async () => {
//   try {
//     const response = await axios.get("http://localhost:8080/api");
//     data.value = response.data.items;
//   } catch (error) {
//     console.error(error);
//   }
// });
</script>