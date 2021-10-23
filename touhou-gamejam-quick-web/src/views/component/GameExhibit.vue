<template>
  <!-- 标题，游戏名字 -->
  <h2 align="left">{{ obj.name }}</h2>
  <!--走马灯-->
  <el-carousel
     v-if="update"
    :interval="3000"
    arrow="always"
    trigger="click"
    :height="carouselHeight + 'px'"
  >
    <el-carousel-item v-for="item in obj.imgs" :key="item">
      <el-image class="affix-container" :src="item">
        <template #error>
          <div class="image-slot">
            <i class="el-icon-picture-outline"></i>
          </div>
        </template>
      </el-image>
    </el-carousel-item>
  </el-carousel>

  <p />
  <div class="tag-text">
    <!-- 一条竖线 -->
    <table class="vertical-line tag-box">
      <tr>
        <td valign="top"></td>
      </tr>
    </table>
    <font size="2"></font>
    <div style="margin: 0 18px">游戏类型</div>
    <div class="tag-box" v-for="tag in obj.tags" :key="tag">
      <el-tag style="margin: 10px">{{ tag }}</el-tag>
    </div>
  </div>

  <div style="float: none;white-space: pre-wrap;">
    <h3>
      {{ obj.description }}
    </h3>
  </div>
</template>

<script>
import * as GameAPI from "@/api/game/game";

export default {
  name: "GameExhibitComponent",
  props: {
    // 游戏ID
    id: {
      type: Number,
      required: true,
    },
    // 组件宽度
    width: String,
  },
  data() {
    return {
      obj: {
        name: "",
        description: "",
        tags: [],
        imgs: [],
      },
      carouselHeight: 100, // 走马灯高度
      update: true,
    };
  },
  created() {
    this.carouselHeight = ((this.getWidth() * 9) / 16).toFixed(0);
    console.log(this.getWidth() + "*" + this.carouselHeight);
  },
  beforeMount() {
    this.getGame(this.getId());
  },
  methods: {
    //强制刷新
    reload() {
      this.update = false
      this.$nextTick(() => {
        this.update = true
      })
    },
    // 获取游戏ID
    getId() {
      return this.$props.id;
    },
    // 获取整体组件宽度
    getWidth() {
      if (this.$props.width == null) {
        console.log(this.$props.width);
        return 1440;
      }
      console.log(this.$props.width);
      return this.$props.width.substr(0, this.$props.width.length - 2);
    },
    // 获取游戏数据
    getGame(id) {
      if (id == null) {
        id = 2;
      }
      GameAPI.getGame(id).then((res) => {
        this.obj = res;
      });
    },
  },
};
</script>

<style>
.tag-text {
  width: 100%;
  height: 100px;
}
.tag-box {
  float: left;
}
.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
  text-align: center;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}

.div-left {
  width: 300px;
  height: 120px;
  border: 1px solid #000;
  float: left;
}

.affix-container {
  height: 100%;
  width: 100%;
}

.vertical-line {
  height: 80px;
  border-color: #cbd0d8;
  border-left-style: solid;
  border-width: 1px;
}

.text-des {
  line-height: 25px;
  text-align: left;
  height: 30%;
  width: 100%;
}
</style>