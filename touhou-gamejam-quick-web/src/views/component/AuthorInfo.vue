<template>
  <div class="overall">
    <el-space wrap>
      <el-avatar shape="square" :size="100" :src="authorObj.img" />

      <el-space direction="vertical">
        <el-link type="info" style="margin-bottom:-6px;color: black">{{ authorObj.name }}</el-link>
        <el-button type="primary" style="margin-top:36px">关注</el-button>
      </el-space>
    </el-space>
  </div>
</template>

<script>
import * as GameApi from "@/api/game/game";
export default {
  name: "AuthorInfoComponent",
  props: {
    // 作者ID
    "id": {
      type: Number,
      required: true,
    },
  },
  data() {
    return {
      authorObj: {
        name: "",
        img: "",
      },
    };
  },
  created() {
    this.getUser()
  },
  methods:{
    getUser(){
      GameApi.getUserImg(this.$route.params.id).then((res)=>{
        this.authorObj=res;
        console.log("aaaaaaaaaaaaaa")
      })
    }
  },
};
</script>

<style>
.overall {
  text-align: center;
}

.spacing {
  line-height: 20px;
  text-align: left;
}
</style>