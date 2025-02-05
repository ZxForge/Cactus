<template>
    <div class="main_content">
        <div class="page_header">
            <HeaderPages :info_header="breadcrumbs" />
        </div>
        <div class="content_block">
            <div class="left_block">
                <label for="name">Название</label>
                <label for="description">Описание</label>
                <label for="status">Статус</label>
                <label for="priority">Приоритет</label>
            </div>
            <div class="right_block">
                <input v-model="newUser.email" type="text" id="name" name="name" placeholder="Введите название" required>
                <input v-model="newUser.password" type="text" id="description" name="description" placeholder="Введите пароль" required>
                <div>
                    <input v-model="newUser.changePassword" type="checkbox" id="status" name="status">
                    <label class="text" for="status">Активный</label>
                </div>
                <select v-model="newUser.role" id="priority" name="priority" required>
                    <option value="" disabled selected>Не выбрано</option>
                    <option value="Администратор">Администратор</option>
                    <option value="Модератор">Модератор</option>
                    <option value="Гость">Гость</option>
                    <option value="Пользователь">Пользователь</option>
                </select>
            </div>
        </div>
        <div class="buttons">
            <button type="reset" @click="resetData()" class="resetButton">Очистить</button>
            <button type="submit" @click="submitData()" class="createButton">Создать</button>
        </div>
    </div>
</template>


<script setup lang="ts">
import HeaderPages from '@/components/HeaderPages.vue'
import { useRouter } from 'vue-router'
import { ref } from "vue"
 
interface InfoHeader {
    url_info: string
    url: string
    name_pages?: string
}

const breadcrumbs: InfoHeader[] = [
    { url_info: 'Главная', url: '/', name_pages: 'Создать систему' },
    { url_info: 'Системы', url: '/systems'},
    { url_info: 'Создать систему', url: '/systems/create'},
]

const newUser = ref({
    email: "",
    login: "",
    password: "", 
    role: "", 
    changePassword: false, 
})

const router = useRouter()

const submitData = () => {
    router.push({ 
        name: 'users', 
        params: { newUser: JSON.stringify(newUser.value) } 
    })
    resetData()
}

const resetData = () => {
    newUser.value = {
        email: "", 
        login: "", 
        password: "", 
        role: "", 
        changePassword: false,
    }
}
</script>

<style scoped lang="less">
/* Основной контейнер */
.main_content{
    width: 100vw;
    height: 100vh;
    box-sizing: border-box;
    padding: 10px 20px;
    background-color: rgba(250, 251, 252, 1);
}

/* Гридовая структура контента */
.content_block {
    font-family: Arial, Helvetica, sans-serif;
    font-size: 18px;
    display: grid;
    grid-template-columns: 200px 1fr;
    gap: 16px;
    margin: 10px;
    color: black;
}

/* Левый блок с текстом меток */
.left_block label {
    height: 40px;
    display: flex;
    justify-content: left;
    align-items: center;
    font-size: 18px;
    margin-bottom: 16px;
}

/* Правый блок с полями ввода */
.right_block input[type="text"],
.right_block input[type="email"],
.right_block input[type="password"],
.right_block div,
.right_block select {
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 18px;
    margin-bottom: 17px; 
}

.right_block div{
    border: none;
}

.right_block input[type="checkbox"] {
    margin-right: 10px;
}

.text {
    font-size: 18px;
}

/* Кнопки */
.buttons {
    display: flex;
    justify-content: end;
    gap: 10px;
    padding: 0 10px 0 0;
}

button {
    height: 40px;
    width: 160px;
    border: none;
    border-radius: 8px;
    font-size: 18px;
    cursor: pointer;
}

.resetButton {
    border: 1px solid rgba(146, 147, 171, 1);
    background-color: rgba(255, 255, 255, 1);
    color: rgba(146, 147, 171, 1);
    font-weight: bold;
}

.createButton {
    background-color: rgba(85, 111, 246, 1);
    color: rgba(255, 255, 255, 1);
}

</style>