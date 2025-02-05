<template>
    <div>
        <div class="main_content">
            <HeaderPages :info_header="breadcrumbs" />
            <div class="content_block">
                <div v-for="element of blockName" :key="element.id" class="block">
                    <RouterLink :to="{ name: element.slug }">
                        <span class="text">{{ element.name }}</span>
                    </RouterLink>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

</script>

<style>

</style>
