import { createWebHistory, createRouter } from 'vue-router'
import Main from './pages/Main.vue';

const routes = [
    {path: '/', component: Main}
]

// Динамическое получение всех файлов с расширением vue в папке pages
let requireComponent = import.meta.glob(['./pages/**/*.vue', '!./pages/Main.vue'])

Object.keys(requireComponent).forEach(fileName => {
    const componentConfig = requireComponent[fileName];

    const componentName = fileName
        .split('/')
        .map(path => {
            // PascalCase to kebab-case
            return path.replace(/[A-Z]/g, (match, offset)=>{
                const prepend = offset <= 0? "": "-"
                return prepend + match.toLowerCase()
            });
        })
        .slice(2)
        .join('/')
        .replace(/\.\w+$/g, '/');

    const component =  componentConfig.default || componentConfig;
    routes.push({path: `/${componentName}`, component});
});

export default createRouter({
  history: createWebHistory(),
  routes,
})