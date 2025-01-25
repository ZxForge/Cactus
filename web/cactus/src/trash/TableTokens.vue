<template>
    <div class="table">
        <div class="table_header" v-for="note of tableHeader" :key="note.public_token">
            <div><span class="table_text table_header_text">{{ note.choice }}</span></div>
            <div><span class="table_text table_header_text">{{ note.system }}</span></div>
            <div><span class="table_text table_header_text">{{ note.public_token }}</span></div>
            <div><span class="table_text table_header_text">{{ note.status }}</span></div>
            <div><span class="table_text table_header_text">{{ note.process }}</span></div>
            <div><span class="table_text table_header_text">{{ note.function }}</span></div>
        </div>
        <div class="table_row" v-for="note of tableRow" :key="note.public_token">
            <div><span class="table_text center_align"><input type="checkbox"></span></div>
            <div><span class="table_text">{{ note.system }}</span></div>
            <div><span class="table_text">{{ note.public_token }}</span></div>
            <div><span class="table_text">{{ note.status }}</span></div>
            <div><span class="table_text">{{ note.process }}</span></div>
            <div class="functions">
                <div>
                    <button @click="openEditModal" class="buttons">
                        <EditPencil />
                    </button>

                    <ModalEdit
                        :isOpen="isEditModalOpen"
                        :title="'Редактирование системы'"
                        :fields="fields"
                        :initialData="initialData"
                        @close="closeEditModal"
                        @save="saveEditModal"
                    >
                        <template #icon>
                            <EditPencil />
                        </template>                    
                    </ModalEdit>
                </div>
                <div>
                    <button @click="openModal" class="buttons">
                        <TrashIcon />
                    </button>

                    <ModalDelete
                        :isOpen="isModalOpen"
                        title="Удаление токена"
                        message="Вы действительно хотите удалить токен: (название)?"
                        confirmButtonText="Удалить"
                        cancelButtonText="Отменить"
                        @close="closeModal"
                        @confirm="handleConfirm"
                    >

                        <template #icon>
                            <EditPencil />
                        </template>
                    </ModalDelete>
                </div>
            </div>
        </div>
    </div>
</template> 

<script setup>
import TrashIcon from '@/components/icons/TrashIcon.vue';
import EditPencil from '@/components/icons/EditPencil.vue';
import ModalDelete from '@/components/ui/ModalDelete.vue';
import ModalEdit from './ui/ModalEdit.vue';
import { ref } from 'vue';

const props = defineProps({
    tableHeader: {
        type: Array,
        required: true,
    }, 
    tableRow: {
        type: Array,
        required: true,
    }
})

const isModalOpen = ref(false);

const openModal = () => {
  isModalOpen.value = true;
};

const closeModal = () => {
  isModalOpen.value = false;
};

const handleConfirm = () => {
  console.log('Действие подтверждено');
  closeModal();
};

const isEditModalOpen = ref(false);

const openEditModal = () => {
  isEditModalOpen.value = true;
};

const closeEditModal = () => {
  isEditModalOpen.value = false;
};

const saveEditModal = (data) => {
  console.log('Сохраненные данные:', data);
  closeEditModal();
};

// Поля для модального окна
const fields = [
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
      sms: 'SMS',
      oneC: '1C',
    },
  },
];

// Начальные данные
const initialData = {
  system_name: 'Моя система',
  description: 'Описание системы',
  status: true,
  processes: {
    email: true,
    telegram: false,
    push: true,
    sms: false,
    oneC: false,
  },
};
</script>

<style scoped>
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
    justify-content: flex-start;
    gap: 30px;
}

.functions div {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 25px;
    width: 25px;
    border: none;
}

.center_align{
    display: flex;
    justify-content: center;
}

.buttons{
    background-color: white;
    border: none;
    width: 20px;
    height: 20px;
    cursor: pointer;
}
</style>