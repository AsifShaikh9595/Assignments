<template>
    <div>
        <h1>{{msg}}</h1>
        <table border="1">
            <tr>
                <td>Name</td>
                <td>Class</td>
                <td>Section</td>
            </tr>
            <tr v-for="stud in myarr" :key="stud.studenName"> 
                <td>{{stud.studenName}}</td>
                <td>{{stud.studenClass}}</td>
                <td>{{stud.studenSection}}</td>
            </tr>
        </table>
        <ChildComponent :msg="message"/>

    </div>
</template>
<script>
import ChildComponent from '../components/ChildComponent.vue'
    export default{
        components:{
            ChildComponent
        },
        props:{
            msg:String,
            myarr:Array
        },
       data(){
        return{
            message:this.msg,
           
        }
       }
       
    }
</script>

<style>
table{
   margin-left: auto;
   margin-right: auto;
}

</style>