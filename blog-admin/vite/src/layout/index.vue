<template>
  <div class="app-wrapper">
    <sidebar class="sidebar-container"/>
  </div>

</template>

<script>
import {defineComponent} from "vue";

import AppMain from './components/AppMain'
import Navbar from './components/Navbar'
import Sidebar from './components/Sidebar'

export default defineComponent({
  name: "Layout",
  components: {
    AppMain,
    Navbar,
    Sidebar
  },
  setup() {

  }
})
</script>

<style scoped>

</style>
