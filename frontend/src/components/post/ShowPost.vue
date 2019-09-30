<template>
  <div class="Post">

    <h1>{{msg}}</h1>
    <h2>{{post.title}}</h2>
    <p>{{post.description}}</p>
    <p class="edit" @click="pageto(post.slug)">編集する</p>

  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'shPost',
  data () {
    return {
      msg: 'Post 詳細',
      post: ''
    }
  },
  created () {
    if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'test') {
      axios.get(`${process.env.VUE_APP_DEVELOPMENT_BACKENDHOST}/api/v1/post/${this.$route.params.slug}`).then(res => {
        console.log('%cGET show RESPONSE DATA :%c ' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
        this.post = res.data
      })
    } else {
      axios.get(`https://backend-api-dot-testing-190927-golang.appspot.com/api/v1/post/${this.$route.params.slug}`).then(res => {
        console.log('%cGET show RESPONSE DATA :%c ' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
        this.post = res.data
      })
    }
    //   axios.get(`${process.env.VUE_APP_PRODUCTION_BACKENDHOST}/api/v1/posts`).then(res => {
    //     console.log('GET show: ' + JSON.stringify(res.data, null, 4))
    //     this.post = res.data
    //   })
    // }
  },
  methods: {
    pageto: function (slug) {
      this.$router.push(`${slug}/edit`)
    }
  }
}

</script>

<style>
  .edit {
    text-decoration: underline;
    cursor: pointer;
  }
</style>
