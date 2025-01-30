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
                    @click="handleActionClick(action, row)"
                    class="buttons"
                >   
                    <component :is="action.icon" />
                </button>

              <!-- Универсальное модальное окно -->
                <ModalWindow 
                  v-if="selectedRow"
                  :isOpen = "isModalOpen"
                  :title = "`${modalTitle}`"
                  @close="closeModal"

                > 
                  <template #titleicon>
                    <component :is="currentAction?.icon" />
                  </template>
                  <template #crossicon>
                    <component :is="currentAction?.crossIcon" />
                  </template>
                  <template #body>
                    <component :is="currentAction?.bodyComponent" v-bind="modalProps"></component>
                  </template>
                </ModalWindow>
            </div>
          </div>
        </div> 
    </div>
</template>
  
<script setup>
import { ref, defineProps, computed } from 'vue';
import ModalWindow from './ui/ModalWindow.vue';
  
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
  fields: {
    type: Array,
    required: true,
  },
  pageName: {
    type: String, 
    required: true,
  },
});

const titleFormat = (pageName) => {
  const lowerPageName = pageName.toLowerCase();

  const wordForms = {
    'пользователи': 'пользователя',
    'токены': 'токен',
    'системы': 'систему',
  };

  return wordForms[lowerPageName] || lowerPageName;
  
};

const selectedRow = ref(null);
const isModalOpen = ref(false);
const currentAction = ref(null);


// Форматирование процессов (если нужно)
const formatProcesses = (processes) => {
  if (!processes) return '';
  return Object.entries(processes)
    .filter(([_, value]) => value)
    .map(([key]) => key)
    .join(', ');
};

const handleActionClick = (action, row) => {
  currentAction.value = action; 
  selectedRow.value = row;
  console.log(selectedRow.value)
  isModalOpen.value = true; 
}

const closeModal = () => {
  isModalOpen.value = false;
  currentAction.value = null; 
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

const modalTitle = computed(() => {
  if (currentAction.value?.type === 'edit') {
    return `Редактировать ${titleFormat(props.pageName)}`;
  } else if (currentAction.value?.type === 'delete') {
    return `Удалить ${titleFormat(props.pageName)}`;
  }
  return '';
}); 

// Динамические пропсы для компонента
const modalProps = computed(() => {
  if (currentAction.value?.type === 'edit') {
    return {
      fields: props.fields,
    };
  } else if (currentAction.value?.type === 'delete') {
    return {
      modalTitle: titleFormat(props.pageName), 
      itemName: selectedRow.value.public_token || selectedRow.value.login || 'элемент',
    };
  }
  return {};
});

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