<template>
  <div class="flex flex-row">
    <div class="basis-1/5 flex flex-row justify-center items-center">
      <input
        class="border-solid border-2 bg-slate-50 caret-pink-500"
        type="text"
        v-model="searchText"
      />
      <button
        @click="handleClick"
        class="bg-cyan-500 hover:bg-cyan-600 rounded-none font-sans"
      >
        Search
      </button>
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

<script setup lang="ts">
import { ref, toRefs } from "vue";

interface Props {
  title: string;
  version: string;
  releaseTime: string;
}

const props = defineProps<Props>();
const { title, version, releaseTime } = toRefs(props);

const emit = defineEmits<{ (e: "click", text: string): void }>();
const searchText = ref("");
const handleClick = () => {
  emit("click", searchText.value);
};
</script>

