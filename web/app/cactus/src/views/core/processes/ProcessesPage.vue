<template>
    <div>
        <div class="main_content">
            <HeaderPages :info_header="breadcrumbs" />
        </div>
    </div>
</template>

<script setup lang="ts">
import HeaderPages from '@/components/HeaderPages.vue';

interface InfoHeader {
  url_info: string;
  url: string;
  name_pages?: string;
}

const breadcrumbs: InfoHeader[] = [
    { url_info: 'Главная', url: '/', name_pages: 'Процессы' },
    { url_info: 'Процессы', url: '/processes' },
    { url_info: 'Рассылка писем', url: '/emails' }
];
</script>

<style></style>
