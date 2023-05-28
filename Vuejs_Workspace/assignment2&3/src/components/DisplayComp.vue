<template>
    <p>
        <label>Name: </label>{{ user.name }}
       <label>UserName:</label> {{ user.username }}
       <label>Password:</label> {{ user.password }}<br>
       <label>Date of Birth:</label> {{ user.dob }}<br>
       <label>Email:</label> {{ user.email }}<br>
       <label>Phone:</label>{{ user.phone }}<br>
       <label>Graduation:</label>{{ user.graduation }}<br>
       <label>Address:</label>{{ user.address }}<br>
       <label>Gender:</label>{{ user.gender }}<br>
       <ol v-for="t in user.technologies" :key="t.technologies">
                {{ t }}
    </ol>
        

    </p>
  
</template>

<script>
export default {
    props:["user"]

}
</script>

<style>
form{
    background-color:rgb(120, 106, 165)
}

</style>