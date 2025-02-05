<template>
    <div>
        <div class="main_content">
            <div class="page_header">
                <HeaderPages :info_header="breadcrumbs" />
                <RouterLink to="/tokens/create">
                    <button class="create_user">
                        <span class="create_user_text">Создать токен</span>
                    </button>
                </RouterLink>
            </div>
            <div class="content_block">
                <TableData 
                    :headers="headersForTokens" 
                    :rows="rowsForTokens" 
                    :columnRows="columnRows"
                    @action="(data: any) => modalFunction(data)"
                />
            </div>
        </div>
        <ModalWindow 
          v-if="isOpen"
          :title="eventModalContent[activeModalEvent]['title']"
          @close="closeModal"
        >
          <template #titleicon>
            <component :is="eventModalContent[activeModalEvent]['icon']"></component>
          </template>
          <template #crossicon>
            <component :is="eventModalContent[activeModalEvent]['crossicon']"></component>
          </template>
          <template #body>
            <component 
                :is="eventModalContent[activeModalEvent]['component']" 
                v-bind="eventModalContent[activeModalEvent]['bind']"
                @close="closeModal"
            >
            </component>
          </template>
          <template #footer>
            <button 
              v-for="(button, key) in eventModalContent[activeModalEvent]['buttons']" 
              :key="key"
              :class="`buttons ${button.style}`"
              :style="button.customStyle"
              @click="button.function"
            >
              {{ button.name }}
            </button>
          </template>
        </ModalWindow>
    </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import HeaderPages from '@/components/HeaderPages.vue'
import TableData from '@/components/TableData.vue';
import TrashIcon from '@/components/icons/TrashIcon.vue';
import DeleteForm from '@/components/DeleteForm.vue';
import EditPencil from '@/components/icons/EditPencil.vue';
import EditForm from '@/components/EditForm.vue';
import CrossIcon from '@/components/icons/CrossIcon.vue';
import ModalWindow from '@/components/ui/ModalWindow.vue';

interface InfoHeader {
    url_info: string
    url: string
    name_pages?: string
}

const breadcrumbs: InfoHeader[] = [
    { url_info: 'Главная', url: '/', name_pages: 'Токены' },
    { url_info: 'Токены', url: '/tokens' },
]

const columnRows = 6; 

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
        type: 'edit',
        icon: EditPencil,
      },
      {
        type: 'delete',
        icon: TrashIcon,
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

const editFields = [
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

interface ActionData {
  actionType: string; // или другой тип, если actionType не строка
  selectedRow: any; // уточните тип, если возможно
  isOpen: boolean;
}

const activeModalEvent = ref('')
const selectedRow = ref(undefined)
const isOpen = ref(false)

const modalFunction = (actionData: ActionData) => {
  activeModalEvent.value = actionData.actionType
  selectedRow.value = actionData.selectedRow
  isOpen.value = actionData.isOpen
  console.log(selectedRow)

}

const closeModal = () => {
  isOpen.value = false;
  activeModalEvent.value = ''; 
};

const saveChanges = () => {
  console.log('Данные успешно изменены')
  isOpen.value = false;
  activeModalEvent.value = ''; 
};

const deleteElement = () => {
  console.log('Элемент удален')
  isOpen.value = false;
  activeModalEvent.value = ''; 
};

const eventModalContent = {
  'edit': {
    "icon": EditPencil,
    "crossicon": CrossIcon,
    "title":"Редактировать токен",
    "component": EditForm,
    "bind": {
      fields: editFields,
      
    },
    "buttons": {
      "cancel": {
        "name": "Отмена",
        "style": "cancel_button",
        "customStyle": {},
        "function": () => closeModal()
      },
      "save": {
        "name": "Сохранить",
        "style": "save_button",
        "customStyle": {},
        "function": () => saveChanges()
      },
      
      }
    }, 

  'delete': {
    "icon": TrashIcon,
    "crossicon": CrossIcon,
    "title":"Удалить токен",
    "component": DeleteForm,
    "bind": {
      modalTitle: `Вы уверены, что хотите удалить токен`,
      itemName: columnRows,
      
    },
    "buttons": {
      "cancel": {
        "name": "Отмена",
        "style": "cancel_button",
        "customStyle": {},
        "function": () => closeModal()
      },
      "delete": {
        "name": "Удалить",
        "style": "delete_button",
        "customStyle": {},
        "function": () => deleteElement()
      },
    }, 
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