<template>
  <el-header>
  </el-header>
  <el-container>
    <el-main >
      <el-card class="box-card main">
        <!--头部-->
        <template #header>
          <div class="card-header">
            <span>{{ name }}</span>
          </div>
        </template>
        <!--走马灯-->
        <div class="card-img">
          <el-carousel trigger="click" height="350px">
            <el-carousel-item v-for="item in img" :key="item">
              <img class="affix-container" :src=item>
            </el-carousel-item>
          </el-carousel>
        </div>
        <div class="text-des">
          <h2>简介:</h2>
          <h3>
            {{game_des}}
          </h3>
        </div>
      </el-card>
    </el-main>
    <el-aside width="30%">
      <div class="affix-container">
        <el-affix target=".affix-container" :offset="80">
          <div class="des">
            <el-avatar :v-model="author_img" :size="100" :src=author_img></el-avatar>
            <el-divider></el-divider>
            <span>作者:{{author_name}}</span>
            <el-divider></el-divider>
            <span>类型:{{game_tag}}</span>
            <el-divider></el-divider>
            <el-button type="primary" round>Download</el-button>
          </div>
        </el-affix>
      </div>
    </el-aside>
  </el-container>
</template>

<script>
export default {
  data(){
    return{
      name:"游戏名字",
      author_name:"我是作者",
      game_tag:"rpg,stg",
      game_des:"这是一串很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长的介绍",
      author_img:"https://img2.baidu.com/it/u=3043039063,2684560819&fm=26&fmt=auto",
      img:[
        "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
      ]
    }
  },
  name: 'App',
  components: {
  },
}

</script>

<style>
.text-des{
  line-height: 25px;
  text-align: left;
  height: 30%;
  width: 100%;
}
.des{
  padding-top: 30px;
  line-height: 0;
  width: 100%;
  height: 50%;
}
.affix-container{
  height: 100%;
  width: 100%;
}
.card-img{
  height: 370px;
  width: 100%;
}
.box-card {
  width: 480px;
}
.card-header {
  height: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.main{
  width: 100%;
  height: 700px;
}
.el-container{
  margin: 0 auto;
  width: 70%;
}
.el-header{
  background-color: #b3c0d1;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 60px;
  margin: 0 auto;
  width: 70%;
}

.el-aside {
  background-color: #d3dce6;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 200px;
}

.el-main {
  background-color: #e9eef3;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 160px;
}
.el-card__header{
  height: 10%;
}
.el-card__body{
  height: 70%;
}



.el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 150px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>
