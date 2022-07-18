<template>
  <div class="text-center items-center mt-5">
    <h1 class="text-4xl font-bold font-sans">{{ title }}</h1>
    <div class="flex flex-row mt-8 justify-center">
      <form action="#" class="flex flex-row w-96">
        <el-input
          v-model="searchText"
          placeholder="Please input a search key"
          clearable
        />
        <el-button @click="handleSearchBtnClick"> Search </el-button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: "HomeHeader",
};
</script>

<script lang="ts" setup>
import { ref } from "vue";

interface Props {
  title: string;
}

const { title } = defineProps<Props>();

const searchText = ref("");

const emit = defineEmits<{
  (e: "click-search", text: string): void;
}>();

const handleSearchBtnClick = () => {
  emit("click-search", searchText.value);
};
</script>


