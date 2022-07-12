<template>
  <div :style="{ height: height, overflowY: 'auto' }">
    <el-tree :data="data" :props="defaultProps" @node-click="handleNodeClick">
      <template #default="{ node, data }">
        <template v-if="data.fileType == 2">
          <el-icon color="#409EFC" class="no-inherit"><Folder /></el-icon>
        </template>
        <template v-else>
          <el-icon><Document /></el-icon>
        </template>
        <span class="node-label">{{ node.label }}</span>
      </template>
    </el-tree>
  </div>
</template>

<script lang="ts">
export default {
  name: "Catalogue",
};
</script>

<script setup lang="ts">
import { toRefs } from "vue";
import { Tree } from "./index";

interface Props {
  data: Tree[];
}

const props = defineProps<Props>();
const { data } = toRefs(props);

const height = `${window.innerHeight - 90}px`;

const emit = defineEmits<{ (e: "click", node: Tree): void }>();

const handleNodeClick = (node: Tree) => {
  emit("click", node);
};

const defaultProps = {
  id: 1,
  parentId: 0,
  children: "children",
  docType: 2,
  label: "label",
};
</script>

<style scoped>
.node-label {
  margin-left: 5px;
}
</style>
