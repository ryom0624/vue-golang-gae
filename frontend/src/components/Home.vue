<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <h1>Vue.js golang Application</h1>
    <p>{{msg}}</p>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data () {
    return {
      msg: 'Welcome to vue-golang application! This web site\'s data uses Google Cloud Datastore.'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
