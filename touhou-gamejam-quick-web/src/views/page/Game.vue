<template>
  <el-header style=":width= mainWidth ">
    {{ $route.params.id }}
  </el-header>
  <el-container>
    <el-main>
      <GameExhibitComponent :id="$route.params.id" width="600px" />
    </el-main>
    <el-aside width="300px">
      <div class="affix-container">
        <el-affix target=".affix-container" :offset="80">
          <div class="des">
            <AuthorInfoComponent />
            <GameDownload/>

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
      obj: {
        name: "游戏",
        tags: "rpg,stg",
        downloadId: "https://www.baidu.com",
        authorId: "",
        description:
          "这是一段很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长很长的简介",
        imgs: [],
      },
      authorObj: {
        name: "舍僻阁",
        img: "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fi0.hdslb.com%2Fbfs%2Farticle%2F96cc3eaca8df7cce99cd3b5a22843206dc522ec2.jpg&refer=http%3A%2F%2Fi0.hdslb.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1637070918&t=e9d5d7a77aca295eecd2b43a9543be7a",
      },
      mainWidth:"800px",
    };
  },
  created() {
    this.mainWidth=window.innerWidth<800?800:window.innerWidth+"px";
    console.log(this.mainWidth);
  },
  beforeMount() {
    this.getGame();
  },
  methods: {
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
  width: 1000px;
}
.el-header {
  background-color: #b3c0d1;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 60px;
  margin: 0 auto;
  width: 1000px;
}
.el-aside {
  background-color: #d3dce6;
  color: var(--el-text-color-primary);
  text-align: center;
}

</style>
