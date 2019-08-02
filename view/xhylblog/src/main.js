// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import iView from 'iview';
import axios from 'axios'
import {post,fetch,patch,put} from './request/http'
//定义全局变量
Vue.prototype.$post=post;
Vue.prototype.$fetch=fetch;
Vue.prototype.$patch=patch;
Vue.prototype.$put=put;
import 'iview/dist/styles/iview.css'
import 'iview/dist/iview.min.js'
import VueRouter  from 'vue-router'

import AddBlog from './components/AddBlog'
import Home from './components/Home'

Vue.config.productionTip = false
Vue.use(iView)
Vue.use(VueRouter)


//配置路由
const router=new VueRouter({
  routes:[
    {path:"/",component:Home},
    {path:"/article/add",component:AddBlog},
  ],
  mode:"history"
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  components: { App },
  template: '<App/>'
})