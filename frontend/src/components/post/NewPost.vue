<template>
  <div class="newPost">
    <h1>{{msg}}</h1>
    title : <input v-model="title" type="text">
    <p>title : {{title}}</p>
    slug: <input v-model="slug" type="text">
    <p>slug : {{slug}}</p>
    description: <br>
    <textarea v-model="description" rows="8"></textarea>
    <p>description : {{description}}</p>
    <input type="submit" @click="post()">
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'newPost',
  data () {
    return {
      msg: 'Post投稿',
      title: '',
      description: '',
      slug: ''
    }
  },
  methods: {
    post () {
      let postData = {
        title: this.title,
        description: this.description,
        slug: this.slug
      }
      if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'test') {
        axios.post(`${process.env.VUE_APP_DEVELOPMENT_BACKENDHOST}/api/v1/post`, postData).then(res => {
          console.log('%cPOST new RESPONSE data :%c' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
        })
      } else {
        axios.post(`https://backend-api-dot-testing-190927-golang.appspot.com/api/v1/post`, postData).then(res => {
          console.log('%cPOST new RESPONSE data :%c' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
        })
      }
      //   axios.post(`${process.env.VUE_APP_PRODUCTION_BACKENDHOST}/api/v1/post`, postData).then(res => {
      //     console.log('%cPOST new RESPONSE data :%c' + JSON.stringify(res.data, null, 4), 'color:red; font-weight:bold;', '')
      //   })
      // }
      console.log('%cpostData : %c' + JSON.stringify(this.post), 'color: red; font-weight:bold', '')
      postData = ''
      this.$router.push(`/posts`)
    }
  }
}

</script>

<style>

.newPost input, textarea {
  border: 1px solid #ddd;
}

</style>
