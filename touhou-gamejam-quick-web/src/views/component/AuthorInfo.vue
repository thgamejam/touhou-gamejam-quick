<template>
  <div class="overall">
    <el-space wrap>
      <el-avatar shape="square" :size="100" :src="authorObj.img" />

      <el-space direction="vertical">
        <el-link type="info" style="margin-bottom:-6px">{{ authorObj.name }}</el-link>
        <el-button type="primary" style="margin-top:36px">关注</el-button>
      </el-space>
    </el-space>
  </div>
</template>

<script>
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
        name: "舍僻阁",
        img: "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fi0.hdslb.com%2Fbfs%2Farticle%2F96cc3eaca8df7cce99cd3b5a22843206dc522ec2.jpg&refer=http%3A%2F%2Fi0.hdslb.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1637070918&t=e9d5d7a77aca295eecd2b43a9543be7a",
      },
    };
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