<template>
  <div class="common-layout">
    <el-container>
      <el-header>标题</el-header>
      <el-main>
        <div v-for="item in 5" :key="item" class="main-box">
          <el-row>
            <el-col :span="8">
              <div class="grid-content bg-purple">
                <div class="des-img">
                  <img class="box-img" src="../public/testimg.png">
                </div>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="grid-content bg-purple-light">

                这是一段文字这是一段文字
              </div>
            </el-col>
            <el-col :span="4">
              <div class="grid-content bg-purple">
                3
              </div>
            </el-col>
          </el-row>
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
.el-header{
  background-color: #b3c0d1;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 60px;
  width: 80%;
  margin: 0 auto;
}
.el-main {
  background-color: #e9eef3;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 160px;
  width: 80%;
  margin: 0 auto;
}

.box-img{
  border-radius: 10px;
  width: 100%;
  height: 100%;
}
.des-img{
  width: 100%;
  height: 100%;
}
.bg-purple {
  background: #d3dce6;
}
.bg-purple-light {
  background: #e5e9f2;
}
.grid-content {
  text-align: justify;
  border-radius: 4px;
  height: 160px;
}
.main-box{
  height: 100%;
  width: 100%;
  margin: 10px;
}
</style>
