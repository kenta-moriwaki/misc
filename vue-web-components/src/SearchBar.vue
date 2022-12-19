<script setup lang="ts">
import { ref } from "vue";

import IconSupport from "./components/icons/IconSupport.vue";

const props = defineProps({
  text: {
    type: String,
    default: "text",
  },
});

const queryText = ref(props.text);
</script>

<template>
  <div class="wrapper">
    <div class="search-bar">
      <button class="search-button">
        <IconSupport />
      </button>
      <input class="search-input" type="search" :value="queryText" />
    </div>
  </div>
</template>

<style scoped>
.wrapper {
  padding: 20px;
}
.search-bar {
  width: 100%;
  height: 100%;
  display: flex;
  height: 32px;
  border-radius: 25px;
  background: #ffffff;
  box-shadow: 10px 10px 20px #7a7a7a, -10px -10px 20px #ffffff;
}

.search-button {
  width: 64px;
  flex-shrink: 0;
  border-radius: 25px 0 0 25px;
  border-width: 0;
}

.search-input {
  width: 100%;
  border-radius: 0 25px 25px 0;
  border-width: 0;
}
</style>
