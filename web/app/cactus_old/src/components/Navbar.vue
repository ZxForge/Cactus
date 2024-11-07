<template>
    <div class="left_navigation">
        <div class="header">
            <div class="logo_and_name">
                <div class="logo">
                    <LogoIcon />
                </div>
                <div class="name">
                    <RouterLink to="/">Cactus</RouterLink>
                </div>
            </div>
        </div>

        <div class="select_menu">
            <div class="container">
                <NavbarButton to="/">
                    <template #icon><DashboardIcon /></template>
                    <template #title>Панель управления</template>
                </NavbarButton>
                <NavbarButton to="/process/">
                    <template #icon><ProcessIcon /></template>
                    <template #title>Процессы</template>
                    <template #subbutton="{ to }">
                        <RouterLink v-for="subTo of listSubProcessButton" :key="subTo" :class="{active:activeRouter(`${to}${subTo}/`)}" class="subbutton" :to="`${to}${subTo}/`">
                            <DotIcon class="icon" />
                            <span>{{ subTo }}</span>
                        </RouterLink>
                    </template>
                </NavbarButton>
                <NavbarButton to="/users/">
                    <template #icon><UsersIcon /></template>
                    <template #title>Пользователи</template>
                </NavbarButton>
                <NavbarButton to="/tokens/">
                    <template #icon><TokenIcon /></template>
                    <template #title>Токены</template>
                </NavbarButton>
                <NavbarButton to="/tools/">
                    <template #icon><ToolsIcon /></template>
                    <template #title>Настройки</template>
                </NavbarButton>
                <NavbarButton to="/docs/">
                    <template #icon><DocumentationIcon /></template>
                    <template #title>Документация</template>
                </NavbarButton>
            </div>
        </div>

        <div class="footer">
            <div class="logo_prof_and_name_user_and_exit">
                <div class="logo_user">
                    <img src="../src/assets/img/ava.png" alt="" />
                </div>
                <div class="name_user">
                    <RouterLink class="user_name_text" to="/"
                        >Zalberix</RouterLink
                    >
                </div>
                <div class="exit">
                    <RouterLink class="user_name_text" to="/">
                        <ExitIcon />
                    </RouterLink>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import LogoIcon from "./Icon/LogoIcon.vue";
import ExitIcon from "./Icon/ExitIcon.vue";
import DashboardIcon from "./Icon/DashboardIcon.vue";
import ProcessIcon from "./Icon/ProcessIcon.vue";
import DocumentationIcon from "./Icon/DocumentationIcon.vue";
import ToolsIcon from "./Icon/ToolsIcon.vue";
import TokenIcon from "./Icon/TokenIcon.vue";
import UsersIcon from "./Icon/UsersIcon.vue";
import NavbarButton from "./UI/NavbarButton.vue";
import DotIcon from "./Icon/DotIcon.vue";

import { useRoute } from "vue-router"

const route = useRoute();

const listSubProcessButton = ['email', 'telegram', 'sms', 'push', '1c'];

function activeRouter (url) {
    return route.path === url
}
</script>

<style scoped>
.left_navigation {
    width: 199px;
    height: 100vh;
    background-color: rgba(17, 20, 61, 1);
    position: absolute;
}

.header {
    width: 199px;
    height: 70px;
    background-color: rgba(41, 43, 80, 1);
}

.logo_and_name {
    width: 170px;
    height: 60px;
    margin: 0 auto;
    padding: 15px;
    display: flex;
    position: relative;
    align-items: center;
}
.logo {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background-color: #fff;
}

.logo svg {
    margin: 1px 0 0 8px;
}

.button.active ~ .menu {
    line-height: 1.2;
    color: currentColor;
}

.name {
    position: absolute;
    right: 35px;
    top: 25px;
    color: #fff;
    text-align: center;
    font-family: "Inter";
    text-transform: uppercase;
}

.container {
    width: 190px;
    margin: 15px auto;
}

.footer {
    width: 100%;
    height: 70px;
    background-color: rgba(41, 43, 80, 1);
    position: absolute;
    bottom: 0;
}
.logo_prof_and_name_user_and_exit {
    width: 200px;
    margin: auto;
    display: flex;
}
.logo_user {
    padding: 10px;
}
.logo_user img {
    width: 40px;
    height: 40px;
    border-radius: 50%;
}

.name_user {
    margin-top: 20px;
    margin-left: 20px;
    color: #fff;
}

.user_name_text {
    color: #fff;
    text-decoration: none;
}

.exit {
    margin-top: 20px;
    margin-left: 20px;
}
</style>