<template>
  <Teleport to="body">
    <div class="modal">
      <div class="modal_overlay" @click.self="closeModal"></div>
      <div class="modal_content">
        <header class="modal_header">
          <div class="left_header">
            <slot name="titleicon">
              
            </slot>
            <slot name="titletext" class="modal_text">{{ title }}</slot>
          </div>
          <button class="close_button" @click="closeModal">
            <slot name="crossicon"></slot>
          </button>
        </header>
        <main class="modal_body" >
          <slot name="body"></slot>
        </main>
        <footer class="modal_footer">
          <slot name="footer">

          </slot>
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

.modal_footer {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
}

:slotted(.buttons){
  border: none;
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 700;
}

:slotted(.cancel_button) {
  color: white;
  background-color: rgba(85, 111, 246, 1);
}

:slotted(.cancel_button:hover) {
  background-color: rgba(85, 111, 246, 0.7);
}

:slotted(.save_button) {
  color: white;
  background-color: rgb(92, 204, 82); 
}

:slotted(.save_button:hover) {
  color: white;
  background-color: rgba(92, 204, 82, 0.7);
}

:slotted(.delete_button) {
  color: white;
  background-color: rgba(246, 85, 95, 1);
}

:slotted(.delete_button:hover) {
  color: white;
  background-color: rgba(246, 85, 95, 0.7);
}
</style>