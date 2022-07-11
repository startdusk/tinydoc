<template>
  <div :style="{ height: height, overflowY: 'auto' }">
    <el-tree :data="data" :props="defaultProps" @node-click="handleNodeClick" />
  </div>
</template>

<script lang="ts">
export default {
  name: "Catalogue",
};
</script>

<script setup lang="ts">
import { Tree } from "./index";

interface Props {
  data: Tree[];
}

const { data } = defineProps<Props>();

const height = `${window.innerHeight - 17}px`;

const emit = defineEmits<{ (e: "click", node: Tree): void }>();

const handleNodeClick = (node: Tree) => {
  emit("click", node);
};

const defaultProps = {
  id: 1,
  parentId: 0,
  children: "children",
  label: "label",
};
</script>
