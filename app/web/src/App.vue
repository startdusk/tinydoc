<template>
  <Vditor :markdown="markdown" />
</template>

<script setup lang="ts">
import Vditor from "./components/Vditor.vue";
const markdown = `
        ## Vue Composition API + Vditor + TypeScript Minimal Example

        Topic 1

        Topic 2
        `;
</script>
<style scoped></style>
