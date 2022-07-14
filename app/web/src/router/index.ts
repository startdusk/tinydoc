import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router"

const DetailView = () => import("../views/detail.vue")
const HomeView = () => import("../views/home.vue")

const routes: RouteRecordRaw[] = [
	{
		path: '/',
		name: 'home',
		component: HomeView,
	},
	{
		path: '/details/:id',
		name: 'details',
		component: DetailView
	},
	{
		path: '/:catchAll(.*)*',
		redirect: { name: 'home' },
	},
]

const router = createRouter({
	history: createWebHashHistory(),
	routes
});

export default router;


