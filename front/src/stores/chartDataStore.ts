import axios from "axios";
import { defineStore } from "pinia";
import { computed, ref } from "vue";

interface charItemInteface {
  id: string;
  ticker: string;
  porcent: number;
}

export const useChartDataStore = defineStore("chartData", () => {
  const API_URL = "http://localhost:8080/api/"
  const grpItems = ref<charItemInteface[]>([]);
  const dcpItems = ref<charItemInteface[]>([]);
  const grpItemSelected = ref<charItemInteface>();
  const dcpItemSelected = ref<charItemInteface>();
  const toBuyItems = ref<charItemInteface[]>([]);

  const getItems = async () => {
    try {
      const response = await axios.get(`${API_URL}grp`);
      grpItems.value = await response.data;
      const response2 = await axios.get(`${API_URL}dcp`);
      dcpItems.value = await response2.data;
      grpItemSelected.value = grpItems.value[0];
      dcpItemSelected.value = dcpItems.value[0];
    } catch (error) {
      console.log(error);
    }
  };

  const getToBuyItems = async () => {
    try {
      const response = await axios.get(`${API_URL}to-buy`);
      toBuyItems.value = await response.data;
    } catch (error) {
      console.log(error);
    }
  }

  const selectItem = (index: number, growth:boolean) => {
    if(growth){
      grpItemSelected.value = grpItems.value[index];
    }else{
      dcpItemSelected.value = dcpItems.value[index];
    }
  }

  const getSelectedItem = (growth:boolean):charItemInteface|undefined => {
    return growth ? grpItemSelected.value : dcpItemSelected.value;
  }

  const items = (growth:boolean):charItemInteface[] => {
    return growth ? grpItems.value : dcpItems.value;
  }

  return {
    grpI: computed(() => grpItems.value),
    dcpI: computed(() => dcpItems.value),
    getSelectedItem: computed(() => getSelectedItem),
    items: computed(() => items),
    getGrowthNames: computed(() => toBuyItems.value.map((item) => item.ticker)),
    getGrowthPorcent: computed(() => toBuyItems.value.map((item) => Number(item.porcent.toFixed(2)))),
    getItems,
    getToBuyItems,
    selectItem,
  }
});