<template>
  <div class="body-box">
    <div class="header" :style="'height'+imgHeight+'px'">
      <img class="gamejam-img" :style="'height'+imgHeight+'px'" src="http://thjam.cc:9000/public-download/head.jpg">
    </div>
    <div v-for="item in gamelist" :key="item" class="header"  :style="'height'+imgHeight+'px'">
      <a :href="'/game/'+item.id">
        <img class="gamejam-img" :src="item.img" :style="'height'+imgHeight+'px'">
      </a>
    </div>
  </div>
</template>

<script>
import * as GameApi from "@/api/game/game.js"
export default {
  data(){
    return{
      gamelist:[],
      imgHeight:100,
    };
  },
  created() {
    this.setHeight();
    this.getGameList();
    console.log(this.gamelist);
  },
  methods:{
    setHeight(){
      this.imgHeight=(window.innerWidth*0.6)/16*9;
    },
    getGameList(){
      GameApi.getGameList(1).then((res)=>{
        this.gamelist=res.list;
      })
    }
  },
}
</script>

<style>
.gamejam-img{
  border-radius: 5%;
  width: 100%;
}
.body-box{
  margin: 0 20%;
  height: 1000px;
  width: 60%;
}
.header{
  width: 100%;
}
.gamejam-main {
  background-color: #e9eef3;
  color: var(--el-text-color-primary);
}
</style>
