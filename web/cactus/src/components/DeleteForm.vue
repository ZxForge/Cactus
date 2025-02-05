<template>
    <div class="delete_modal">
      <main class="modal_body">
        <p> {{ modalTitle }}: <strong>{{ itemName }}</strong>?</p>
      </main>
    </div>
</template>
  
<script setup lang="ts">
import { defineProps } from 'vue';

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
</style>