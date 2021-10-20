<template>
  <el-header>
    {{$route.params.id}}
  </el-header>
  <el-container>
    <el-main >
      <el-card class="box-card main">
        <!--头部-->
        <template #header>
          <div class="card-header">
            <span>{{ obj.name }}</span>
          </div>
        </template>
        <!--走马灯-->
        <div class="card-img">
          <el-carousel trigger="click" height="350px">
            <el-carousel-item v-for="item in obj.imgs" :key="item">
              <img class="affix-container" :src=item>
            </el-carousel-item>
          </el-carousel>
        </div>
        <div class="text-des">
          <h2>简介:</h2>
          <h3>
            {{obj.description}}
          </h3>
        </div>
      </el-card>
    </el-main>
    <el-aside width="30%">
      <div class="affix-container">
        <el-affix target=".affix-container" :offset="80">
          <div class="des">
            <el-avatar :v-model="authorObj.img" :size="100" :src=authorObj.img></el-avatar>
            <el-divider></el-divider>
            <span>作者:{{authorObj.name}}</span>
            <el-divider></el-divider>
            <span>类型:{{obj.tags}}</span>
            <el-divider></el-divider>
            <el-card class="aside-card">
              <el-divider><i class="el-icon-star-on"></i>协议<i class="el-icon-star-on"></i></el-divider>
              <el-row>
                <el-col :span="20"><div>允许商用</div></el-col>
                <el-col :span="4"><div><i class="el-icon-close"></i></div></el-col>
              </el-row>
              <el-row>
                <el-col :span="20"><div>允许再创作</div></el-col>
                <el-col :span="4"><div><i class="el-icon-close"></i></div></el-col>
              </el-row>
              <el-row>
                <el-col :span="20"><div>允许转载</div></el-col>
                <el-col :span="4"><div><i class="el-icon-close"></i></div></el-col>
              </el-row>
              <el-row>
                <el-col :span="20"><div>允许下载</div></el-col>
                <el-col :span="4"><div><i class="el-icon-check"></i></div></el-col>
              </el-row>
              <hr>
              <el-button class="button" type="primary" round>
                <el-link type="warning" :underline="false" :href=obj.downloadId target="_blank">
                  DownLoad
                </el-link>
              </el-button>
            </el-card>
          </div>
        </el-affix>
      </div>
    </el-aside>
  </el-container>
</template>

<script>
import * as GameAPI from "@/api/game/game";

export default {
  data(){
    return{
      obj:{
        name:"游戏",
        tags:"rpg,stg",
        downloadId:"https://www.baidu.com",
        authorId:"",
        description:"这是一段很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长的简介",
        imgs:[
          // "https://img1.baidu.com/it/u=1939001640,3189615103&fm=26&fmt=auto",
          // "https://gimg2.baidu.com/image_search/src=http%3A%2F%2F5b0988e595225.cdn.sohucs.com%2Fimages%2F20170822%2Ff908dec9bd9146e08f5396212d0a335b.jpeg&refer=http%3A%2F%2F5b0988e595225.cdn.sohucs.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1637070918&t=ec55041f0d5f009cc62b8b887f5ebe17",
        ],
      },
      authorObj:{
        name:"舍僻阁",
        img:"https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fi0.hdslb.com%2Fbfs%2Farticle%2F96cc3eaca8df7cce99cd3b5a22843206dc522ec2.jpg&refer=http%3A%2F%2Fi0.hdslb.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1637070918&t=e9d5d7a77aca295eecd2b43a9543be7a",
      },
    }
  },
  methods:{
    getGame(){
      GameAPI.getGame(this.$route.params.id).then((res) => {
        this.obj = res;
      });
    },
  },
  beforeMount() {
    this.getGame();
  },
  name: 'App',
  components: {
  },
}

</script>

<style>
.el-icon-close{
  color: red;
}
.el-icon-check{
  color: green;
}
.aside-card{
  text-align: left;
  width: 80%;
  margin: 0 10%;
}
.button{
  width: 100%;
  height: 35px;
}

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
