<template>
    <div class="user_edit_modal">
        <div class="user_edit_modal_backdrop">
            <div class="user_edit_modal_container">
                <header class="modal_header">
                    <div class="left_block_header">
                        <EditPencil style="cursor: default;"/>
                        <h2 class="modal_title">Редактирование - Albus Dumbledore</h2>
                    </div>
                    <div class="left_block_header">
                        <RouterLink to="/users">
                            <button class="close_button"><CrossIcon /></button>
                        </RouterLink>
                    </div>
                </header>
                <div class="main_content">
                    <main class="modal_body">
                        <div class="modal_body_field">
                            <label for="login" class="field_label">Логин</label>
                            <input id="login" type="text" class="field_input" placeholder="Введите логин"/>
                        </div>
                        <div class="modal_body_field">
                            <label for="email" class="field_label">Почта</label>
                            <input id="email" type="email" class="field_input" placeholder="Введите почту"/>
                        </div>
                        <div class="modal_body_field">
                            <label for="role" class="field_label">Роль</label>
                            <select id="role" class="field_input">
                                <option value="" disabled selected>Выберите роль</option>
                                <option value="admin">Администратор</option>
                                <option value="guest">Гость</option>
                            </select>
                        </div>
                        <div class="modal_body_field">
                            <label for="status" class="field_label">Статус</label>
                            <label><input id="status" type="checkbox" class="field_checkbox"/>Активен</label>
                        </div>
                        <div class="modal_body_field">
                            <label for="password" class="field_label">Новый пароль</label>
                            <input id="password" type="password" class="field_input" placeholder="Введите пароль">
                        </div>
                        <div class="modal_body_field">
                            <div class="field_label"></div>
                            <label><input id="change_on_login" type="checkbox" class="field_checkbox"/> Сменить при входе</label> 
                        </div>
                    </main>
                    <footer class="modal_footer">
                        <RouterLink to="/users">
                            <button class="cancel_button">Отмена</button>
                        </RouterLink>
                        <button class="save_button">Изменить</button>
                    </footer>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import CrossIcon from '@/components/icons/CrossIcon.vue';
import EditPencil from '@/components/icons/EditPencil.vue';
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
    background-color:rgba(22, 44, 67, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
}

.user_edit_modal_container {
    width: 605px;
    height: 590px;
    background-color: rgba(255, 255, 255, 1);
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    display: flex;
    flex-direction: column;
    padding: 16px;
}

.modal_header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.left_block_header{
    display: flex;
    align-items: center;
    gap: 16px;
}

.modal_title {
    font-size: 20px;
    font-weight: 400;
    line-height: 24px;
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

.main_content{
    width: 100%;
    padding: 0;
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.modal_body {
    display: flex;
    flex-direction: column;
    padding-top: 16px;
}

.modal_body_field{
    display: flex;
    align-items: center;
    gap: 10px;
    margin-bottom: 16px;
}

.edit_form {
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.field_label {
    width: 150px;
    height: 30px; 
    margin-bottom: 4px;
    font-size: 14px;
    font-weight: 500;
    color: #555;
    display: flex;
    align-items: center;
}

.field_input {
    flex-grow: 1;
    width: 395px;
    height: 30px;
    padding: 5px 10px;
    font-size: 14px;
    border: 1px solid rgba(146, 147, 171, 0.5);
    border-radius: 4px;
}

.field_checkbox {
    margin-right: 8px;
}

.modal_footer {
    display: flex;
    justify-content: flex-end;
    gap: 10px;
}

.cancel_button {
    width: 100px;
    height: 40px;
    padding: 10px 20px;
    font-size: 14px;
    font-weight: 700;
    color: rgba(146, 147, 171, 1);
    background-color: rgba(255, 255, 255, 1);
    border: 1px solid rgba(146, 147, 171, 1);
    border-radius: 8px;
    cursor: pointer;
}

.save_button {
    width: 110px;
    height: 40px;
    padding: 10px 20px;
    font-size: 14px;
    font-weight: 700;
    color: rgba(255, 255, 255, 1);
    background-color: rgba(85, 111, 246, 1);
    border: none;
    border-radius: 8px;
    cursor: pointer;
}
</style>
  