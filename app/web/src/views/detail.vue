<template>
  <div v-loading.fullscreen.lock="fullscreenLoading" class="flex flex-col p-3">
    <Header
      @click="handleSearchBtnClick"
      :title="title"
      :version="version"
      :releaseTime="releaseTime"
    />
    <div class="flex flex-row">
      <div class="basis-1/5">
        <Catalogue :data="data" @click="handleCatalogueNodeClick" />
      </div>
      <div v-loading.lock="vditorLoading" class="basis-4/5">
        <Vditor :markdown="markdown" outlinePosition="right" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, watch } from "vue";
import { ElMessage } from "element-plus";
import { useRoute } from "vue-router";

import Vditor from "../components/vditor/Vditor.vue";
import Catalogue from "../components/catalogue/Catalogue.vue";
import Header from "../components/header/Header.vue";
import { Tree } from "../components/catalogue";

// ===========================================================================
// Route
const route = useRoute();

const markdowns = [
  `
        # 这是首页
        ## Vue Composition API + Vditor + TypeScript Minimal Example

        Topic 1

        Topic 2
        `,
  `
  ### 文章1

  没有内容
  `,
  `
  ### 文章2

  暂无内容
  `,

  `
  ### 文章3

  暂无内容
  `,
  `## 文章4`,
  `## 文章5`,
  `## 文章6`,
  `## 文章7`,
  `## 文章8`,
];

const markdown = ref("");
// ===========================================================================
// On mounted
onMounted(() => {
  const idRaw = route.params.id as string;
  const docId = parseInt(idRaw);
  console.log(`docId = ${docId}`);
  markdown.value = markdowns[docId % markdowns.length];
  openFullScreen();
});

// ===========================================================================
// Loading
const fullscreenLoading = ref(false);
const openFullScreen = () => {
  fullscreenLoading.value = true;
  setTimeout(() => {
    fullscreenLoading.value = false;
  }, 1500);
};

const vditorLoading = ref(false);
const openVditorScreen = () => {
  vditorLoading.value = true;
  setTimeout(() => {
    vditorLoading.value = false;
  }, 1500);
};

let currentDocId: number = 0;
const handleCatalogueNodeClick = (node: Tree) => {
  if (node.fileType != 1) {
    return;
  }
  if (currentDocId === node.id) {
    return;
  }
  currentDocId = node.id;
  markdown.value = markdowns[node.id];
  openVditorScreen();
};

const data = ref<Tree[]>([]);
data.value = [
  {
    id: 1,
    paraentId: 0,
    label: "Level one 1",
    fileType: 2,
    children: [
      {
        id: 2,
        paraentId: 1,
        label: "Level two 1-1",
        fileType: 2,
        children: [
          {
            id: 3,
            paraentId: 2,
            fileType: 1,
            label: "Level three 1-1-1",
          },
        ],
      },
    ],
  },
  {
    id: 4,
    paraentId: 0,
    label: "Level one 2",
    fileType: 2,
    children: [
      {
        id: 5,
        paraentId: 4,
        label: "Level two 2-1",
        fileType: 2,
        children: [
          {
            id: 6,
            paraentId: 5,
            label: "Level three 2-1-1",
            fileType: 1,
          },
        ],
      },
      {
        id: 7,
        paraentId: 4,
        label: "Level two 2-2",
        fileType: 2,
        children: [
          {
            id: 8,
            paraentId: 7,
            label: "Level three 2-2-1",
            fileType: 1,
          },
        ],
      },
    ],
  },
];

const title = ref("This is a header");
const version = ref("V1.01");
const releaseTime = ref("2022-07-12 10:46:32");

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
