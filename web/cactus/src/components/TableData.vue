<template>
    <div class="table">
        
        <div class="table_header">
          <div v-for="header in headers" :key="header.key" class="table_header_cell">
            <span class="table_text table_header_text">{{ header.label }}</span>
          </div>
        </div>
  
        <div class="table_row" v-for="(row, rowIndex) in rows" :key="rowIndex">
          <div v-for="header in headers" :key="header.key" class="table_cell">

            <span v-if="header.type === 'checkbox'" class="table_text center_align">
              <input type="checkbox" v-model="row[header.key]" />
            </span>

            <span v-else-if="header.type === 'text'" class="table_text">
              {{ row[header.key] }}
            </span>

            <span v-else-if="header.type === 'status'" class="table_text">
              {{ row[header.key] }}
            </span>

            <span v-else-if="header.type === 'processes'" class="table_text">
              {{ formatProcesses(row[header.key]) }}
            </span>

            <div v-else-if="header.type === 'functions'" class="functions">
                <button
                    v-for="(action, actionIndex) in header.actions"
                    :key="actionIndex"
                    @click="handleActionClick(action, row)"
                    class="buttons"
                >   
                    <component :is="action.icon" />
                </button>
            </div>
          </div>
        </div> 
    </div>
</template>
  
<script setup>
import { ref, defineProps, computed } from 'vue';
  
const props = defineProps({
  headers: {
    type: Array,
    required: true,
  },
  rows: {
    type: Array,
    required: true,
  },
  columnRows: {
    type: Number, 
    required: true, 
  },
});

const emit = defineEmits(['action'])

const formatProcesses = (processes) => {
  if (!processes) return '';
  return Object.entries(processes)
    .filter(([_, value]) => value)
    .map(([key]) => key)
    .join(', ');
};

const handleActionClick = (action, row) => {
  emit('action', {
    actionType: action.type, 
    selectedRow: row, 
    isOpen: true, 
  })
}

</script>
  
<style scoped>
  .table {
    background-color: rgba(255, 255, 255, 1);
    display: grid;
    grid-template-columns: repeat(v-bind(columnRows), minmax(0, 1fr));
    border: 1px solid rgba(146, 147, 171, 0.5);
    border-radius: 10px;
    padding: 10px;
  }
  
  .table_header,
  .table_row {
    display: contents;
  }

  .table_cell:first-child{
    padding-left: 35px;
  }

  .table_header,
  .table_row {
    display: contents;
  }
  
  .table_header_cell,
  .table_cell {
    padding: 15px;
    border-bottom: 1px solid #e5e5e5;
    text-align: left;
  }
  
  .table_row:last-child .table_cell {
    border-bottom: none;
  }
  
  .table_text {
    width: 37px;
    height: 17px;
    font-family: 'Inter', sans-serif;
    font-weight: 400;
    font-size: 16px;
    line-height: 16.94px;
  }
  
  .table_header_text {
    font-weight: bold;
  }
  
  .functions {
    display: flex;
    justify-content: flex-start;
    gap: 30px;
  }
  
  .functions div {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 25px;
    width: 25px;
    border: none;
  }
  
  .buttons {
    background-color: white;
    border: none;
    width: 20px;
    height: 20px;
    cursor: pointer;
  }
</style>