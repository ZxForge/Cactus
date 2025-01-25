<template>
    <div class="main_content">
        <div class="page_header">
            <HeaderPages :info_header="breadcrumbs" />
        </div>
        <div class="content_block">
            <div class="left_block" v-for="field of Fields" :key="field.public_token">
                <label for="public_token">{{field.public_token}}</label>
                <label for="private_token">{{field.private_token}}</label>
                <label for="system_name">{{field.system_name}}</label>
                <label for="description">{{field.description}}</label>
                <label for="processes">{{field.processes}}</label>
            </div>
            <div class="right_block">
                <input v-model="newToken.public_token" type="text" id="public_token" name="public_token" placeholder="Введите публичный токен" required>
                <input v-model="newToken.private_token" type="password" id="private_token" name="private_token" placeholder="Введите приватный токен" required>
                <input v-model="newToken.system_name" type="text" id="system_name" name="system_name" placeholder="Введите название системы" required>
                <input v-model="newToken.description" type="text" id="description" name="description" placeholder="Введите описание системы" required>
                <div v-for="(value, key) in newToken.processes" :key="key" class="checkbox_container"> 
                    <input v-model="newToken.processes[key]" type="checkbox" :id="key" :name="key">
                    <label :for="key">{{ key }}</label>
                </div>
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
    { url_info: 'Главная', url: '/', name_pages: 'Создать токен' },
    { url_info: 'Токены', url: '/tokens' },
    { url_info: 'Создать токен', url: '/tokens/create_token'},
]

interface FieldsInt{
    public_token: string
    private_token: string
    system_name: string
    description: string
    processes: string
}

const Fields: FieldsInt[] = [
    { 
        public_token: 'Публичный токен', 
        private_token: 'Приватный токен', 
        system_name: 'Название', 
        description: 'Описание', 
        processes: 'Процессы', 
    }
]

const newToken = ref({
    public_token: "",
    private_token: "",
    system_name: "", 
    description: "", 
    processes: {
        email: false,
        telegram: false,
        push: false,
        sms: false,
        oneC: false,
    }, 

})

const router = useRouter()

const submitData = () => {
    router.push({ 
        name: 'users', 
        params: { newUser: JSON.stringify(newToken.value) } 
    })
    resetData()
}

const resetData = () => {
    newToken.value = {
        public_token: "",
        private_token: "",
        system_name: "", 
        description: "", 
        processes: {
            email: false,
            telegram: false,
            push: false,
            sms: false,
            oneC: false,
        },
    }
}
</script>

<style scoped lang="less">
.main_content{
    width: 100vw;
    height: 100vh;
    box-sizing: border-box;
    padding: 10px 20px;
    background-color: rgba(250, 251, 252, 1);
}

.content_block {
    font-family: Arial, Helvetica, sans-serif;
    font-size: 18px;
    display: grid;
    grid-template-columns: 200px 1fr;
    gap: 16px;
    margin: 10px;
    color: black;
}

.left_block label {
    height: 40px;
    display: flex;
    justify-content: left;
    align-items: center;
    font-size: 18px;
    margin-bottom: 16px;
}

.right_block label{
    font-size: 18px;
}

.right_block input[type="text"],
.right_block input[type="password"] {
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

.checkbox_container{
    display: flex;
    align-items: center;
    font-size: 18px;
    margin: 12px 0 17px 0; 
}
</style>