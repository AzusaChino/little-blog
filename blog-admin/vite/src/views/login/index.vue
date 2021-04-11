<template>
  <h1>{{ msg }}</h1>
  <button @click="realTime.count++">count is: {{ realTime.count }}</button>
  <button @click="handleClick">点击跳转其它路由</button>
  <p>Edit <code>components/HelloWorld.vue</code> to test hot module replacement.</p>
</template>

<script>
import {defineComponent, reactive} from "vue";
import {useRouter} from 'vue-router'

export default defineComponent({
  name: 'Index',
  props: {msg: {type: String, default: '欢迎来到vue3'}},
  setup(props) {
    const router = useRouter();
    let realTime = reactive({count: 0});

    function handleClick() {
      router.push({path: '/other'})
    }

    return {
      msg: props.msg,
      realTime,
      handleClick: handleClick
    }
  }
})
</script>
