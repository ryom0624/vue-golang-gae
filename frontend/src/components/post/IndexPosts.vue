<template>
  <div class="Posts">
    <h1>{{msg}}</h1>
    <ul>
      <li v-for="post in posts" :key="post.Slug" @click="pageto(post.slug)">
        <p>{{post.title}}</p>
      </li>
    </ul>
    <router-link to="posts/new">新規投稿</router-link>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'IndexPosts',
  data () {
    return {
      msg: 'Post 一覧',
      posts: []
    }
  },
  created () {
    if (process.env.NODE_ENV === 'development' | process.env.NODE_ENV === 'test') {
      axios.get(`${process.env.VUE_APP_DEVELOPMENT_BACKENDHOST}/api/v1/posts`).then(res => {
        console.log('%cGET index RESPONSE DATA :%c ' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
        this.posts = res.data
      })
    } else {
      axios.get(`https://backend-api-dot-testing-190927-golang.appspot.com/api/v1/posts`).then(res => {
        console.log('%cGET index RESPONSE DATA :%c ' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
        this.posts = res.data
      })
    }
    //   axios.get(`${process.env.VUE_APP_PRODUCTION_BACKENDHOST}/api/v1/posts`).then(res => {
    //     console.log('%cGET index RESPONSE DATA :%c ' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
    //     this.posts = res.data
    //   })
    // }
  },
  methods: {
    pageto: function (slug) {
      this.$router.push(`/post/${slug}`)
    }
  }
}
</script>

<style>

.Posts ul {
  list-style: none;
  padding-left: 0;
}

.Posts li {
  margin: 0 auto;
}

</style>
