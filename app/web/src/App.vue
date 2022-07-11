<template>
  <el-row>
    <el-col :span="4">
      <Catalogue :data="data" @click="handleClick" />
    </el-col>
    <el-col :span="20">
      <Vditor :markdown="markdown" outlinePosition="right" />
    </el-col>
  </el-row>
</template>

<script setup lang="ts">
import Vditor from "./components/vditor/Vditor.vue";
import Catalogue from "./components/catalogue/Catalogue.vue";
import { Tree } from "./components/catalogue";
import { ref } from "vue";

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
];

const markdown = ref("");
markdown.value = markdowns[0];

const handleClick = (node: Tree) => {
  markdown.value = markdowns[node.id];
};

const data = ref<Tree[]>([]);
data.value = [
  {
    id: 1,
    paraentId: 0,
    label: "Level one 1",
    children: [
      {
        id: 2,
        paraentId: 1,
        label: "Level two 1-1",
        children: [
          {
            id: 3,
            paraentId: 2,
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
    children: [
      {
        id: 5,
        paraentId: 4,
        label: "Level two 2-1",
        children: [
          {
            id: 6,
            paraentId: 5,
            label: "Level three 2-1-1",
          },
        ],
      },
      {
        id: 7,
        paraentId: 4,
        label: "Level two 2-2",
        children: [
          {
            id: 8,
            paraentId: 7,
            label: "Level three 2-2-1",
          },
        ],
      },
    ],
  },
];
</script>
<style scoped></style>
