import Vuex from 'vuex';
import Vue from 'vue';

Vue.use(Vuex);

const store= new Vuex.Store({
    state:{
        counter:0
    },
    mutations:{
        increment(state){
            state.counter++
        },
        decrement(state){
            state.counter--
        }
    }
})
export default store;
 