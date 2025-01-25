<template>
    <teleport to="body">
      <div v-if="isOpen" class="edit_modal_overlay" @click.self="closeModal">
        <div class="edit_modal_container">
          <header class="modal_header">
            <div class="left_block_header">
              <slot name="icon">
                <!-- Слот для иконки -->
              </slot>
              <h2 class="modal_title">{{ title }}</h2>
            </div>
            <button class="close_button" @click="closeModal">
              <CrossIcon />
            </button>
          </header>
          <main class="modal_body">
            <div class="content_block">
              <div class="left_block">
                <label v-for="field in fields" :key="field.key" :for="field.key">
                  {{ field.label }}
                </label>
              </div>
              <div class="right_block">
                <template v-for="field in fields" :key="field.key">
                    <!-- Текстовые поля -->
                    <input
                        v-if="field.type === 'text'"
                        v-model="formData[field.key]"
                        type="text"
                        :id="field.key"
                        :placeholder="field.placeholder || ''"
                    />
  
                    <!-- Чекбоксы -->
                    <div v-else-if="field.type === 'checkbox'" class="checkbox_style">
                      <input
                        v-model="formData[field.key]"
                        type="checkbox"
                        :id="field.key"
                      />
                      <label :for="field.key">{{ field.labelcb }}</label>
                    </div>
                  
                    <!-- Группа чекбоксов -->
                    <div
                      v-else-if="field.type === 'checkbox-group'"
                      class="checkbox_container"
                    >
                      <div
                        v-for="(option, optionKey) in field.options"
                        :key="optionKey"
                        class="checkbox_style_group"
                      >
                        <input
                          v-model="formData[field.key][optionKey]"
                          type="checkbox"
                          :id="`${field.key}-${optionKey}`"
                        />
                        <label :for="`${field.key}-${optionKey}`">{{ option }}</label>
                      </div>
                    </div>

                    <!-- Выпадающий список -->
                    <div v-else-if="field.type === 'select'" class="select_style">
                      <select
                        v-model="formData[field.key]"
                        :id="field.key"
                      >
                        <option selected disabled :value="undefined">Выберите роль</option>
                        <option
                          v-for="(optionValue, optionKey) in field.options"
                          :key="optionKey"
                          :value="optionKey"
                        >
                          {{ optionValue }}
                        </option>
                        </select>
                    </div>

                    <!-- Пароль -->
                    <input
                        v-if="field.type === 'password'"
                        v-model="formData[field.key]"
                        type="password"
                        :id="field.key"
                        :placeholder="field.placeholder || ''"
                    />
                </template>
              </div>
            </div>
          </main>
          <footer class="modal_footer">
            <button class="cancel_button" @click="closeModal">Отмена</button>
            <button class="save_button" @click="saveChanges">Сохранить</button>
          </footer>
        </div>
      </div>
    </teleport>
</template>
  
<script setup lang="ts">
import { ref, watch } from 'vue';
import CrossIcon from '@/components/icons/CrossIcon.vue';

// Пропсы
const props = defineProps({
  isOpen: {
    type: Boolean,
    required: true,
  },
  title: {
    type: String,
    default: 'Редактирование',
  },
  fields: {
    type: Array as () => Field[],
    required: true,
  },
  initialData: {
    type: Object,
    default: () => ({}),
  },
});

// Типы для полей
interface Field {
  key: string;
  label: string;
  labelcb?: string,
  type: 'text' | 'checkbox' | 'checkbox-group' | 'select' | 'password';
  placeholder?: string;
  options?: Record<string, string>; // Для checkbox-group
}

// Эмиты
const emit = defineEmits(['close', 'save']);

// Данные формы
const formData = ref<Record<string, any>>({});

// Инициализация формы
watch(
  () => props.initialData,
  (newData) => {
    formData.value = { ...newData };
  },
  { immediate: true, deep: true }
);

// Закрытие модального окна
const closeModal = () => {
  emit('close');
};

// Сохранение изменений
const saveChanges = () => {
  emit('save', formData.value);
};
</script>

<style scoped>
.edit_modal_overlay {
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

.edit_modal_container {
  width: 600px;
  background-color: #fff;
  border-radius: 8px;
  padding: 16px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.modal_header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.left_block_header {
  display: flex;
  align-items: center;
  gap: 8px;
}

.modal_title {
  font-size: 20px;
  font-weight: 500;
  margin: 0;
}

.close_button {
  background: none;
  border: none;
  cursor: pointer;
  font-size: 24px;
  color: #aaa;
}

.modal_body {
  margin-bottom: 16px;
}

.content_block {
  display: grid;
  grid-template-columns: 150px 1fr;
  gap: 16px;
}

.left_block {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.left_block label {
  height: 40px;
  display: flex;
  align-items: center;
  font-size: 16px;
}

.right_block {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.right_block input[type="text"], 
.right_block input[type="password"] {
  height: 40px;
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 16px;
}

.checkbox_style, .checkbox_style_group {
  display: flex;
  align-items: center;
  height: 40px;
  gap: 8px;
}
.checkbox_style_group {
  height: 20px;
  padding-top: 20px;
}


.checkbox_container {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.modal_footer {
  margin-top: 100px;
  display: flex;
  justify-content: flex-end;
  gap: 8px;
}

.cancel_button {
  padding: 8px 16px;
  font-size: 14px;
  color: #9293ab;
  background-color: #fff;
  border: 1px solid #9293ab;
  border-radius: 4px;
  cursor: pointer;
}

.save_button {
  padding: 8px 16px;
  font-size: 14px;
  color: #fff;
  background-color: #556ff6;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.select_style {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.select_style select {
    height: 40px;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 16px;
    background-color: #fff;
    cursor: pointer;
}

</style>