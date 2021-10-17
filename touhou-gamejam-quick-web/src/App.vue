<template>
  <div class="common-layout">
    <el-container>
      <el-header>标题</el-header>
      <el-main>
        <div v-for="item in 5" :key="item" class="main-box">
            <el-card class="box-card" body-style="{padding: '0px'}" >
              <template #header>
                <!--标题-->
                <div class="card-header">
                  <h2>《这是很长很长很长的标题》</h2>
                </div>
              </template>
              <!--图片div-->
              <div class="head-img">
                <div class="box-img">
                  <img class="text-img" src="../public/test.png">
                </div>

              </div>
            </el-card>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: 'App',
  components: {
  },
}

</script>

<style>
.head-img{
  width: 100%;
  height: 50%;
}

.el-card__body{
  height: 80%;
}
.box-card {
  margin: 0 0 30px;
  width: 100%;
  height: 800px;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.el-header{
  background-color: #b3c0d1;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 60px;
  width: 70%;
  margin: 0 auto;
}
.el-main {
  background-color: #e9eef3;
  color: var(--el-text-color-primary);
  text-align: center;
  width: 70%;
  margin: 0 auto;
  height: 100%;
}
.text-img{
  width: 100%;
  height: 100%;
  border-radius: 5%;
}
.box-img{
  float: left;
  border-radius: 10px;
  width: 50%;
  height: 100%;
}

.main-box{
  height: 100%;
  width: 100%;
}
</style>
