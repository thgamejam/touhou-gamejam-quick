<template>
  <el-header :style="'width:'+width+'px'">
    {{ $route.params.id }}
  </el-header>
  <el-container :style="'width:'+width+'px'">
    <el-main>
      <GameExhibitComponent :id="$route.params.id" width="100" />
    </el-main>
    <el-aside width="300px">
      <div class="affix-container">
        <el-affix target=".affix-container" :offset="80">
          <div class="des">
            <AuthorInfoComponent />
            <GameDownload />
          </div>
        </el-affix>
      </div>
    </el-aside>
  </el-container>
</template>

<script>
import * as GameAPI from "@/api/game/game";
import GameExhibitComponent from "@/views/component/GameExhibit.vue";
import AuthorInfoComponent from "@/views/component/AuthorInfo.vue";
import GameDownload from "@/views/component/GameDownload.vue";

export default {
  name: "GamePage",
  components: {
    GameExhibitComponent,
    AuthorInfoComponent,
    GameDownload,
  },

  data() {
    return {
      gameID: 0,
      width: 0,
      mainWidth: 0,
    };
  },
  created() {
    this.adaptive();
  },
  beforeMount() {
    this.getGame();
  },
  methods: {
    adaptive() {
      const lowestWidth = 800;
      this.width = (window.innerWidth * 0.8).toFixed(0);
      this.width = window.innerWidth < lowestWidth / 0.8 ? 800 : this.width;
      this.mainWidth = this.width - 300;
      console.log("width: " + this.width);
      console.log("mainWidth: " + this.mainWidth);
    },
    getGame() {
      GameAPI.getGame(this.$route.params.id).then((res) => {
        this.obj = res;
      });
    },
  },
};
</script>

<style>
.el-container {
  margin: 0 auto;
  /*width: 80%;*/
}
.el-header {
  background-color: #b3c0d1;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 60px;
  margin: 0 auto;
  /*width: 80%;*/
}
.el-aside {
  background-color: #d3dce6;
  color: var(--el-text-color-primary);
  text-align: center;
}
</style>
