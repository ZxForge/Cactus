<template>
    <teleport to="body">
      <div v-if="isOpen" class="modal_overlay" @click.self="closeModal">
        <div class="modal_container">
          <header class="modal_header">
            <div class="modal_header_left">
              <slot name="icon">
                <!-- Слот для иконки (по умолчанию пустой) -->
              </slot>
              <h2 class="modal_title">{{ title }}</h2>
            </div>
            <button class="modal_close_button" @click="closeModal">
              <CrossIcon />
            </button>
          </header>
          <main class="modal_body">
            <p>{{ message }}</p>
          </main>
          <footer class="modal_footer">
            <button class="modal_button delete_button" @click="onConfirm">
              {{ confirmButtonText }}
            </button>
            <button class="modal_button cancel_button" @click="closeModal">
              {{ cancelButtonText }}
            </button>
          </footer>
        </div>
      </div>
    </teleport>
  </template>
  
  <script setup>
  import { ref, watch } from 'vue';
  import CrossIcon from '../icons/CrossIcon.vue'; // Импортируйте вашу иконку
  
  // Пропсы
  const props = defineProps({
    isOpen: {
      type: Boolean,
      required: true,
    },
    title: {
      type: String,
      default: 'Удаление',
    },
    message: {
      type: String,
      required: true,
    },
    confirmButtonText: {
      type: String,
      default: 'Удалить',
    },
    cancelButtonText: {
      type: String,
      default: 'Отменить',
    },
  });
  
  // Эмиты
  const emit = defineEmits(['close', 'confirm']);
  
  // Закрытие модального окна
  const closeModal = () => {
    emit('close');
  };
  
  // Подтверждение действия
  const onConfirm = () => {
    emit('confirm');
  };
  </script>
  
  <style scoped>
  .modal_overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }
  
  .modal_container {
    background-color: white;
    border-radius: 8px;
    width: 400px;
    padding: 20px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  }
  
  .modal_header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }
  
  .modal_header_left {
    display: flex;
    align-items: center;
    gap: 10px;
  }
  
  .modal_title {
    margin: 0;
    font-size: 1.25rem;
    font-weight: bold;
  }
  
  .modal_close_button {
    background: none;
    border: none;
    cursor: pointer;
    padding: 0;
    display: flex;
    align-items: center;
  }
  
  .modal_body {
    margin-bottom: 20px;
  }
  
  .modal_footer {
    display: flex;
    justify-content: flex-end;
    gap: 10px;
  }
  
  .modal_button {
    padding: 10px 20px;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    font-size: 14px;
    font-weight: 700;
  }
  
  .delete_button {
    border: 1px solid rgba(246, 85, 95, 0.5);
    background-color: white;
    color: rgba(246, 85, 95, 1);
  }
  
  .delete_button:hover {
    color: white;
    background-color: rgba(246, 85, 95, 1);
  }
  
  .cancel_button {
    color: white;
    background-color: rgba(85, 111, 246, 1);
  }
  
  .cancel_button:hover {
    background-color: rgba(85, 111, 246, 0.7);
  }
  </style>