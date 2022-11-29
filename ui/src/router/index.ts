import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import DashboardView from '@/views/dashboard/DashboardView.vue';
import NodeList from '@/views/nodes/NodeList.vue';
import LocalVolumeIndex from '@/views/local-volumes/LocalVolumeIndex.vue';
import SettingsView from '@/views/settings/SettingsView.vue';
import RouterContent from '@/views/RouterContent.vue';
import NodeDetailIndex from '@/views/nodes/NodeDetailIndex.vue';
import NodeDetailLocalDisks from '@/views/nodes/NodeDetailLocalDisks.vue';
import NodeDetailMigrates from '@/views/nodes/NodeDetailMigrates.vue';
import PoolList from '@/views/pools/PoolList.vue';
import PoolDetail from '@/views/pools/PoolDetail.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: DashboardView,
  },
  {
    path: '/nodes',
    name: 'Node',
    component: RouterContent,
    redirect: { name: 'NodeList' },
    children: [
      {
        path: '',
        name: 'NodeList',
        component: NodeList,
      },
      {
        path: ':name',
        name: 'NodeDetailIndex',
        component: NodeDetailIndex,
        redirect: { name: 'NodeDetailLocalDisks' },
        children: [
          {
            path: 'disks',
            name: 'NodeDetailLocalDisks',
            component: NodeDetailLocalDisks,
          },
          {
            path: 'migrates',
            name: 'NodeDetailMigrates',
            component: NodeDetailMigrates,
          },
        ],
      },
    ],
  },
  {
    path: '/pools',
    name: 'Pool',
    component: RouterContent,
    redirect: { name: 'PoolList' },
    children: [
      {
        path: '',
        name: 'PoolList',
        component: PoolList,
      },
      {
        path: ':name',
        name: 'PoolDetail',
        component: PoolDetail,
      },
    ],
  },
  {
    path: '/local-volumes',
    name: 'LocaleVolume',
    component: LocalVolumeIndex,
  },
  {
    path: '/settings',
    name: 'Settings',
    component: SettingsView,
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: { name: 'Dashboard' },
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
