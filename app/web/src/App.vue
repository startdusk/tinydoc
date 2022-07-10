<template>
  <div id="vditor" />
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import Vditor from "vditor";
import $ from "jquery";
import "vditor/dist/index.css";

const vditor = ref<Vditor | null>(null);

onMounted(async () => {
  vditor.value = new Vditor("vditor", {
    height: window.innerHeight,
    cache: { enable: false },
    value: "",
    toolbar: [
      {
        name: "more",
        toolbar: [
          "code-theme",
          "content-theme",
          "export",
          "outline",
          "preview",
        ],
      },
    ],
    toolbarConfig: {
      hide: true,
    },
    outline: {
      enable: true,
      position: "left",
    },
    preview: {
      mode: "editor",
      actions: [],
      hljs: {
        enable: true,
        lineNumber: true,
      },
    },
    after: async () => {
      // vditor.value is a instance of Vditor now and thus can be safely used here
      vditor.value!.setValue(
        `## Vue Composition API + Vditor + TypeScript Minimal Example
        
        ### Topic 1

        #### Topic 2
        `
      );
      const previewBtn = $("button[data-type='preview']");
      console.log(previewBtn);
      previewBtn.click();
    },
  });
});
</script>
<style scoped>
:global(.vditor-toolbar--hide) {
  transition: all 0.15s ease-in-out;
  height: 0px !important;
  overflow: hidden;
}
</style>
