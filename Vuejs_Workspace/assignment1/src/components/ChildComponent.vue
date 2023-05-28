<template>
    <div>
        <h2>From ChildComponent</h2>
        <h1>{{msg}}</h1>
    </div>
</template>
<script>
    export default{
        data(){
            return{

            }
        },
        props:{
            msg:String
        }
    }
</script>
<style>

</style>