import Vue, { ref } from 'vue'
import { getMenuContent } from '@/api/menuAPI'
import { withAsync } from '@/api/helpers/withAsync'

export const objectTypes = {
  REFERENCE: 1,
  CLASSIFIER: 2
}

export const newRoute = ({
  name,
  routeTo,
  iconName = null,
  childrens = null,
  hasSubMenu = false,
  refType = null,
  isTree = false,
  resourceName = null,
  permissionName = null
}) => {
  return {
    name,
    routeTo,
    iconName,
    childrens,
    isOpen: false,
    hasSubMenu,
    refType,
    isTree,
    resourceName,
    permissionName
  }
}

const state = Vue.observable({
  navs: [
    newRoute({
      name: 'Информационная панель',
      routeTo: { name: 'Analytics' },
      iconName: 'home',
      resourceName: 'analytics'
    }),
    newRoute({
      name: 'Справочники',
      routeTo: { name: 'ReferenceRecordsList' },
      iconName: 'view-grid-detail',
      childrens: [],
      hasSubMenu: true,
      refType: objectTypes.REFERENCE,
      resourceName: 'references'
    }),
    newRoute({
      name: 'Классификаторы',
      routeTo: { name: 'ClassifierTree' },
      iconName: 'box',
      childrens: [],
      hasSubMenu: true,
      refType: objectTypes.CLASSIFIER,
      resourceName: 'references'
    }),
    newRoute({
      name: 'Структура данных',
      routeTo: { name: 'ConstructorList' },
      iconName: 'upload-three',
      resourceName: 'structure'
    }),
    newRoute({
      name: 'Импорт данных',
      routeTo: { name: 'ImportDataList' },
      iconName: 'import-and-export',
      resourceName: 'import-export'
    }),
    newRoute({
      name: 'Интеграция',
      routeTo: { name: null },
      iconName: 'sphere',
      resourceName: 'integration',
      childrens: [
        newRoute({
          name: 'Сценарий',
          routeTo: { name: 'IntegrationScenarioList' },
          iconName: 'import-and-export'
        }),
        newRoute({
          name: 'Маршруты',
          routeTo: { name: 'IntegrationRouteList' },
          iconName: 'import-and-export'
        }),
        newRoute({
          name: 'Процессы',
          routeTo: { name: 'IntegrationProcessesList' },
          iconName: 'import-and-export'
        }),
        newRoute({
          name: 'Журнал',
          routeTo: { name: 'IntegrationLogsList' },
          iconName: 'import-and-export'
        })
      ],
      hasSubMenu: true
    }),
    newRoute({
      name: 'Админстрирование',
      routeTo: { name: null },
      iconName: 'sphere',
      resourceName: 'adminstration',
      childrens: [
        // newRoute({
        //   name: 'Роли',
        //   routeTo: { name: 'RolesList' },
        //   iconName: 'import-and-export'
        // }),
        newRoute({
          name: 'Пользователи',
          routeTo: { name: 'UsersList' },
          iconName: 'import-and-export'
        })
      ],
      hasSubMenu: true
    }),
    newRoute({
      name: 'FAQ',
      routeTo: { name: 'FaqList' },
      iconName: 'file-question',
      resourceName: 'faq'
    })
  ]
})

export const navsComputed = {
  navs: {
    get() {
      return state.navs
    }
  }
}

export const toggleMenu = async (idx, refType) => {
  if (!state.navs[idx].isOpen && state.navs[idx].refType) {
    state.navs[idx].childrens = await fetchRefObjects(refType)
    state.navs[idx].isOpen = !state.navs[idx].isOpen
  } else {
    state.navs[idx].isOpen = !state.navs[idx].isOpen
  }
}

export const updateChildrens = async (refType) => {
  for (let idx in state.navs) {
    if (state.navs[idx].refType == refType) {
      state.navs[idx].childrens = await fetchRefObjects(refType)
    }
  }
}

async function fetchRefObjects(objectId) {
  const { response, error } = await withAsync(getMenuContent, {
    table_type_id: objectId
  })
  if (error) {
    return
  }
  return response.data.map((el) =>
    newRoute({
      name: el.table_show,
      routeTo: {
        name:
          objectId == TableType.REFERENCE.id || !el.is_tree
            ? TableType.REFERENCE.pageName
            : TableType.CLASSIFIER.pageName,
        params: {
          reference_name: el.table_name
        },
        query: {
          table_id: el.table_id,
          is_classifier: objectId == TableType.CLASSIFIER.id,
          table_show: el.table_show
        }
      },
      isTree: el.is_tree
    })
  )
}

const TableType = {
  REFERENCE: {
    pageName: 'ReferenceRecordsList',
    id: 1
  },
  CLASSIFIER: {
    pageName: 'ClassifierTree',
    id: 2
  }
}
