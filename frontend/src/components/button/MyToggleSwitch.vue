<script setup lang="ts">
interface CheckButtonProps {
  isChecked?: boolean
}

const isCheckedProps = withDefaults(defineProps<CheckButtonProps>(), {isChecked: false})
import {ref, watch} from 'vue'

const check_value = ref(isCheckedProps.isChecked)
watch(() => isCheckedProps.isChecked, v => check_value.value = v)
const back = ref("lightgray")
const left = ref("4px")
watch(check_value, v => {
  back.value = v ? "skyblue" : "lightgray"
  left.value = v ? "28px" : "4px"
}, {immediate: true})
</script>
<template>
  <div class="switch">
    <button class="toggle-span"></button>
  </div>
</template>
<style scoped>
.switch {
  position: relative;
  width: 48px;
  height: 24px;
  background-color: v-bind(back);
  border-radius: 34px;
  cursor: pointer;
  transition: 0.2s;
  box-shadow: 2px 2px 5px gray;
}

.toggle-span {
  border-radius: 50%;
  cursor: pointer;
  width: 16px;
  height: 16px;
  margin: 4px;
  margin-left: v-bind(left);
  border: 0;
  background-color: #fff;
  transition: 0.2s;
  box-shadow: 2px 2px 5px gray;
}
</style>