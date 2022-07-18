<template>
  <HomeHeader title="Tiny Docs" @click-search="handleSearchBtnClick" />
  <Documents @click="handleClick" :datas="datas" />
  <el-backtop :right="100" :bottom="100" />
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

import { ElMessage } from "element-plus";
import HomeHeader from "../components/header/HomeHeader.vue";
import Documents from "../components/document/Documents.vue";
import { DocData } from "../components/document";

import { routing } from "../utils/routing";
import { time } from "../utils/time";

const currentDate = time.formatTime(new Date(), "yyyy-MM-dd");
const router = useRouter();

const handleClick = (data: DocData) => {
  router.push(routing.detail({ id: data.docId }));
};

// ==========================================================================
const datas = ref<DocData[]>([]);
datas.value = [
  {
    docId: 1,
    title: "文章标题1",
    publishTime: currentDate,
  },
];

// ===========================================================================
// Click search button
const handleSearchBtnClick = (text: string) => {
  if (!text) {
    ElMessage({
      message: "请输入搜索关键字",
      showClose: true,
    });
    return;
  }
  ElMessage({
    message: "TODO: 搜索功能待开发",
    type: "warning",
    showClose: true,
  });
};
</script>
