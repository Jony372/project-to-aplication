import axios from "axios";
import { defineStore } from "pinia";
import { ref, computed } from "vue";
// import { useRoute } from "vue-router";

interface ItemInterface {
  action: string;
  brokerage: string;
  company: string;
  rating_from: string;
  rating_to: string;
  target_from: string;
  target_to: string;
  ticker: string;
  time: string;
}

export const useItemStore = defineStore("items",() => {
  const items = ref<ItemInterface[]>([]);
  const API_URL = "http://localhost:8080/api"
  let pages = ref<string[]>([]);

  const fetchItems = async (prev?: boolean) => {
    const page = prev ? previousPage() : getLastPage(); 
    const response = await axios.get(API_URL, {
      params: {
        page: page,
      },
    });
    pages.value.push(response.data?.next_page);
    const data = await response.data?.items;
    items.value = data;
  };

  const getLastPage = (): string => {
    return pages.value[pages.value.length - 1] || "";
  }

  const previousPage = (): string => {
    if (pages.value.length <= 2) {
      pages.value = [];
      return "";
    }
    pages.value.pop();
    pages.value.pop();
    return pages.value[pages.value.length - 1];
  }

  return {
    items: computed(() => items.value),
    pages: computed(() => pages.value.length),
    fetchItems,
  }
});