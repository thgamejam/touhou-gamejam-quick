<template>
  <el-header>
    {{ $route.params.id }}
  </el-header>
  <el-container>
    <el-main>
      <GameExhibitComponent :id="$route.params.id" width="600px" />
    </el-main>
    <el-aside width="30%">
      <div class="affix-container">
        <el-affix target=".affix-container" :offset="80">
          <div class="des">
            <AuthorInfoComponent />
            
            <el-card class="aside-card">
              <el-divider
                ><i class="el-icon-star-on"></i>协议<i
                  class="el-icon-star-on"
                ></i
              ></el-divider>
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
              <el-button class="button" type="primary" round>
                <el-link
                  type="warning"
                  :underline="false"
                  :href="obj.downloadId"
                  target="_blank"
                >
                  DownLoad
                </el-link>
              </el-button>
            </el-card>
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

export default {
  name: "GamePage",
  components: {
    GameExhibitComponent,
    AuthorInfoComponent,
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
    };
  },
  created() {},
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
.el-icon-close {
  color: red;
}
.el-icon-check {
  color: green;
}

.main {
  width: 100%;
  height: 700px;
}
.el-container {
  margin: 0 auto;
  width: 70%;
}
.el-header {
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
}

</style>
