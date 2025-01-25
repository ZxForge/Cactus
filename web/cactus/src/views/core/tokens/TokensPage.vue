<template>
    <div>
        <div class="main_content">
            <div class="page_header">
                <HeaderPages :info_header="breadcrumbs" />
                <RouterLink to="/tokens/create_token">
                    <button class="create_user">
                        <span class="create_user_text">Создать токен</span>
                    </button>
                </RouterLink>
            </div>
            <div class="content_block">
                <TableData 
                    :headers="headersForTokens" 
                    :rows="rowsForTokens" 
                    :pageName="pageName"
                    :editFieldsForToken="editFieldsForToken"
                    :initialDataForToken="initialData"
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
    { url_info: 'Главная', url: '/', name_pages: 'Токены' },
    { url_info: 'Токены', url: '/tokens' },
]

const pageName = 'Токены'

const headersForTokens = [
  { key: 'choice', label: 'Выбор', type: 'checkbox' },
  { key: 'system', label: 'Система', type: 'text' },
  { key: 'public_token', label: 'Публичный токен', type: 'text' },
  { key: 'status', label: 'Статус', type: 'status' },
  { key: 'processes', label: 'Процессы', type: 'processes' },
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

const rowsForTokens = [
  {
    choice: false,
    system: 'agt72.ru',
    public_token: 'agt1234',
    status: 'Активен',
    processes: { email: true, telegram: false, push: true },
  },
  {
    choice: false,
    system: 'dom72.ru',
    public_token: 'dom1234',
    status: 'Активен',
    processes: { email: true, telegram: false, push: true },
  },
];

const editFieldsForToken = [
  { key: 'system_name', label: 'Название', type: 'text', placeholder: 'Введите название системы' },
  { key: 'description', label: 'Описание', type: 'text', placeholder: 'Введите описание системы' },
  { key: 'status', label: '', labelcb: 'Активен',type: 'checkbox' },
  {
    key: 'processes',
    label: 'Процессы',
    type: 'checkbox-group',
    options: {
      email: 'Email',
      telegram: 'Telegram',
      push: 'Push',
    },
  },
];

const initialData = {
  system_name: '',
  description: '',
  status: true,
  processes: {
    email: false,
    telegram: false,
    push: false,
  },
};
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