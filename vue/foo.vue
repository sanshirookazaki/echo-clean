<template>
  <div class="foo">
    <h1 class="header"> Foo component</h1>
    <p>This is Foo component</p>
  </div>
</template>

<style>
  p {text-decoration: underline;}
</style>

<style scoped>
  .foo {
      border: solid 1px green;
      margin: 4px;
      padding: 4px;
  }
  .header {
      font-size: 150%;
  }
</style>