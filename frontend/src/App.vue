<template>
<!--  <div id="app">-->
<!--    <div id="nav">-->
<!--      <ul>-->
<!--        <li><router-link to="/">Home</router-link> | </li>-->
<!--        <li><router-link to="/articles">Article</router-link> | </li>-->
<!--        <li><router-link to="/posts">Posts</router-link> | </li>-->
<!--      </ul>-->
<!--    </div>-->
<!--    <router-view/>-->
<!--  </div>-->

  <v-app>
    <v-app-bar app>
      <v-toolbar-title class="headline text-uppercase">
        <span class="font-weight-light">Vue golang Application</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <router-link to="/" class="router-link">
        <v-btn text>
          <span class="mr-2">home</span>
        </v-btn>
      </router-link>
      <router-link to="/posts" class="router-link">
        <v-btn text>
          <span class="mr-2">posts</span>
        </v-btn>
      </router-link>
      <router-link to="/HelloWorld" class="router-link">
        <v-btn text>
          <span class="mr-2">helloWord</span>
        </v-btn>
      </router-link>
    </v-app-bar>

    <v-content>
      <router-view/>
    </v-content>
  </v-app>

</template>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
  }
  #nav {
    padding: 30px;
  }

  #nav a {
    font-weight: bold;
    color: #2c3e50;
  }

  #nav ul {
    list-style: none;
  }

  .router-link {
    text-decoration: none;
  }

  .router-link span {
    color: #2c3e50;
  }

  #nav li {
    display: inline;
  }

  #nav a.router-link-exact-active {
    color: #42b983;
  }
</style>
