<template>
    <div class="delete_modal">
      <main class="modal_body">
        <p>Вы уверены, что хотите удалить {{ modalTitle }}: <strong>{{ itemName }}</strong>?</p>
      </main>
      <footer class="modal_footer">
        <button class="buttons cancel_button" @click="closeModal">Отмена</button>
        <button class="buttons delete_button" @click="confirmDelete">Удалить</button>
      </footer>
    </div>
</template>
  
<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';

// Пропсы
const props = defineProps({
  modalTitle: {
    type: String,
    required: true,
  },
  itemName: {
    type: String,
    required: true,
  },
});

// Эмиты
const emit = defineEmits(['close', 'delete']);

// Закрытие модального окна
const closeModal = () => {
  emit('close');
};

// Подтверждение удаления
const confirmDelete = () => {
  emit('delete');
};
</script>

<style scoped>
.delete_modal {
  width: 385px;
}

.modal_body {
  font-size: 18px;
  padding-left: 5px;
  margin-bottom: 20px;
}

.modal_body strong {
  font-weight: bold;
}

.modal_footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.buttons{
    border: none;
    padding: 10px 20px;
    border-radius: 8px;
    cursor: pointer;
    font-size: 14px;
    font-weight: 700;
}

.cancel_button {
  color: white;
  background-color: rgba(85, 111, 246, 1);
}

.cancel_button:hover {
  background-color: rgba(85, 111, 246, 0.7);
}

.delete_button {
  color: white;
  background-color: rgba(246, 85, 95, 1);
}

.delete_button:hover {
  color: white;
  background-color: rgba(246, 85, 95, 0.7);
}
</style>