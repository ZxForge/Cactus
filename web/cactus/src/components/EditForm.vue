<template>
    <div class="edit_form">
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
                type="text"
                :id="field.key"
                :placeholder="field.placeholder || ''"
              />
  
              <!-- Чекбоксы -->
              <div v-else-if="field.type === 'checkbox'" class="checkbox_style">
                <input
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
                    type="checkbox"
                    :id="`${field.key}-${optionKey}`"
                  />
                  <label :for="`${field.key}-${optionKey}`">{{ option }}</label>
                </div>
              </div>
  
              <!-- Выпадающий список -->
              <div v-else-if="field.type === 'select'" class="select_style">
                <select :id="field.key">
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
                type="password"
                :id="field.key"
                :placeholder="field.placeholder || ''"
              />
            </template>
          </div>
        </div>
      </main>
    </div>
</template>
  
<script setup lang="ts">

const props = defineProps({
  fields: {
    type: Array as () => Field[],
    required: true,
  },
});

// Типы для полей
interface Field {
  key: string;
  label: string;
  labelcb?: string;
  type: 'text' | 'checkbox' | 'checkbox-group' | 'select' | 'password';
  placeholder?: string;
  options?: Record<string, string>;
}

</script>
  
<style scoped>
.edit_form {
  width: 100%;
  width: 600px;
  height: 500px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
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

.checkbox_style,
.checkbox_style_group {
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