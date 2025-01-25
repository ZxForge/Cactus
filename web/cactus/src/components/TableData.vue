<template>
    <div class="table">
        <!-- Заголовок таблицы -->
        <div class="table_header">
          <div v-for="header in headers" :key="header.key" class="table_header_cell">
            <span class="table_text table_header_text">{{ header.label }}</span>
          </div>
        </div>
  
        <!-- Строки таблицы -->
        <div class="table_row" v-for="(row, rowIndex) in rows" :key="rowIndex">
          <div v-for="header in headers" :key="header.key" class="table_cell">

            <!-- Чекбокс для выбора -->
            <span v-if="header.type === 'checkbox'" class="table_text center_align">
              <input type="checkbox" v-model="row[header.key]" />
            </span>

            <!-- Текстовое поле -->
            <span v-else-if="header.type === 'text'" class="table_text">
              {{ row[header.key] }}
            </span>

            <!-- Статус (текст или иконка) -->
            <span v-else-if="header.type === 'status'" class="table_text">
              {{ row[header.key] }}
            </span>

            <!-- Процессы (список) -->
            <span v-else-if="header.type === 'processes'" class="table_text">
              {{ formatProcesses(row[header.key]) }}
            </span>

            <!-- Функции (кнопки) -->
            <div v-else-if="header.type === 'functions'" class="functions">
                <button
                    v-for="(action, actionIndex) in header.actions"
                    :key="actionIndex"
                    @click="openEditModal"
                    class="buttons"
                >   
                    <component :is="action.icon" />
                </button>

              <!-- Модальное окно редактирования -->
                <ModalEdit
                  v-if="selectedRow"
                  :isOpen="isEditModalOpen"
                  :title="`Редактирование ${title}`"
                  :fields="fields"
                  :initialData="initialData"
                  @close="closeEditModal"
                  @save="saveEditModal"
                >
                  <template #icon>
                    <EditPencil />
                  </template>
                </ModalEdit>
            
                <!-- Модальное окно удаления -->
                <ModalDelete
                  v-if="selectedRow"
                  :isOpen="isDeleteModalOpen"
                  :title="`Удаление ${title}`"
                  :message="`Вы действительно хотите удалить токен: ${selectedRow.system}?`"
                  confirmButtonText="Удалить"
                  cancelButtonText="Отменить"
                  @close="closeDeleteModal"
                  @confirm="handleConfirm"
                >
                  <template #icon>
                    <TrashIcon />
                  </template>
                </ModalDelete>
            </div>
          </div>
        </div> 
    </div>
</template>
  
<script setup>
import { ref, defineProps } from 'vue';
import TrashIcon from '@/components/icons/TrashIcon.vue';
import EditPencil from '@/components/icons/EditPencil.vue';
import ModalDelete from './ui/ModalDelete.vue';
import ModalEdit from './ui/ModalEdit.vue';
  
  // Пропсы
const props = defineProps({
  headers: {
    type: Array,
    required: true,
  },
  rows: {
    type: Array,
    required: true,
  },
  editFieldsForToken: {
    type: Array,
    required: true,
  },
  editFieldsForUser: {
    type: Array,
    required: true,
  },
  pageName: {
    type: String, 
    required: true,
  },
  initialDataForToken: {
    type: Array,
    required: true,
  },
  initialDataForUser: {
    type: Array,
    required: true,
  },
});

const getModalSettings = (pageName) => {
  const lowerPageName = pageName.toLowerCase();

  const wordForms = {
    'пользователи': 'пользователя',
    'токены': 'токена',
  };

  const fieldsMapping = {
    'пользователи': props.editFieldsForUser,
    'токены': props.editFieldsForToken,
  };

  const initialData = {
    'пользователи': props.initialDataForUser,
    'токены': props.initialDataForToken,
  };

  return {
    title:  wordForms[lowerPageName] || lowerPageName,
    fields: fieldsMapping[lowerPageName] || [],
    initialData: initialData[lowerPageName] || [],
  };
};

const { title, fields, initialData } = getModalSettings(props.pageName);


const isEditModalOpen = ref(false);
const isDeleteModalOpen = ref(false);
const selectedRow = ref(null);

// Форматирование процессов (если нужно)
const formatProcesses = (processes) => {
  if (!processes) return '';
  return Object.entries(processes)
    .filter(([_, value]) => value)
    .map(([key]) => key)
    .join(', ');
};

// Функции для открытия модальных окон
const openEditModal = (row) => {
  selectedRow.value = row;
  isEditModalOpen.value = true;
};

const openDeleteModal = (row) => {
  selectedRow.value = row;
  isDeleteModalOpen.value = true;
};

// Функции для закрытия модальных окон
const closeEditModal = () => {
  isEditModalOpen.value = false;
};

const closeDeleteModal = () => {
  isDeleteModalOpen.value = false;
};

// Обработка сохранения изменений
const saveEditModal = (data) => {
  console.log('Сохранённые данные:', data);
  closeEditModal();
};

// Обработка удаления
const handleConfirm = () => {
  console.log('Удалённые данные:', selectedRow.value);
  closeDeleteModal();
};
</script>
  
<style scoped>
  .table {
    background-color: rgba(255, 255, 255, 1);
    display: grid;
    grid-template-columns: repeat(6, minmax(0, 1fr));
    border: 1px solid rgba(146, 147, 171, 0.5);
    border-radius: 10px;
    padding: 10px;
  }
  
  .table_header,
  .table_row {
    display: contents;
  }

  .table_cell:first-child{
    padding-left: 35px;
  }

  .table_header,
  .table_row {
    display: contents;
  }
  
  .table_header_cell,
  .table_cell {
    padding: 15px;
    border-bottom: 1px solid #e5e5e5;
    text-align: left;
  }
  
  .table_row:last-child .table_cell {
    border-bottom: none;
  }
  
  .table_text {
    width: 37px;
    height: 17px;
    font-family: 'Inter', sans-serif;
    font-weight: 400;
    font-size: 16px;
    line-height: 16.94px;
  }
  
  .table_header_text {
    font-weight: bold;
  }
  
  .functions {
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
  
  .buttons {
    background-color: white;
    border: none;
    width: 20px;
    height: 20px;
    cursor: pointer;
  }
</style>