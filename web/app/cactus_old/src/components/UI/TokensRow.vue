<script setup>
    import Trash from "../Icon/Trash.vue";
    import EditPencil from "../Icon/EditPencil.vue";
    import CheckboxStateOne from "./CheckboxStateOne.vue";
    import CheckboxStateTwo from "./CheckboxStateTwo.vue";
    import TokensEdit from "../../pages/TokensEdit.vue";
    import TokensDelete from "../../pages/TokensDelete.vue";
    import { ref } from 'vue';

    const props = defineProps({
        name: String,
        publicKey: String,
        status: String, 
        description: String, 
        process: String, 
    })

    const showFirstCheckbox = ref(true);
    const isVisibleEdit = ref(false)
    const isVisibleDelete = ref(false)

    const toggleCheckboxes = function() {
      showFirstCheckbox.value = !showFirstCheckbox.value;
    };

    const isVisEdit = function(){
        isVisibleEdit.value = !isVisibleEdit.value;
    }
    const isVisTrash = function(){
        isVisibleDelete.value = !isVisibleDelete.value;
    }

</script>

<template>
    <div id="row">
        <TokensEdit v-if="isVisibleEdit"/>
        <TokensDelete v-if="isVisibleDelete"/>
        <div class="selected" @click="toggleCheckboxes">
            <CheckboxStateOne v-if="showFirstCheckbox"/>
            <CheckboxStateTwo v-else/>
        </div>
        <div class="name" @click="toggleCheckboxes">
            <span>{{name}}</span>
        </div>
        <div class="public-key">
            <span>{{publicKey}}</span>
        </div>
        <div class="status">
            <span>{{status}}</span>
        </div>
        <div class="description">
            <span>{{description}}</span>
        </div>
        <div class="process">
            <span>{{process}}</span>
        </div>
        <div class="action">
            <Trash  @click="isVisTrash"/>
            <EditPencil @click="isVisEdit"/>
        </div>
    </div>
</template>



<style scoped>
@import "../../../src/assets/css/TokensRow.css"
</style>
