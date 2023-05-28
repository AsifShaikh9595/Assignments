<template>
  <div id="app" >
   
    <nav>
       <ul>
        
       <li><h3><router-link to="/">Home</router-link></h3></li>
      <li><h3><router-link to="/about">About</router-link></h3></li>
      <li><h3><router-link to="/contact">Contact</router-link></h3></li>
      <li><h3><router-link to="/signup">SignUp</router-link></h3></li>
     </ul>
    </nav>
    <router-view></router-view>
  
   
  </div>
</template>

<script>

export default {
  name: 'App',
  components: {
   
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #080808;
  margin-top: 60px;
}
nav{
  background-color:rgb(95, 74, 215);
  padding: 10px;
}
li {
  display: inline-block;
  margin-right: 10px;
}

</style>
