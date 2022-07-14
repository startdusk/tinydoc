<template>
  <div :style="{ height: height, overflowY: 'auto' }">
    <el-tree
      :props="defaultProps"
      @node-click="handleNodeClick"
      :load="HandleLoadNodeClick"
      lazy
      highlight-current
    >
      <template #default="{ node, data }">
        <template v-if="data.fileType == 2">
          <el-icon color="#409EFC" class="no-inherit"><Folder /></el-icon>
        </template>
        <template v-else>
          <el-icon><Document /></el-icon>
        </template>
        <span class="ml-1">{{ node.label }}</span>
      </template>
    </el-tree>
  </div>
</template>

<script lang="ts">
export default {
  name: "Catalogue",
};
</script>

<script lang="ts" setup>
import type Node from "element-plus/es/components/tree/src/model/node";

import { Tree } from "./index";

const height = `${window.innerHeight - 90}px`;

const emit = defineEmits<{
  (e: "click", node: Tree): void;
  (e: "load", node: Node, resolve: (data: Tree[]) => void): void;
}>();

const handleNodeClick = (node: Tree) => {
  emit("click", node);
};

const HandleLoadNodeClick = (node: Node, resolve: (data: Tree[]) => void) => {
  if (node.data.fileType == 1) {
    return;
  }
  emit("load", node, resolve);
};

const defaultProps = {
  id: 1,
  parentId: 0,
  children: "children",
  docType: 2,
  label: "label",
  isLeaf: "leaf",
};
</script>
