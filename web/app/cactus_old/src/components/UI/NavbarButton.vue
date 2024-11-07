<template>
    <div class="navbar_buttons" :class="{ open: open }">
        <RouterLink :class="{active: isExactActive(to)}" class="button" :to="to">
            <div class="content">
                <slot name="icon" class="icon"></slot>
                <slot name="title" class="title"> Кнопка </slot>
            </div>
            <div
                v-if="$slots.subbutton"
                @click="open = !open"
                class="roll_button"
            >
                <ArrowIcon class="roll_icon" :class="{ open: open }" />
            </div>
        </RouterLink>
        <div v-if="$slots.subbutton && open" class="subbutton_list">
            <slot name="subbutton" :to="to"></slot>
        </div>
    </div>
</template>

<script setup>
import ArrowIcon from "../Icon/ArrowIconR.vue";

import { ref } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();

const props = defineProps({
    to: {
        type: String,
        required: true,
    },
});

function isExactActive(url) {
    return route.path === url || (url.split('/').filter(e=>e!=='').length !== 0 && route.path.startsWith(url));
}

const open = ref(false);
</script>

<style lang="less" scope>
.navbar_buttons {
    width: 190px;
    color: #fff;
    font-size: 12px;
    font-family: "Inter";
    margin: 10px 0;
    border-radius: 8px;
    cursor: pointer;

    .button {
        display: flex;
        justify-content: space-between;
        text-align: center;
        width: 100%;

        padding: 7px 10px;
        border-radius: 8px;
        cursor: pointer;

        &:hover {
            background-color: rgba(85, 111, 246, 1);
        }

        &.active {
            background-color: rgba(85, 111, 246, 1);
        }

        .content {
            display: flex;
            align-items: center;
            gap: 10px;

            .icon {
                display: flex;
                justify-content: center;
                align-items: center;
                width: 20px;
                height: 20px;
            }
        }
    }

    &.open {
        background-color: rgba(85, 111, 246, 0.1);
    }

    .roll_button .roll_icon {
        transition: 0.2s;
        transform: rotate(-90deg);
        display: flex;
        justify-items: center;
        &.open {
            transform: rotate(90deg);
        }
    }

    .subbutton_list {
        display: flex;
        flex-direction: column;
        border-radius: 10px;
        transition: 0.5s;

        .subbutton {
            padding: 7px 10px;
            color: rgba(146, 147, 171, 1);
            display: flex;
            align-items: center;
            gap: 10px;

            &:hover,
            &:focus-visible {
                color: #fff;
                .icon {
                    color: #fff;
                }
            }

            &.active {
                color: #fff;
                .icon {
                    color: #fff;
                }
            }

            .icon {
                width: 20px;
                color: rgba(146, 147, 171, 1);
            }

            
        }
    }
}
</style>