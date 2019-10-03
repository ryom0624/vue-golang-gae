<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="6">
        <div v-if="post">
          <v-form
            ref="form"
            max-width="400"
          >
            <h1>Edit Post Form</h1>
            <v-text-field
              v-model="post.title"
              :counter="25"
              label="title"
              requird>
            </v-text-field>
            <v-textarea
              v-model="post.description"
              rows="8"
              clearable
              clear-icon="cancel"
              label="description"
              value="This is description of the post"
            >
            </v-textarea>
            <v-text-field
              v-model="post.slug"
              :counter="100"
              label="slug"
              requird>
            </v-text-field>
            <v-btn @click="updatePost()">Submit</v-btn>
          </v-form>
        </div>
      </v-col>
    </v-row>
  </v-container>
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
      } else {
        console.log('%cNODE_ENV: %c' + process.env.NODE_ENV, 'font-weight: bold; color: red;', '')
        axios.get(`https://backend-api-dot-testing-190927-golang.appspot.com/api/v1/post/` + this.$route.params.slug).then(res => {
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
        axios.post(`https://backend-api-dot-testing-190927-golang.appspot.com/api/v1/post/` + this.$route.params.slug, this.post).then(res => {
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
