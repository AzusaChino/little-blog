<template>
  <div>
    <HelloWorld :msg="msg"/>
    <Demo/>
  </div>
</template>

<script lang="ts">
import HelloWorld from "../../components/HelloWorld.vue";
import Demo from '../login/index.vue'

import {defineComponent} from "vue";

export default defineComponent({
  name: "Home",
  components: {HelloWorld, Demo},
  setup() {
    return {
      msg: "hello World",
    };
  },
});
</script>
