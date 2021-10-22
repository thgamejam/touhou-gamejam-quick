<template>
  <el-card class="aside-card" style="width: 80%;margin: 10%">
    <div class="card-header">
      <el-divider>
        <i class="el-icon-star-on"></i>协议<i class="el-icon-star-on" />
      </el-divider>
    </div>
    <el-row>
      <el-col :span="20"><div>允许商用</div></el-col>
      <el-col :span="4"
        ><div><i class="el-icon-close"></i></div
      ></el-col>
    </el-row>
    <el-row>
      <el-col :span="20"><div>允许再创作</div></el-col>
      <el-col :span="4"
        ><div><i class="el-icon-close"></i></div
      ></el-col>
    </el-row>
    <el-row>
      <el-col :span="20"><div>允许转载</div></el-col>
      <el-col :span="4"
        ><div><i class="el-icon-close"></i></div
      ></el-col>
    </el-row>
    <el-row>
      <el-col :span="20"><div>允许下载</div></el-col>
      <el-col :span="4"
        ><div><i class="el-icon-check"></i></div
      ></el-col>
    </el-row>

    <hr />
    
    <el-button
        @click="getGameDownloadURL"
        style="width: 100%"
        :type="type"
        :loading="loading"
    >
      Download
    </el-button>

  </el-card>
</template>

<script>
import * as GameApi from "@/api/game/game";
export default {
  data(){
    return{
      downloadURL:"",
      loading:false,
      type:"primary",
    };
  },
  methods:{
    getGameDownloadURL(){
      //修改按钮样式
      this.loading=true;
      this.type="danger";
      //延时请求链接并重置按钮样式
      setTimeout(()=>{
        this.loading=false;
        this.type="primary";
        GameApi.getGameDownloadURL(this.$route.params.id).then((res)=>{
          this.downloadURL=res.url;
          window.open(this.downloadURL)
        })
      },3000)
    },
  },
};
</script>

<style>
.el-icon-close {
  color: red;
}
.el-icon-check {
  color: green;
}

.aside-card {
  text-align: left;
  width: 100%;
}
</style>