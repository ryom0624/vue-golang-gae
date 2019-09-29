<template>
  <div id="app">
    <div id="nav">
      <ul>
        <li><router-link to="/">Home</router-link> | </li>
        <li><router-link to="/articles">Article</router-link> | </li>
        <li><router-link to="/posts">Posts</router-link> | </li>
      </ul>
    </div>
    <router-view/>
  </div>
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

#nav li {
  display: inline;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
