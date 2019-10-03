<template>
  <v-container class="Posts">
    <h1 class="headline">
      {{msg}}
    </h1>
    <v-row>
      <v-col
        v-for="post in posts"
        :key="post.slug"
        cols="12"
        sm="4"
      >
        <v-card
          :elevation="15"
          class="mx-auto"
          max-width="344"
          outlined
          dark
        >
          <v-list-item three-line>
            <v-list-item-content>
              <div class="overline mb-4">post</div>
              <v-list-item-title class="headline mb-1">{{post.title}}</v-list-item-title>
              <v-list-item-content class="text-left">{{post.description}}</v-list-item-content>
            </v-list-item-content>

            <v-list-item-avatar
              tile
              size="80"
              color="grey"
            ></v-list-item-avatar>
          </v-list-item>

          <v-card-actions>
            <v-btn text @click="pageto(post.slug)">Link</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>

    <v-card-actions>
      <v-btn text>
        <router-link to="posts/new">New Post</router-link>
      </v-btn>
    </v-card-actions>

  </v-container>
</template>

<script>
import axios from 'axios'

export default {
  name: 'IndexPosts',
  data () {
    return {
      msg: 'Posts',
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
