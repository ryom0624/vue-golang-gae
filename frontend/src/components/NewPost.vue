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
      if (process.env.NODE_ENV === 'development' || process.env.NODE_ENV === 'test') {
        let postData = {
          title: this.title,
          description: this.description,
          slug: this.slug
        }
        axios.post(`${process.env.VUE_APP_DEVELOPMENT_BACKENDHOST}/api/v1/post`, postData).then(res => {
          console.log(res.data)
        })
        console.log(postData)
        postData = ''
      }
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
