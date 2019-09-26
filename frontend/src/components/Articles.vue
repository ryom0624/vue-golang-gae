<template>
  <div class="Articles">
    <h1>{{ msg }}</h1>
    <ul>
      <!-- articlesの配列をループで回して取得 -->
      <li v-for="article in articles" :key="article.Id" @click="pageto(article.Id)">
        <p>{{article.Title}}</p>
      </li>
    </ul>
    <h2><router-link to="/">Homeへ</router-link></h2>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Articles',
  data () {
    return {
      msg: 'This is Articles Page',
      articles: []
    }
  },
  methods: {
    pageto: function (id) {
      this.$router.push(`/articles/${id}`)
    }
  },

  created () {
    if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'test') {
      axios.get(`${process.env.VUE_APP_DEVELOPMENT_BACKENDHOST}/api/v1/articles`).then(res => {
        console.log(res.data)
        this.articles = res.data
      })
    } else if (process.env.NODE_ENV === 'production') {
      axios.get('https://backend-api-dot-testing-190927-golang.appspot.com/api/v1/articles').then(res => {
        console.log(res.data)
        this.articles = res.data
      })
    // } else if (process.env.NODE_ENV === 'production') {
    //   console.log('fuga')
    //   axios.get(process.env.VUE_APP_PRODUCTION_BACKENDHOST + '/api/v1/articles').then(res => {
    //     console.log(res.data)
    //     this.articles = res.data
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
  display: block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
