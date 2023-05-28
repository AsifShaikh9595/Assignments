<template>
  <div>
    <h3>Home Component</h3>
  <div id="div">
    <div class="conter"><h4>Counter:</h4><h4>{{ $store.state.counter }}</h4></div><br>
    <div class="buttons">
      <button @click="$store.commit('increment')">Increment</button>
      <button @click="$store.commit('decrement')">Decrement</button>
      </div>
   
  </div>
  </div>
   

  
</template>

<script>
export default {

}
</script>

<style>
.conter {
    font-size: medium;
    color: blueviolet;
}
/* .buttons button{
  margin:0 10px;
} */
#div{
  text-align: center;
}


</style>