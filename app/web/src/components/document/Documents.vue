<template>
  <div class="m-7">
    <div class="grid grid-cols-8 gap-8">
      <div v-for="data of datas" :key="data">
        <el-card
          :body-style="{ padding: '0px', cursor: 'pointer' }"
          shadow="hover"
          @click="handleClick(data)"
        >
          <img
            src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
            class="image"
          />
          <div style="padding: 14px">
            <span>Yummy hamburger-{{ data.title }}</span>
            <div class="bottom">
              <time class="text-sm text-neutral-400"
                >Published: {{ data.publishTime }}</time
              >
            </div>
          </div>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: "Documents",
};
</script>

<script lang="ts" setup>
import { ref, toRefs } from "vue";
import { DocData } from "./index";

interface Props {
  datas: DocData[];
}

const props = defineProps<Props>();
const { datas } = toRefs(props);

const emit = defineEmits<{
  (e: "click", data: DocData): void;
}>();

const handleClick = (data: DocData) => {
  emit("click", data);
};
</script>


