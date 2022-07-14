<template>
  <div id="vditor" />
</template>
<script lang="ts">
export default {
  name: "Vditor",
};
</script>
<script lang="ts" setup>
import { ref, onMounted, watch } from "vue";
import Vditor from "vditor";
import $ from "jquery";
import { VditorConfig } from "./index";
import "vditor/dist/index.css";

interface Props {
  markdown: string;
  outlinePosition: "left" | "right";
}

const props = defineProps<Props>();

onMounted(() => {
  renderMarkdown({
    markdown: props.markdown,
    outlinePosition: props.outlinePosition,
  });
});

watch(
  () => props.markdown,
  (newMarkdown, _oldMarkdown) => {
    renderMarkdown({
      markdown: newMarkdown,
      outlinePosition: props.outlinePosition,
    });
  }
);

const renderMarkdown = (config: VditorConfig) => {
  const vditor = new Vditor("vditor", {
    height: window.innerHeight - 90,
    cache: { enable: false },
    value: "",
    toolbar: [
      {
        name: "more",
        toolbar: ["preview"],
      },
    ],
    toolbarConfig: {
      hide: true,
    },
    outline: {
      enable: true,
      position: config.outlinePosition,
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
      vditor.setValue(config.markdown);
      const previewBtn = $("button[data-type='preview']");
      previewBtn.trigger("click");
    },
  });
};
</script>
<style scoped>
:deep(.vditor-toolbar--hide) {
  transition: all 0.15s ease-in-out;
  height: 0px !important;
  overflow: hidden;
}
</style>
