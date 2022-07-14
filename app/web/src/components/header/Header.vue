<template>
  <div class="flex flex-row">
    <div class="basis-1/5 flex flex-row justify-center items-center">
      <el-page-header @back="handleBackBtnClick" />
      <el-input
        v-model="searchText"
        placeholder="Please input a search key"
        clearable
      />
      <el-button @click="handleSearchBtnClick"> Search </el-button>
    </div>
    <div class="basis-4/5 flex flex-col items-center">
      <div class="font-sans text-4xl">{{ title }}</div>
      <div class="font-sans">
        发布时间: {{ releaseTime }} 版本: {{ version }}
      </div>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: "Header",
};
</script>

<script lang="ts" setup>
import { ref, toRefs } from "vue";

interface Props {
  title: string;
  version: string;
  releaseTime: string;
}

const props = defineProps<Props>();
const { title, version, releaseTime } = toRefs(props);
const searchText = ref("");

const emit = defineEmits<{
  (e: "click-search", text: string): void;
  (e: "click-back"): void;
}>();

const handleSearchBtnClick = () => {
  emit("click-search", searchText.value);
};

const handleBackBtnClick = () => {
  emit("click-back");
};
</script>

