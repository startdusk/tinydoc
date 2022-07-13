<template>
  <div class="text-center font-bold text-4xl font-sans mt-5">
    Tiny doc home page
  </div>
  <el-divider />
  <div class="m-5">
    <el-row :gutter="12">
      <el-col
        v-for="(o, index) in 60"
        :key="o"
        :span="3"
        :offset="0.5"
        :style="{ marginBottom: '10px' }"
      >
        <el-card
          shadow="hover"
          :body-style="{ padding: '0px', cursor: 'pointer' }"
        >
          <div
            class="flex flex-col items-center p-4"
            @click="handleClick(index)"
          >
            <img
              src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
              class="image"
            />
            <span>Yummy Document - {{ index }}</span>
            <div class="bottom">
              <time class="time">{{ currentDate }}</time>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

import { routing } from "../utils/routing";

const currentDate = ref(new Date().toDateString());
const router = useRouter();

const handleClick = (docId: number) => {
  router.push(routing.detail({id: docId}));
};
</script>

<style>
.time {
  font-size: 12px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.image {
  width: 180px;
  height: 200px;
  display: block;
}
</style>
