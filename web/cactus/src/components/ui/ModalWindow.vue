<template>
  <Teleport to="body">
    <div v-if="isOpen" class="modal">
      <div class="modal_overlay" @click.self="closeModal"></div>
      <div class="modal_content">
        <header class="modal_header">
          <div class="left_header">
            <slot name="titleicon"></slot>
            <p class="modal_text">{{ props.title }}</p>
          </div>
          <button class="close_button" @click="closeModal">
            <slot name="crossicon" class="crossicon"></slot>
          </button>
        </header>
        <main class="modal_body">
          <!-- Слот для основного содержимого модального окна -->
          <slot name="body"></slot>
        </main>
        <footer class="modal_footer">
          <!-- Слот для кнопок в футере модального окна -->
          <slot name="footer"></slot>
        </footer>
      </div>
    </div>
  </Teleport>
</template>
  
<script setup>
import { defineProps } from 'vue';

const props = defineProps({
    title: {
        type: String,
        required: true,
    },
    isOpen: {
      type: Boolean,
      required: true,
    },
});
const emit = defineEmits(['close']);

const closeModal = () => {
  emit('close');
}
</script>
  
<style scoped>
*{
  margin: 0;
  padding: 0;
}
.modal_text{
  font-size: 18px;
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.5);
}

.modal_overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.modal_content {
  background: white;
  padding: 20px;
  border-radius: 8px;
  position: relative;
  z-index: 1;
}

.modal_header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.left_header{
  font-size: 18px;
  display: flex;
  align-items: center;
  gap: 8px;
}

.close_button{
  border: none;
}
</style>