<template>
  <div class="article">
    <h1>{{article.Title}}</h1>
    <h2>{{article.Description}}</h2>
    <p><router-link to="/">Homeへ</router-link></p>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Article',
  data () {
    return {
      // createdでデータを取得したデータを入れてくれるので一旦null
      article: ''
    }
  },
  created () {
    // urlで指定された動的パラメーターから商品情報をとってくる。
    if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'test') {
      axios.get(`${process.env.VUE_APP_DEVELOPMENT_BACKENDHOST}/api/v1/articles/${this.$route.params.id}`).then(res => {
        console.log(res.data)
        this.article = res.data
      })
    } else if (process.env.NODE_ENV === 'production') {
      axios.get(`https://backend-api-dot-testing-190927-golang.appspot.com/api/v1/articles/${this.$route.params.id}`).then(res => {
        console.log(res.data)
        this.article = res.data
      })
    // } else if (process.env.NODE_ENV === 'production') {
    //   axios.get(process.env.VUE_APP_PRODUCTION_BACKENDHOST + `/api/v1/articles/${this.$route.params.id}`).then(res => {
    //     console.log(res.data)
    //     this.article = res.data
    //   })
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
