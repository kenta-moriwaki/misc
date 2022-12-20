<script setup lang="ts">
import { ref, watchEffect } from "vue";

const queryText = ref("");

const stateA = ref(true);

window.addEventListener("updateSearchState", (event) => {
  queryText.value = (event as CustomEvent).detail.queryText;
  stateA.value = (event as CustomEvent).detail.stateA;
});

// 待つ or リトライがいるかも
window.dispatchEvent(new CustomEvent("requestSearchState"));

watchEffect(() => {
  const detail = {
    queryText: queryText.value,
    stateA: stateA.value,
  };
  const event = new CustomEvent("setSearchState", { detail });
  window.dispatchEvent(event);
});
</script>

<template>
  <div class="wrapper">
    <p>query: {{ queryText }}</p>
    <p>
      stateA: {{ stateA }}
      <input class="search-checkbox" type="checkbox" v-model="stateA" />
    </p>
  </div>
</template>

<style scoped></style>
