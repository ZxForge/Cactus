<template>
    <div class="user_delete_modal">
        <div class="user_delete_modal_backdrop">
            <div class="user_delete_modal_container">
                <header class="modal_header">
                    <div class="left_block_header">
                        <TrashIcon style="cursor: default;"/>
                        <h2 class="modal_title">Удаление</h2>
                    </div>
                    <div class="left_block_header">
                        <RouterLink to="/users">
                            <button class="close_button"><CrossIcon /></button>
                        </RouterLink>
                    </div>
                </header>
                <div class="main_content">
                    <main class="modal_body">
                        <p>Вы действительно хотиите удалить пользователя Федя Пупкин?
                        </p>
                    </main>
                    <footer class="modal_footer">
                        <button class="delete_button">Удалить</button>
                        <RouterLink to="/users">
                            <button class="cancel_button">Отменить</button>
                        </RouterLink>   
                    </footer>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import CrossIcon from '@/components/icons/CrossIcon.vue';
import TrashIcon from '@/components/icons/TrashIcon.vue';
// Add any needed script logic here
</script>

<style scoped>
.user_delete_modal {
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

.user_delete_modal_backdrop {
    position: absolute;
    width: 100%;
    height: 100%;
    background-color:rgba(22, 44, 67, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
}

.user_delete_modal_container {
    width: 384px;
    height: 158px;
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

.delete_button {
    width: 100px;
    height: 40px;
    padding: 10px 20px;
    font-size: 14px;
    font-weight: 700;
    color: rgba(246, 85, 95, 1);
    background-color: rgba(255, 255, 255, 1);
    border: 1px solid rgba(246, 85, 95, 1);
    border-radius: 8px;
    cursor: pointer;
}

.cancel_button {
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