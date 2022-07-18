<template>
  <div v-loading.fullscreen.lock="fullscreenLoading" class="flex flex-col p-3">
    <DetailHeader
      @click-search="handleSearchBtnClick"
      @click-back="handleBackBtnClick"
      :title="title"
      :version="version"
      :releaseTime="releaseTime"
    />
    <div class="flex flex-row">
      <div class="basis-1/5">
        <Catalogue
          @click="handleCatalogueNodeClick"
          @load="handleCatalogueLoadNodeClick"
        />
      </div>
      <div v-loading.lock="vditorLoading" class="basis-4/5">
        <Vditor :markdown="markdown" outlinePosition="right" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from "vue";
import { ElMessage } from "element-plus";
import { useRoute, useRouter } from "vue-router";

import Vditor from "../components/vditor/Vditor.vue";
import Catalogue from "../components/catalogue/Catalogue.vue";
import DetailHeader from "../components/header/DetailHeader.vue";
import { Tree } from "../components/catalogue";

import type Node from "element-plus/es/components/tree/src/model/node";
import { routing } from "../utils/routing";

// ===========================================================================
// Route
const route = useRoute();
const router = useRouter();

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
const title = ref("This is a header");
const version = ref("V1.01");
const releaseTime = ref("2022-07-12 10:46:32");
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

const levelData1: Tree[] = [
  {
    id: 1,
    paraentId: 0,
    label: "Level one 1",
    fileType: 2,
    leaf: true,
  },
  {
    id: 2,
    paraentId: 0,
    label: "Level one 2",
    fileType: 2,
    children: [],
  },
];

const levelData2: Tree[] = [
  {
    id: 5,
    paraentId: 4,
    label: "Level two 2-1",
    fileType: 2,
    children: [],
  },
];

const levelData3: Tree[] = [
  {
    id: 6,
    paraentId: 5,
    label: "Level three 2-1-1",
    fileType: 1,
    leaf: true,
  },
];

const handleCatalogueLoadNodeClick = (
  node: Node,
  resolve: (data: Tree[]) => void
) => {
  if (node.level === 0) {
    return resolve(levelData1);
  }
  if (!node.data.children) {
    return resolve([]);
  }

  if (node.data.id === 2) {
    return resolve(levelData2);
  }
  if (node.data.id === 5) {
    return resolve(levelData3);
  }
};

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

// ===========================================================================
// Click back button
const handleBackBtnClick = () => {
  router.push(routing.home());
};
</script>
