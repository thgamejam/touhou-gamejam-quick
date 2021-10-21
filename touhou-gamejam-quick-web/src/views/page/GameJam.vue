<template>
  <div class="body-box">
    <div class="header" :style="'height'+imgHeight+'px'">
      <img class="gamejam-img" :style="'height'+imgHeight+'px'" src="http://test.strluck.com:9000/public-download/logo.jpg">
    </div>
    <div v-for="item in gamelist" :key="item" class="header"  :style="'height'+imgHeight+'px'">
      <a :href="'/game/'+item.id">
        <img class="gamejam-img" :src="item.img" :style="'height'+imgHeight+'px'">
      </a>
    </div>
  </div>
</template>

<script>
export default {
  data(){
    return{
      gamelist:[
        {
          id:1,
          img:"https://img1.baidu.com/it/u=1514254201,1610764859&fm=26&fmt=auto",
          name:"aaa",
        },
        {
          id:2,
          img:"https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fpic18.nipic.com%2F20111216%2F9047182_205018652177_2.jpg&refer=http%3A%2F%2Fpic18.nipic.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1637387402&t=4a63f0b795e435b8e24294b63bf4562c",
          name:"bbb",
        },
      ],
      imgHeight:100,
    };
  },
  created() {
    this.setHeight();
  },
  methods:{
    setHeight(){
      this.imgHeight=(window.innerWidth*0.6)/16*9;
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
