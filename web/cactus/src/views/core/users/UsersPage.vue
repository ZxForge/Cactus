<template>
    <div>
        <div class="main_content">
            <div class="page_header">
                <HeaderPages :info_header="breadcrumbs" />
                <RouterLink to="/users/create_user">
                    <button class="create_user">
                        <span class="create_user_text">Создать пользователя</span>
                    </button>
                </RouterLink>
            </div>
            <div class="content_block">
                <TableData 
                    :headers="headersForUsers" 
                    :rows="rowsForUsers" 
                    :pageName="pageName"
                    :editFieldsForUser="editFieldsForUser"
                    :initialDataForUser="initialData"
                />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import HeaderPages from '@/components/HeaderPages.vue'
import TableData from '@/components/TableData.vue';
import TrashIcon from '@/components/icons/TrashIcon.vue';
import EditPencil from '@/components/icons/EditPencil.vue';

interface InfoHeader {
    url_info: string
    url: string
    name_pages?: string
}

const breadcrumbs: InfoHeader[] = [
    { url_info: 'Главная', url: '/', name_pages: 'Пользователи' },
    { url_info: 'Пользователи', url: '/users' },
]

const pageName = 'Пользователи'

const headersForUsers = [
  { key: 'id', label: 'Номер', type: 'text' },
  { key: 'email', label: 'Почта', type: 'text' },
  { key: 'login', label: 'Логин', type: 'text' },
  { key: 'role', label: 'Роль', type: 'text' },
  { key: 'status', label: 'Статус', type: 'status' },
  {
    key: 'functions',
    label: 'Функции',
    type: 'functions',
    actions: [
      {
        icon: EditPencil,
        handler: 'openEditModal',
      },
      {
        icon: TrashIcon,
        handler: 'openDeleteModal',
      },
    ],
  },
];

const rowsForUsers = [
  {
    id: 1,
    email: 'user1@example.com',
    login: 'user1',
    role: 'Администратор',
    status: 'Активен',
  },
  {
    id: 2,
    email: 'user2@example.com',
    login: 'user2',
    role: 'Пользователь',
    status: 'Неактивен',
  },
];

const editFieldsForUser = [
  { key: 'login', label: 'Логин', type: 'text', placeholder: 'Введите новый логин' },
  { key: 'email', label: 'Почта', type: 'text', placeholder: 'Введите новую почту' },
  { key: 'role', label: 'Роль', type: 'select', options: {
    admin: 'Администратор', 
    user: 'Пользователь', 
  }},
  { key: 'status', label: '', labelcb: 'Активен', type: 'checkbox' },
  { key: 'change_password', label: 'Новый пароль', type: 'password'},
  { key: 'change', label: '', labelcb: 'Сменить при входе', type: 'checkbox' },
];

const initialData = {
    login: '',
    email: '', 
    role: '', 
    status: false, 
    change_password: '', 
    change: false,
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
</style>