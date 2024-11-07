<template>
  <div class="header_main">
      <div class="url_info">
      <template v-for="(info_hed, index) in info_header" :key="info_hed.url_info">
          <a :href="info_hed.url">{{ info_hed.url_info }}</a>
          <ArrowIcon v-if="index < info_header.length - 1" class="arrow" />
      </template>
      </div>
      <div class="name_pages" v-for="info_hed in info_header" :key="info_hed.name_pages">
      <h2>{{ info_hed.name_pages }}</h2>
      </div>
  </div>
</template>

<script>
import ArrowIcon from './Icon/ArrowIconR.vue';
export default {
    components: {
        ArrowIcon
    },
    props: {
        info_header: {
            type: Array,
            required: true
        }
    }
}
</script>

<style scoped>
.url_info {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-left: 10px;
}

.url_info a {
  color:#9293AB; 
  font-size: 12px;
  text-decoration: none; 
  margin: 10px 0 10px 0; 
}

.url_info a:hover {
  text-decoration: underline;
}

.name_pages h2 {
    margin-left: 10px;
    font-size: 32px;
}

.arrow {
  width: 10px;
  height: 10px;
  fill: currentColor;
}
</style>  