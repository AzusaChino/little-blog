<template>
  <div>
    <el-tag>
      Home
    </el-tag>
  </div>
</template>

<script lang="ts">
import {Vue} from 'vue-class-component';


export default class Home extends Vue {
  data() {
    return {
      message: "Hello World",
      router: "The router"
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
