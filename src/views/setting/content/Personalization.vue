<script setup lang="ts">
    import MySelectCard from '../../../components/card/MySelectCard.vue'
    import MyNormalButton from '../../../components/button/MyNormalButton.vue'
    import { dark_mode } from '../../../logic/changeBody'
    function changeDarkMode() {
        dark_mode.value = !dark_mode.value
    }
</script>
<template>
    <div style="display: flex; flex-direction: column;">
        <MySelectCard :maxHeight="60" :isExpand="true" :title="'暗色模式'" :isLast="true">
            <div style="text-align: center; margin: 10px;">
                <MyNormalButton id="dark-mode-button" @click="changeDarkMode">
                    切换
                </MyNormalButton>
            </div>
        </MySelectCard>
    </div>
</template>
<style scoped>
    #dark-mode-button {
        width: 200px;
        height: 40px;
        border: 1px solid blue;
    }
</style>