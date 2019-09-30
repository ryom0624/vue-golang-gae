import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Articles from '@/components/article/Articles'
import Article from '@/components/article/Article'
import IndexPosts from '@/components/post/IndexPosts'
import ShowPost from '@/components/post/ShowPost'
import NewPost from '@/components/post/NewPost'
import EditPost from '@/components/post/EditPost'

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
      path: '/article/:id',
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
