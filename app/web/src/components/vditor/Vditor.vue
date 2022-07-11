<template>
  <div id="vditor" />
</template>
<script lang="ts">
export default {
  name: "Vditor",
};
</script>
<script setup lang="ts">
import { ref, onMounted, watch, watchEffect, toRefs } from "vue";
import Vditor from "vditor";
import $ from "jquery";
import "vditor/dist/index.css";

interface Props {
  markdown: string;
  outlinePosition: "left" | "right";
}
const props = defineProps<Props>();
const { markdown, outlinePosition } = toRefs(props);

const vditor = ref<Vditor | null>(null);

onMounted(() => {
  renderMarkdown(markdown.value, outlinePosition.value);
});

watch(
  () => markdown.value,
  (newVal, oldVal) => {
    renderMarkdown(newVal, outlinePosition.value);
  }
);

const renderMarkdown = (
  markdown: string,
  outlinePosition: "left" | "right"
) => {
  vditor.value = new Vditor("vditor", {
    height: window.innerHeight - 17,
    cache: { enable: false },
    value: "",
    toolbar: [
      {
        name: "more",
        toolbar: [
          "preview",
        ],
      },
    ],
    toolbarConfig: {
      hide: true,
    },
    outline: {
      enable: true,
      position: outlinePosition,
    },
    preview: {
      mode: "editor",
      actions: [],
      hljs: {
        enable: true,
        lineNumber: true,
      },
    },
    after: () => {
      // vditor.value is a instance of Vditor now and thus can be safely used here
      vditor.value!.setValue(markdown);
      const previewBtn = $("button[data-type='preview']");
      previewBtn.trigger("click");
    },
  });
};
</script>
<style scoped>
:global(.vditor-toolbar--hide) {
  transition: all 0.15s ease-in-out;
  height: 0px !important;
  overflow: hidden;
}
</style>
