<template>
  <el-card class="aside-card">
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
    
    <el-button type="primary" style="width: 100%">Download</el-button>

  </el-card>
</template>

<script>
export default {};
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
  /* margin: 0 10%; */
}
</style>