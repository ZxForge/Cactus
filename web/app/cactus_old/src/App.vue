<template>
    <div>
        <Navbar />
        <RouterView />
    </div>
</template>

<style>
@import "../src/assets/css/main.css";
</style>

<script setup>
import Navbar from "./components/Navbar.vue";
</script>
