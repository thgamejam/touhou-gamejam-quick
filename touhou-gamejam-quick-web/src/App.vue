<template>
  <el-header>标题</el-header>
  <el-container>
    <el-main >
      <el-card class="box-card main">
        <!--头部-->
        <template #header>
          <div class="card-header">
            <span>名字</span>
          </div>
        </template>
        <!--走马灯-->
        <div class="card-img">
          <el-carousel trigger="click" height="350px">
            <el-carousel-item v-for="item in 1" :key="item">
              <h3 class="small">{{ item }}</h3>
            </el-carousel-item>
          </el-carousel>
        </div>
      </el-card>
    </el-main>
    <el-aside width="30%">
      <div class="affix-container">
        <el-affix target=".affix-container" :offset="80">
          <div class="des">
            <el-avatar
                :size="100"
                src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
            ></el-avatar>
            <el-divider></el-divider>
            <span>描述:</span>
            <el-divider></el-divider>
            <span>作者:</span>
            <el-divider></el-divider>
            <span>类型:</span>
            <el-divider></el-divider>
          </div>
        </el-affix>
      </div>
    </el-aside>
  </el-container>
</template>

<script>
export default {
  name: 'App',
  components: {
  },
}

</script>

<style>
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
  height: 30%;
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
  height: 1000px;
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
  height: 85%;
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
