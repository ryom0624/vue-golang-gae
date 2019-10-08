<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="6">
        <v-form
          ref="form"
          max-width="400"
        >
          <h1>New Post Form</h1>
          <v-text-field
           v-model="title"
           :counter="10"
           label="title"
           requird>
          </v-text-field>
          <v-textarea
            v-model="description"
            rows="8"
            clearable
            clear-icon="cancel"
            label="description"
            value="This is description of the post"
          >
          </v-textarea>
          <v-text-field
            v-model="slug"
            :counter="10"
            label="slug"
            requird>
          </v-text-field>
          <v-btn @click="post()">Submit</v-btn>
        </v-form>
      </v-col>
    </v-row>
  </v-container>
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
