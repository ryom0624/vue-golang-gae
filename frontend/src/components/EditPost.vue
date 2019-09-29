<template>
  <div class="editPost">
    <div class="loading" v-if="loading">Loading...</div>

    <div class="post" v-if="post">
      <h1>{{msg}}</h1>
      title : <input v-model="post.title" type="text"><br>
      slug: <input v-model="post.slug" type="text"><br>
      description: <br>
      <textarea v-model="post.description" rows="8"></textarea><br>
      <input type="submit" @click="updatePost()">
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'EditPost',
  data () {
    return {
      msg: 'Post編集',
      loading: false,
      post: null
    }
  },
  created () {
    this.fetchData()
  },
  watch: {
    '$route': 'fetchData'
  },
  methods: {
    fetchData () {
      this.loding = false
      console.log('params: ' + this.$route.params.slug)
      if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'test') {
        axios.get(`${process.env.VUE_APP_DEVELOPMENT_BACKENDHOST}/api/v1/post/` + this.$route.params.slug).then(res => {
          console.log('GET res data:' + JSON.stringify(res.data))
          this.post = res.data
        })
      }
    },
    updatePost () {
      if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'test') {
        axios.post(`${process.env.VUE_APP_DEVELOPMENT_BACKENDHOST}/api/v1/post/` + this.$route.params.slug, this.post).then(res => {
          console.log('POST res data:' + JSON.stringify(res.data))
          this.post = res.data
        })
        console.log('postData' + JSON.stringify(this.post))
      }
      this.$router.push(`/post/${this.post.slug}`)
    }
  }
}

</script>
<style>
  .editPost input, textarea {
    border: 1px solid #ddd;
  }
</style>
