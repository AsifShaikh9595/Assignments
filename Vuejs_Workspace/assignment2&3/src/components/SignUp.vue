<template>
    <div>
    <h3>Sign Up </h3>
  <form method="post">
    <fieldset>
        <legend>SignUp</legend>
        <label>Name:</label><input type="text" placeholder="Enter your name" v-model="user.name"/>
        <pre></pre>
        <label>Username:</label><input type="text" placeholder="Enter your Username" v-model="user.username"/>
        <pre></pre>
        <label>Password:</label><input type="password" placeholder="Enter password" v-model="user.password"/>
        <pre></pre>
        <label>Date Of Birth:</label><input type="date" placeholder="Enter DOB" v-model="user.dob"/>
        <pre></pre>
        <label>Email Id:</label><input type="email" placeholder="email" v-model="user.email"/>
        <pre></pre>
        <label>Phone number:</label><input type="tel" v-model="user.phone"/>
        <pre></pre>
        <label>Highest Graduation:</label><br><input list="highgrad" name="highgrd" id="highgrd" v-model="user.graduation"/>
        <datalist id="highgrad">
            <option value="SSC"></option>
            <option value="HSC"></option>
            <option value="Diploma"></option>
            <option value="Graduation"></option>
            <option value="Post-Graduation"></option>
            <option value="PHD"></option>
        </datalist><br>
        <pre></pre>
        <label>Address:</label><br><textarea rows="5" cols="30" v-model="user.address"></textarea>
        <pre></pre>
        <label>Gender:</label><br>
        <label for="male">Male</label><input type="radio" value="male" id="male" name="gender" v-model="user.gender"/>
        <label for="female">Female</label><input type="radio" value="female" id="female" name="gender" v-model="user.gender"/><br>
        <pre></pre>
        <label>Technologies with you are familiar</label><br>
        <input type="checkbox" name="java" id="java" value="java" v-model="user.technologies"/> <label for="java">JAVA</label>
        <input type="checkbox" name="c#" id="c#" value="c#" v-model="user.technologies"/> <label for="c#">C#</label><br>
        <input type="checkbox" name="vue" id="vue" value="vue" v-model="user.technologies"/> <label for="vue">Vue.js</label>
        <input type="checkbox" name="go" id="go" value="go" v-model="user.technologies"/> <label for="go">GoLang</label><br>
        <input type="checkbox" name="react" id="react" value="react" v-model="user.technologies"/> <label for="react">React</label>
        <input type="checkbox" name="mysql" id="mysql" value="mysql" v-model="user.technologies"/> <label for="mysql">MySQL</label><br>
        <input type="checkbox"  name="oracle" id="oracle" value="oracle" v-model="user.technologies"/> <label for="oracle">Oracle</label>
        <input type="checkbox" name="mongo" id="mongo" value="mongo" v-model="user.technologies"/> <label for="mongo">MongoDB</label><br> 
        <button type="submit">SUBMIT</button>
    </fieldset>
  </form>
  <DisplaySs :user="user"/>
</div>
</template>

<script>
import DisplaySs from '../components/DisplayComp.vue'
export default {
    data(){
        return{
            user:{
                name:null,
                username:null,
                password:null,
                dob:null,
                email:null,
                phone:null,
                    graduation:null,
                    address:null,
                    gender:null,
                    technologies:[]
            }

        }
    },
    components:{
        DisplaySs
    }

}
</script>

<style>
/* div{
    background-color: bisque;
} */

</style>