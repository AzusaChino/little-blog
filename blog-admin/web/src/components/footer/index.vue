<template>
  <div>
    {{message}}
  </div>
</template>

<script lang="ts">
import {Options, Vue} from 'vue-class-component';

// register imported Components
@Options({
  components: {},
})

class Home extends Vue {
  data() {
    return {
      message: "Hello World",
      router: "The router"
    }
  }
}

export default Home
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
