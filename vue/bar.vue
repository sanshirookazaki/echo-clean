<template>
  <div class="bar">
    <h1 class="header"> Bar component</h1>
    <p>This is Bar component</p>
    <foo class="header"/>
  </div>
</template>

<script>
import Foo from './foo'
export default {
  components: {
    Foo
  }
}
</script>

<style scoped>
  .bar {
      border: solid 1px red;
      margin: 4px;
      padding: 4px;
  }
  .header {
      font-size: 125%;
  }
</style>