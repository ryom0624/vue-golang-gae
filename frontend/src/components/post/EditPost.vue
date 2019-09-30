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
      console.log('%cPARAMATER: %c' + this.$route.params.slug, 'font-weight: bold; color: red;', '')
      if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'test') {
        axios.get(`${process.env.VUE_APP_DEVELOPMENT_BACKENDHOST}/api/v1/post/` + this.$route.params.slug).then(res => {
          console.log('%cGET edit RESPONSE DATA :%c ' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
          this.post = res.data
        })
      }
    },
    updatePost () {
      if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'test') {
        axios.post(`${process.env.VUE_APP_DEVELOPMENT_BACKENDHOST}/api/v1/post/` + this.$route.params.slug, this.post).then(res => {
          console.log('%cPOST edit RESPONSE DATA :%c ' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
          this.post = res.data
        })
      } else {
        axios.post(`https://backend-api-dot-testing-190927-golang.appspot.com/api/v1/post` + this.$route.params.slug, this.post).then(res => {
          console.log('%cPOST edit RESPONSE DATA :%c ' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
          this.post = res.data
        })
      }
      //   axios.get(`${process.env.VUE_APP_PRODUCTION_BACKENDHOST}/api/v1/post` + this.$route.params.slug, this.post).then(res => {
      //     console.log('%cPOST edit RESPONSE DATA :%c ' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
      //     this.post = res.data
      //   })
      // }
      console.log('%cpostData : %c' + JSON.stringify(this.post), 'color: red; font-weight:bold', '')
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
