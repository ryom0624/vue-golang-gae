import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Articles from '@/components/Articles'
import Article from '@/components/Article'
import IndexPosts from '@/components/IndexPosts'
import ShowPost from '@/components/ShowPost'
import NewPost from '@/components/NewPost'
import EditPost from '@/components/EditPost'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/articles',
      name: 'Articles',
      component: Articles
    },
    {
      path: '/articles/:id',
      name: 'Article',
      component: Article
    },
    {
      path: '/posts',
      name: 'Posts',
      component: IndexPosts
    },
    {
      path: '/post/:slug',
      name: 'Post',
      component: ShowPost
    },
    {
      path: '/posts/new',
      name: 'newPost',
      component: NewPost
    },
    {
      path: '/post/:slug/edit',
      name: 'editPost',
      component: EditPost
    }
  ]
})
