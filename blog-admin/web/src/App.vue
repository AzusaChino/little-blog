<template>
  <div id="app">
    <Home/>
  </div>
</template>

<script lang="ts">
import {Options, Vue} from "vue-class-component";
import Home from './views/home/Home'

@Options({
  components: {
    Home,
  }
})
class App extends Vue {
}

export default App
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
