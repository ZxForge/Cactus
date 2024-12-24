<template>
    <div>
        <div class="main_content">
            <div class="page_header">
                <HeaderPages :info_header="breadcrumbs" />
                <RouterLink to="/users/createnewuser">
                    <button class="create_user">
                        <span class="create_user_text">Создать пользователя</span>
                    </button>
                </RouterLink>
            </div>
            <div class="content_block">
                <div class="table">
                    <div class="table_header" v-for="note of tableHeader" :key="note.id">
                        <div><span class="table_text table_header_text">{{ note.id }}</span></div>
                        <div><span class="table_text table_header_text">{{ note.email }}</span></div>
                        <div><span class="table_text table_header_text">{{ note.login }}</span></div>
                        <div><span class="table_text table_header_text">{{ note.role }}</span></div>
                        <div><span class="table_text table_header_text">{{ note.status }}</span></div>
                        <div><span class="table_text table_header_text">{{ note.function }}</span></div>
                    </div>
                    <div class="table_row" v-for="note of tableRow" :key="note.id">
                        <div><span class="table_text">{{ note.id }}</span></div>
                        <div><span class="table_text">{{ note.email }}</span></div>
                        <div><span class="table_text">{{ note.login }}</span></div>
                        <div><span class="table_text">{{ note.role }}</span></div>
                        <div><span class="table_text">{{ note.status }}</span></div>
                        <div class="functions">
                            <RouterLink to="/edituser">
                                <EditPencil />
                            </RouterLink>
                            <RouterLink to="/deleteuser">
                                <TrashIcon /> 
                            </RouterLink>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import TrashIcon from '@/components/icons/TrashIcon.vue';
import EditPencil from '@/components/icons/EditPencil.vue';
import HeaderPages from '@/components/HeaderPages.vue'
import { useRoute, useRouter } from 'vue-router'
import { ref, onMounted, watchEffect} from 'vue'

interface InfoHeader {
    url_info: string
    url: string
    name_pages?: string
}

const breadcrumbs: InfoHeader[] = [
    { url_info: 'Главная', url: '/', name_pages: 'Пользователи' },
    { url_info: 'Пользователи', url: '/users' },
]

interface Table {
    id?: string
    email: string
    login: string
    password?: string,
    role: string
    status: string
    function?: string
    changePassword?: boolean
}

const tableHeader: Table[] = [
    { id: '№',email: 'Почта', login: 'Логин', role: 'Роль', status: 'Статус',  function: 'Функции'},
]
const tableRow: Table[] = [
    { id: '1',email: 'petrov@gmail.com', login: 'Petr', role: 'Пользователь', status: 'Активный'},
    { id: '2',email: 'famin@gmail.com', login: 'fam', role: 'Администратор', status: 'Активный'},
]

/*const route = useRoute()
const newUser = ref({})
watchEffect(() => {
  if (route.query.newUser) {
    newUser.value = JSON.parse(route.query.newUser)
  }
})*/

let currentId = ref(tableRow.length + 1);
function addUser(email: string, login: string, role: string, status: string = "Активный") {
    tableRow.push({
        id: currentId.value.toString(),
        email,
        login,
        role,
        status,
    });
    currentId.value++;
}
</script>

<style scoped>
.main_content{
    width: 100%;
    height: 100%;
    box-sizing: border-box;
    padding: 10px 20px;
    background-color: rgba(250, 251, 252, 1);
}

.page_header{
    width: 100%;
    display: flex;
    justify-content: space-between;
    gap: 10px;
}

.create_user{
    height: 30px;
    width: 231px;
    margin-top: 20px;
    border: none;
    border-radius: 8px;
    color: white;
    background-color: #556FF6;
}

.create_user_text{
    width: 191px;
    height: 17px;
    font-family: 'Inter', sans-serif;
    font-weight: 700;
    font-size: 14px;
    line-height: 16.94px;
    text-transform: uppercase;
    letter-spacing: 0.5px;
    cursor: pointer;
}

.content_block{
    margin-top: 5px;
}

.table{
    background-color: rgba(255, 255, 255, 1);
    display: grid;
    grid-template-columns: 75px 1fr 1fr 1fr 1fr 150px;
    border: 1px solid rgba(146, 147, 171, 0.5);
    border-radius: 10px;
    padding: 10px;
}
.table_header, .table_row {
    display: contents;
}

.table_header div,
.table_row div {
    padding: 15px;
    border-bottom: 1px solid #e5e5e5;
    text-align: left;
}

.table_row:last-child div {
    border-bottom: none;
}

.table_text{
    width: 37px;
    height: 17px;
    font-family: 'Inter', sans-serif;
    font-weight: 400;
    font-size: 16px;
    line-height: 16.94px;
}

.table_header_text{
    font-weight: bold;
}

.functions{
    display: flex;
    gap: 30px;
}
</style>