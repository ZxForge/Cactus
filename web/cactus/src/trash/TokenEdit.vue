<template>
    <div class="user_edit_modal">
        <div class="user_edit_modal_backdrop">
            <div class="user_edit_modal_container">
                <header class="modal_header">
                    <div class="left_block_header">
                        <EditPencil style="cursor: default;"/>
                        <h2 class="modal_title">Редактирование - (Название системы)</h2>
                    </div>
                    <div class="left_block_header">
                        <RouterLink to="/users">
                            <button class="close_button"><CrossIcon /></button>
                        </RouterLink>
                    </div>
                </header>
                <div class="main_content">
                    <main class="modal_body">
                        <div class="content_block">
                            <div class="left_block"  v-for="field of Fields" :key="field.system_name">
                                <label for="system_name">{{field.system_name}}</label>
                                <label for="description">{{field.description}}</label>
                                <label for="status">{{field.status}}</label>
                                <label for="processes">{{field.processes}}</label>
                            </div>
                            <div class="right_block">
                                <input v-model="newToken.system_name" type="text" id="system_name" name="system_name" placeholder="Введите название системы" required>
                                <input v-model="newToken.description" type="text" id="description" name="description" placeholder="Введите описание системы" required>
                                <div class="checkbox_style">
                                    <input v-model="newToken.status" type="checkbox" id="status" name="status">
                                    <label for="status">Активен</label>
                                </div>
                                <div class="checkbox_container">
                                    <div class="checkbox_style" v-for="(value, key) in newToken.processes" :key="key"> 
                                        <input v-model="newToken.processes[key]" type="checkbox" :id="key" :name="key">
                                        <label :for="key">{{ key }}</label>
                                    </div>
                                </div>
                                
                            </div>
                        </div>
                    </main>
                    <footer class="modal_footer">
                        <RouterLink to="/users">
                            <button class="cancel_button">Отмена</button>
                        </RouterLink>
                        <button class="save_button">Изменить</button>
                    </footer>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import CrossIcon from '@/components/icons/CrossIcon.vue';
import EditPencil from '@/components/icons/EditPencil.vue';

import HeaderPages from '@/components/HeaderPages.vue'
import { useRouter } from 'vue-router'
import { ref } from "vue"

interface FieldsInt{
    system_name: string
    description: string
    status: string
    processes: string
}

const Fields: FieldsInt[] = [
    { 
        system_name: 'Название', 
        description: 'Описание', 
        status: 'Статус',
        processes: 'Процессы', 
    }
]

const newToken = ref({
    system_name: "", 
    description: "", 
    status: false,
    processes: {
        email: false,
        telegram: false,
        push: false,
        sms: false,
        oneC: false,
    }, 
})
</script>

<style scoped>
.user_edit_modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 1000;
}

.user_edit_modal_backdrop {
    position: absolute;
    width: 100%;
    height: 100%;
    background-color: rgba(22, 44, 67, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
}

.user_edit_modal_container {
    width: 605px;
    height: 590px;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    display: flex;
    flex-direction: column;
    padding: 16px;
}

.modal_header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.left_block_header {
    display: flex;
    align-items: center;
    gap: 16px;
}

.modal_title {
    font-size: 20px;
    font-weight: 400;
    line-height: 24px;
    margin: 0;
}

.close_button {
    background: none;
    border: none;
    font-size: 24px;
    cursor: pointer;
    line-height: 1;
    color: #aaa;
}

.main_content {
    width: 100%;
    padding: 0;
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.modal_body {
    display: flex;
    flex-direction: column;
    padding-top: 16px;
}

.content_block {
    display: grid;
    grid-template-columns: 200px 1fr;
    gap: 16px;
    margin: 10px;
}

.left_block {
    display: flex;
    flex-direction: column;
    gap: 16px;
}

.left_block label {
    height: 40px;
    display: flex;
    align-items: center;
    font-size: 18px;
}

.right_block label{
    font-size: 18px;
}

.right_block {
    display: flex;
    flex-direction: column;
    gap: 16px;
}

.right_block input[type="text"] {
    height: 40px;
    width: 100%;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 18px;
}

.right_block #status {
    height: 40px;
    font-size: 18px;
}

.checkbox_style {
    display: flex;
    align-items: center;
    gap: 8px;
    
} 

.checkbox_container{
    display: flex;
    flex-direction: column;
    padding-top: 12px;
    gap: 15px;
}

.modal_footer {
    display: flex;
    justify-content: flex-end;
    gap: 10px;
}

.cancel_button {
    width: 100px;
    height: 40px;
    padding: 10px 20px;
    font-size: 14px;
    font-weight: 700;
    color: #9293ab;
    background-color: #fff;
    border: 1px solid #9293ab;
    border-radius: 8px;
    cursor: pointer;
}

.save_button {
    width: 110px;
    height: 40px;
    padding: 10px 20px;
    font-size: 14px;
    font-weight: 700;
    color: #fff;
    background-color: #556ff6;
    border: none;
    border-radius: 8px;
    cursor: pointer;
}
</style>