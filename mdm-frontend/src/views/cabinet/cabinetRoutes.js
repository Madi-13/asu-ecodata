import CabinetLayout from '@/layouts/cabinet/CabinetLayout.vue'

const cabinetRoutes = [
  {
    path: '/cabinet',
    name: 'CabinetLayout',
    component: CabinetLayout,
    redirect: '/cabinet/analytics',
    meta: {
      requiresAuth: true
    },
    children: [
      {
        path: 'analytics',
        name: 'Analytics',
        component: () =>
          import(
            /* webpackChunkName: "schedule" */ `@/views/cabinet/analytics/AnalyticsList.vue`
          ),
        meta: {
          title: 'Информационная панель',
          resourceName: 'analytics'
        }
      },
      {
        path: 'constructor',
        name: 'ConstructorList',
        component: () =>
          import(
            /* webpackChunkName: "constructor" */ `@/views/cabinet/constructor/ConstructorList.vue`
          ),
        meta: {
          title: 'Структура данных',
          resourceName: 'structure'
        }
      },
      {
        path: 'constructor/create',
        name: 'ConstructorCreatePage',
        component: () =>
          import(
            /* webpackChunkName: "constructor" */ `@/views/cabinet/constructor/ConstructorCreate.vue`
          ),
        meta: {
          title: null,
          resourceName: 'structure'
        }
      },
      {
        path: 'constructor/:id',
        name: 'ConstructorEditPage',
        component: () =>
          import(
            /* webpackChunkName: "constructor" */ `@/views/cabinet/constructor/ConstructorEdit.vue`
          ),
        meta: {
          title: null,
          resourceName: 'structure'
        }
      },
      {
        path: 'reference/records/:reference_name',
        name: 'ReferenceRecordsList',
        component: () =>
          import(
            /* webpackChunkName: "referenceRecordList" */ `@/views/cabinet/references/ReferenceRecordsList.vue`
          ),
        meta: {
          title: null,
          resourceName: 'references'
        }
      },
      {
        path: 'reference/tree/:reference_name',
        name: 'ClassifierTree',
        component: () =>
          import(
            /* webpackChunkName: "classifierTree" */ `@/views/cabinet/classifier/ClassifierTree.vue`
          ),
        meta: {
          title: 'Классификаторы',
          resourceName: 'references'
        }
      },
      {
        path: 'reference/classifierlist/:reference_name',
        name: 'ClassifierList',
        component: () =>
          import(
            /* webpackChunkName: "classifierTree" */ `@/views/cabinet/references/ReferenceRecordsList.vue`
          ),
        meta: {
          title: 'Классификаторы',
          resourceName: 'references'
        }
      },
      {
        path: 'reference/records/:reference_name/create',
        name: 'ReferenceRecordCreate',
        component: () =>
          import(
            /* webpackChunkName: "referenceRecordCreate" */ `@/views/cabinet/references/ReferenceRecordCreate.vue`
          ),
        meta: {
          title: null,
          resourceName: 'references'
        }
      },
      {
        path: 'reference/records/:reference_name/edit/:row_id',
        name: 'ReferenceRecordEdit',
        component: () =>
          import(
            /* webpackChunkName: "referenceRecordEdit" */ `@/views/cabinet/references/ReferenceRecordEdit.vue`
          ),
        meta: {
          title: null,
          resourceName: 'references'
        }
      },
      {
        path: 'reference/records/:reference_name/clone/:row_id',
        name: 'ReferenceRecordClone',
        component: () =>
          import(
            /* webpackChunkName: "referenceRecordEdit" */ `@/views/cabinet/references/ReferenceRecordClone.vue`
          ),
        meta: {
          title: null,
          resourceName: 'references'
        }
      },
      {
        path: 'import',
        name: 'ImportDataList',
        component: () =>
          import(
            /* webpackChunkName: "importDataList" */ `@/views/cabinet/import/ImportDataList.vue`
          ),
        meta: {
          title: 'Импорт данных',
          resourceName: 'import-export'
        }
      },
      {
        path: 'integration/scenario',
        name: 'IntegrationScenarioList',
        component: () =>
          import(
            /* webpackChunkName: "IntegrationScenarioList" */ `@/views/cabinet/integration/scenario/IntegrationScenarioList.vue`
          ),
        meta: {
          title: null,
          resourceName: 'integration'
        }
      },
      {
        path: 'integration/routes',
        name: 'IntegrationRouteList',
        component: () =>
          import(
            /* webpackChunkName: "IntegrationRouteList" */ `@/views/cabinet/integration/route/IntegrationRouteList.vue`
          ),
        meta: {
          title: null,
          resourceName: 'integration'
        }
      },
      {
        path: 'integration/processes',
        name: 'IntegrationProcessesList',
        component: () =>
          import(
            /* webpackChunkName: "IntegrationProcessesList" */ `@/views/cabinet/integration/processes/IntegrationProcessesList.vue`
          ),
        meta: {
          title: null,
          resourceName: 'integration'
        }
      },
      {
        path: 'integration/logs',
        name: 'IntegrationLogsList',
        component: () =>
          import(
            /* webpackChunkName: "IntegrationLogsList" */ `@/views/cabinet/integration/logs/IntegrationLogsList.vue`
          ),
        meta: {
          title: null,
          resourceName: 'integration'
        }
      },
      {
        path: 'faq',
        name: 'FaqList',
        component: () =>
          import(
            /* webpackChunkName: "faqList" */ `@/views/cabinet/faq/FaqList.vue`
          ),
        meta: {
          title: null,
          resourceName: 'faq'
        }
      },
      {
        path: 'roles',
        name: 'RolesList',
        component: () =>
          import(
            /* webpackChunkName: "rolesList" */ `@/views/cabinet/adminstration/roles/RolesList.vue`
          ),
        meta: {
          title: 'Роли',
          resourceName: 'adminstration'
        }
      },
      {
        path: 'users',
        name: 'UsersList',
        component: () =>
          import(
            /* webpackChunkName: "usersList" */ `@/views/cabinet/adminstration/users/UsersList.vue`
          ),
        meta: {
          title: 'Пользователи',
          resourceName: 'adminstration'
        }
      }
    ]
  }
]

export default cabinetRoutes
