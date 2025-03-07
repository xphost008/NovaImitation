<script setup lang="ts">
    import MyNormalButton from '../../../components/button/MyNormalButton.vue'
    import MyTextInput from '../../../components/input/MyTextInput.vue'
    import MyNormalLabel from '../../../components/input/MyNormalLabel.vue'
    import { ref } from 'vue'
    const chess_width = ref('')
    const chess_height = ref('')
    const chess_mine = ref('')
    function reset() {
        chess_width.value = '8'
        chess_height.value = '8'
        chess_mine.value = '10'
    }
    const real_width = ref(0)
    const real_height = ref(0)
    function start() {
        let w = parseInt(chess_width.value)
        let h = parseInt(chess_height.value)
        if(Number.isNaN(w) || Number.isNaN(h) || w < 1 || h < 1) { return }
        real_width.value = parseInt(chess_width.value) * 25
        real_height.value = parseInt(chess_height.value) * 25
    }
</script>
<template>
    <div style="display: flex; overflow-x: auto; overflow-y: auto;">
        <div id="bar">
            <MyNormalLabel for="chess-width">场地宽度</MyNormalLabel><br>
            <MyTextInput :place_holder="'[1,+∞] 区间范围'" id="chess-width" v-model="chess_width" class="input-text" /><br>
            <MyNormalLabel for="chess-height">场地高度</MyNormalLabel><br>
            <MyTextInput :place_holder="'[1,+∞] 区间范围'" id="chess-height" v-model="chess_height" class="input-text" /><br>
            <MyNormalLabel for="chess-mine">场地雷数</MyNormalLabel><br>
            <MyTextInput :place_holder="'[1,w*h-1] 区间范围'" id="chess-mine" v-model="chess_mine" class="input-text" /><br>
            <MyNormalButton class="input-button" @click="reset">默认条件</MyNormalButton>
            <MyNormalButton class="input-button" style="margin-left: 10px;" @click="start">开始游戏</MyNormalButton>
        </div>
        <div id="chess" :style="'width: ' + real_width + 'px; height: ' + real_height + 'px;'"></div>
    </div>
</template>
<style scoped>
    #bar {
        width: 200px;
        height: calc(100% - 20px);
        padding: 10px;
    }
    #chess {
        position: relative;
        background-color: red;
        flex-shrink: 0;
    }
    .input-text {
        width: 180px;
        height: 26px;
        font-size: 15px;
        transition: all 0.2s;
    }
    .input-button {
        width: 95px;
        height: 23px;
        margin-top: 5px;
    }
</style>