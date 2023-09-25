import AuthLayout from '@/layouts/auth/AuthLayout.vue'

const authRoutes = [
  {
    path: '/auth',
    name: 'AuthLayout',
    redirect: '/auth/login',
    component: AuthLayout,
    meta: {
      requiredCheckAuth: true
    },
    children: [
      {
        path: 'login',
        name: 'Login',
        meta: {
          title: 'Вход'
        },
        component: () =>
          import(/* webpackChunkName: "Login" */ `@/views/auth/Login.vue`)
      },
    ]
  }
]

export default authRoutes
