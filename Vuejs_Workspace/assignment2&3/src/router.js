import Vue from 'vue';
import Router from 'vue-router';
import HomeComp from './components/HomeComp.vue'
import AboutComp from './components/AboutComp.vue'
import ContactComp from './components/ContactComp.vue'
import SignUp from './components/SignUp.vue'
Vue.use(Router);

  
export default new Router({
      routes:[
        {path:"/", component:HomeComp},
        {path:"/about",component:AboutComp},
        {path:"/contact",component:ContactComp},
        {path:"/signup",component:SignUp}
      ]
})