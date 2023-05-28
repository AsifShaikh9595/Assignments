<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <StudentComp :msg="messg" :myarr="student"></StudentComp>
  </div>
</template>

<script>
//import HelloWorld from './components/HelloWorld.vue'
 import StudentComp from './components/StudentComponent.vue'
export default {
  name: 'App',
  components: {
   // HelloWorld
    StudentComp
  },
  data(){
    return{
        messg:"vue js a is framework",
        student:[
        { 
          studenName:"jack",
          studenClass:6,
          studenSection:"a"
        },
        {
          studenName:"tom",
          studenClass:6,
          studenSection:"a"
        }
        ]
      
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
