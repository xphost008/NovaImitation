<script setup lang="ts">
    import MyNormalButton from '../../../components/button/MyNormalButton.vue'
    import MyTextInput from '../../../components/input/MyTextInput.vue'
    import MyNormalLabel from '../../../components/input/MyNormalLabel.vue'
    import { ref } from 'vue'
    const chess_width = ref('')
    const chess_height = ref('')
    const chess_mine = ref('')
    function reset() {
        chess_width.value = '9'
        chess_height.value = '9'
        chess_mine.value = '10'
    }
    /**
     * 定义格子类型
     * 其中问号可以被点击，旗子无法被点击。
     */
    interface Grids {
        // 格子横坐标
        x: number,
        // 格子纵坐标
        y: number,
        // 格子数字（-1：雷、0：周围8格没有雷、1：有一个雷（剩下以此类推）、2、3、4、5、6、7、8）
        p: -1 | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8,
        // 格子是否被标记（-1：已被点击、0：未标记、1：旗子、2：问号）
        m: -1 | 0 | 1 | 2,
        // 格子应该被显示成什么（有炸弹、旗子、问号、1~8的数字，其中0不会显示）
        s: string,
        // 是否是当前点击的雷（用于判断点击该雷的时候背景变成红色。）
        c: boolean,
    }
    //剩余旗子数
    const flag = ref(0)
    //已使用时间秒数
    const time = ref(0)
    //分数（每揭开一个格子+10分）
    const score = ref(0)
    //格子数组
    const grids = ref<Grids[][]>([[]])
    //是否开始游戏
    let is_start = false
    //临时时间
    let t = 0
    //锁定全局
    let lock = false
    //是否输赢（0：未判断、1：你赢了、2：你输了）
    const win = ref(0)
    //棋盘实际宽度
    let wi = 0
    //棋盘实际高度
    let he = 0
    //棋盘实际雷数
    let mi = 0
    //棋盘记录雷数
    let count = 0
    //全局计时器函数，用于记录秒数……
    setInterval(() => {
        if(!is_start || lock) return
        if(t > 10) {
            t = 0;
            //每秒+1
            time.value++
        }
        t++
    }, 100)
    //初始化数组
    function init_array() {
        grids.value.splice(0)
        for(let i = 0; i < he; i++) {
            grids.value.push([])
            for(let j = 0; j < wi; j++) {
                grids.value[i][j] = {
                    x: j,
                    y: i,
                    p: 0,
                    m: 0,
                    s: "",
                    c: false,
                }
            }
        }
    }
    //生成雷
    function generate_mine() {
        for(let i = 0; i < mi; i++) {
            while(true){
                let x = Math.floor(Math.random() * he)
                let y = Math.floor(Math.random() * wi)
                if(grids.value[x][y].p == -1) { continue }
                grids.value[x][y].p = -1
                break
            }
        }
    }
    //初始化数字（为雷的周围生成数字）
    function init_number() {
        for(let i = 0; i < he; i++) {
            for(let j = 0; j < wi; j++) {
                if (grids.value[i][j].p == -1) { continue }
                var foo = 0;
                if(i > 0 && j > 0) {
                    if(grids.value[i - 1][j - 1].p == -1) {
                        foo++
                    }
                }
                if(i > 0 && j < wi - 1) {
                    if(grids.value[i - 1][j + 1].p == -1) {
                        foo++
                    }
                }
                if(i < he - 1 && j > 0) {
                    if(grids.value[i + 1][j - 1].p == -1) {
                        foo++
                    }
                }
                if(i < he - 1 && j < wi - 1) {
                    if(grids.value[i + 1][j + 1].p == -1) {
                        foo++
                    }
                }
                if(i > 0) {
                    if(grids.value[i - 1][j].p == -1) {
                        foo++
                    }
                }
                if(i < he - 1) {
                    if(grids.value[i + 1][j].p == -1) {
                        foo++
                    }
                }
                if(j > 0) {
                    if(grids.value[i][j - 1].p == -1) {
                        foo++
                    }
                }
                if(j < wi - 1) {
                    if(grids.value[i][j + 1].p == -1) {
                        foo++
                    }
                }
                grids.value[i][j].p = foo as 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8
            }
        }
    }
    //开始游戏~
    function start() {
        let w = parseInt(chess_width.value)
        let h = parseInt(chess_height.value)
        let m = parseInt(chess_mine.value)
        if(Number.isNaN(w) || Number.isNaN(h) || Number.isNaN(m) || w < 1 || h < 1 || m < 1 || m > w * h - 1) { return }
        wi = w
        he = h
        mi = m
        is_start = false
        init_array()
        generate_mine()
        init_number()
        time.value = 0
        score.value = 0
        flag.value = m
        win.value = 0
        lock = false
        count = 0
    }
    //格子左键点击（揭开旗子）
    function grid_button_click(x: number, y: number) {
        if(lock) return 
        if(!is_start) is_start = true
        if(grids.value[y][x].m != -1 && grids.value[y][x].m != 1) {
            grids.value[y][x].m = -1
            let g = grids.value[y][x].p == -1 ? '💣' : grids.value[y][x].p == 0 ? "" : grids.value[y][x].p.toString()
            grids.value[y][x].s = g
            if(grids.value[y][x].p == -1) {
                lock = true
                win.value = 2
                for(let k = 0; k < he; k++) {
                    for(let l = 0; l < wi; l++) {
                        if(grids.value[k][l].p == -1) {
                            grids.value[k][l].s = '💣'
                        }
                    }
                }
                grids.value[y][x].c = true
                return
            }
            count++
            score.value += 10
            if(count == wi * he - mi) {
                lock = true
                win.value = 1
                return
            }
            if(grids.value[y][x].p != 0) return
            if(x > 0 && y > 0) {
                grid_button_click(x - 1, y - 1)
            }
            if(x > 0 && y < he - 1) {
                grid_button_click(x - 1, y + 1)
            }
            if(x < wi - 1 && y > 0) {
                grid_button_click(x + 1, y - 1)
            }
            if(x < wi - 1 && y < he - 1) {
                grid_button_click(x + 1, y + 1)
            }
            if(x > 0){
                grid_button_click(x - 1, y)
            }
            if(x < wi - 1) {
                grid_button_click(x + 1, y)
            }
            if(y > 0) {
                grid_button_click(x, y - 1)
            }
            if(y < he - 1) {
                grid_button_click(x, y + 1)
            }
        }
    }
    //格子右键点击（标旗或者揭开【周围】格子）
    function grid_button_right(x: number, y: number) {
        if(lock) return
        switch(grids.value[y][x].m) {
            case -1:
                let cc = grids.value[y][x].p
                let k = 0
                if(x > 0 && y > 0)
                    if(grids.value[y - 1][x - 1].m == 1)
                        k++
                if(x > 0 && y < he - 1)
                    if(grids.value[y + 1][x - 1].m == 1)
                        k++
                if(x < wi - 1 && y > 0)
                    if(grids.value[y - 1][x + 1].m == 1)
                        k++
                if(x < wi - 1 && y < he - 1)
                    if(grids.value[y + 1][x + 1].m == 1)
                        k++
                if(x > 0)
                    if(grids.value[y][x - 1].m == 1)
                        k++
                if(x < wi - 1)
                    if(grids.value[y][x + 1].m == 1)
                        k++
                if(y > 0)
                    if(grids.value[y - 1][x].m == 1)
                        k++
                if(y < he - 1)
                    if(grids.value[y + 1][x].m == 1)
                        k++
                console.log(cc, k)
                // 如果格子本身的数字小于或者等于周围插旗子的数字，则开启
                if(cc > k) return
                if(x > 0 && y > 0)
                    grid_button_click(x - 1, y - 1)
                if(x > 0 && y < he - 1)
                    grid_button_click(x - 1, y + 1)
                if(x < wi - 1 && y > 0)
                    grid_button_click(x + 1, y - 1)
                if(x < wi - 1 && y < he - 1)
                    grid_button_click(x + 1, y + 1)
                if(x > 0)
                    grid_button_click(x - 1, y)
                if(x < wi - 1)
                    grid_button_click(x + 1, y)
                if(y > 0)
                    grid_button_click(x, y - 1)
                if(y < he - 1)
                    grid_button_click(x, y + 1)
                break
            case 0:
                if(flag.value > 0) {
                    grids.value[y][x].m = 1
                    grids.value[y][x].s = "🚩"
                    flag.value--
                }else{
                    grids.value[y][x].m = 2
                    grids.value[y][x].s = "❔"
                }
                break
            case 1:
                grids.value[y][x].m = 2
                grids.value[y][x].s = "❔"
                flag.value++
                break
            case 2:
                grids.value[y][x].m = 0
                grids.value[y][x].s = ""
                break
        }
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
            <div id="info">
                <MyNormalLabel>分数：{{ score }}</MyNormalLabel><br>
                <MyNormalLabel>时间：{{ time }}</MyNormalLabel><br>
                <MyNormalLabel>旗子：{{ flag }}</MyNormalLabel><br>
                <MyNormalLabel style="font-size: 50px;">{{ win == 1 ? '你赢了' : win == 2 ? '你输了' : '' }}</MyNormalLabel><br>
            </div>
        </div>
        <div id="chess" :style="'width: ' + (parseInt(chess_width) * 25 || 0) + 'px; height: ' + (parseInt(chess_height) * 25 || 0) + 'px;'">
            <div class="grid-row" v-for="grid1 in grids">
                <div class="grid-column" v-for="grid in grid1">
                    <button :class="grid.m == -1 ? (grid.c ? 'grid-button-red' : 'grid-button-click') : 'grid-button'" :style="'left: ' + (grid.x * 25) + 'px; top: ' + (grid.y * 25) + 'px; color: ' + (
                        grid.p == 1 ? 'rgb(0, 128, 255)' :
                        grid.p == 2 ? 'rgb(0, 128, 0)' :
                        grid.p == 3 ? 'rgb(192, 0, 0)' :
                        grid.p == 4 ? 'rgb(0, 0, 96)' :
                        grid.p == 5 ? 'rgb(128, 0, 48)' :
                        grid.p == 6 ? 'rgb(0, 96, 96)' :
                        grid.p == 7 ? 'rgb(10, 10, 10)' :
                        grid.p == 8 ? 'rgb(100, 100, 100)' : 'black') + ';'" @click="grid_button_click(grid.x, grid.y)" @contextmenu="grid_button_right(grid.x, grid.y)">{{ grid.s }}</button>
                </div>
            </div>
        </div>
    </div>
</template>
<style scoped>
    #bar {
        width: 200px;
        height: calc(100% - 20px);
        padding: 10px;
        border-right: 2px solid gray;
    }
    #chess {
        position: relative;
        background-color: lightgray;
        flex-shrink: 0;
    }
    #info {
        margin-top: 20px;
    }
    #info > label {
        font-size: 24px;
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
    .grid-button,
    .grid-button-click,
    .grid-button-red {
        position: absolute;
        width: 25px;
        height: 25px;
        border: 1px solid black;
        background-color: lightgray;
        font-size: 16px;
        cursor: pointer;
        font-weight: bold;
    }
    .grid-button-click,
    .grid-button-red {
        background-color: gray;
        cursor: default;
    }
    .grid-button-red {
        background-color: red;
    }
    .grid-button:hover {
        background-color: silver;
    }
    .grid-button:active {
        background-color: darkgray;
    }
</style>