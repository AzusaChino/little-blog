<template>
  <div>
    <el-container>
      <el-header>Header</el-header>
      <el-container>
        <el-aside width="200px">Aside</el-aside>
        <el-container>
          <el-main>
            <router-view/>
          </el-main>
          <el-footer>Footer</el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>

<script lang="ts">
import {Options, Vue} from 'vue-class-component';

// register imported Components
@Options({
  components: {},
})
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

</style>
