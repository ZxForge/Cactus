<template>
    <div class="user_edit_modal">
        <div class="user_edit_modal_backdrop">
            <div class="user_edit_modal_container">
                <header class="modal_header">
                    <h2 class="modal_title">Редактирование - Albus Dumbledore</h2>
                    <button class="close_button"><CrossIcon /></button>
                </header>
                <main class="modal_body">
                    <form class="edit_form">
                        <div class="form_field">
                            <label for="login" class="field_label">Login</label>
                            <input id="login" type="text" class="field_input" value="Albus Dumbledore" />
                        </div>
                        <div class="form_field">
                            <label for="email" class="field_label">Email</label>
                            <input id="email" type="email" class="field_input" value="Albus@yandex.ru" />
                        </div>
                        <div class="form_field">
                            <label for="role" class="field_label">Роль</label>
                            <select id="role" class="field_input">
                                <option value="admin" selected>Администратор</option>
                                <option value="guest">Гость</option>
                            </select>
                        </div>
                        <div class="form_field">
                            <label for="status" class="field_label">Статус</label>
                            <input id="status" type="checkbox" class="field_checkbox" checked /> Активен
                        </div>
                        <div class="form_field">
                            <label for="password" class="field_label">Новый пароль</label>
                            <input id="password" type="password" class="field_input" value="password123" />
                        </div>
                        <div class="form_field">
                            <input id="change_on_login" type="checkbox" class="field_checkbox" checked /> Сменить при входе
                        </div>
                    </form>
                </main>
                <footer class="modal_footer">
                    <button class="cancel_button">Отмена</button>
                    <button class="save_button">Изменить</button>
                </footer>
            </div>
        </div>
    </div>
</template>

<script setup>
import CrossIcon from '@/components/icons/CrossIcon.vue';
// Add any needed script logic here
</script>

<style scoped>
.user_edit_modal {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 1000;
}

.user_edit_modal_backdrop {
    position: absolute;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
}

.user_edit_modal_container {
    position: relative;
    top: 25vh;
    left: 35vw;
    width: 605px;
    height: 590px;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    display: flex;
    flex-direction: column;
}

.modal_header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 16px;
    border-bottom: 1px solid #ddd;
}

.modal_title {
    font-size: 18px;
    font-weight: bold;
    margin: 0;
}

.close_button {
    background: none;
    border: none;
    font-size: 24px;
    cursor: pointer;
    line-height: 1;
    color: #aaa;
}

.modal_body {
    padding: 16px;
}

.edit_form {
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.form_field {
    display: flex;
    flex-direction: column;
}

.field_label {
    margin-bottom: 4px;
    font-size: 14px;
    font-weight: 500;
    color: #555;
}

.field_input {
    padding: 8px;
    font-size: 14px;
    border: 1px solid #ddd;
    border-radius: 4px;
}

.field_checkbox {
    margin-right: 8px;
}

.modal_footer {
    display: flex;
    justify-content: flex-end;
    gap: 8px;
    padding: 16px;
    background-color: #f5f5f5;
    border-top: 1px solid #ddd;
}

.cancel_button {
    padding: 8px 16px;
    font-size: 14px;
    color: #555;
    background-color: #f5f5f5;
    border: 1px solid #ddd;
    border-radius: 4px;
    cursor: pointer;
}

.save_button {
    padding: 8px 16px;
    font-size: 14px;
    color: #fff;
    background-color: #007bff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}
</style>
  