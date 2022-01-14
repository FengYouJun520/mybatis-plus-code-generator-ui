import SvgIcon from '@/components/SvgIcon/index.vue'
import { SelectOption } from 'naive-ui'
import { h } from 'vue'

export interface Table {
  name: string
  comment: string
}

export interface TablesOptions {
  label: string
  value: string
  comment: string
}

export const databaseTypeOptions: SelectOption[] = [
  {
    label: () => {
      return h(
        'span',
        {
          style: {
            display: 'flex',
            alignItems: 'center',
            columnGap: '10px',
          },
        },
        [
          h(SvgIcon, {
            name: 'mysql',
            style: {
              width: 24,
              height: 24,
            },
          }),
          h('span', null, 'mysql'),
        ]
      )
    },
    value: 'mysql',
  },
  {
    label: () => {
      return h(
        'span',
        {
          style: {
            display: 'flex',
            alignItems: 'center',
            columnGap: '10px',
          },
        },
        [
          h(SvgIcon, {
            name: 'postgresql',
            style: {
              width: 24,
              height: 24,
            },
          }),
          h('span', null, 'postgres'),
        ]
      )
    },
    value: 'postgres',
    disabled: true,
  },
  {
    label: () => {
      return h(
        'span',
        {
          style: {
            display: 'flex',
            alignItems: 'center',
            columnGap: '10px',
          },
        },
        [
          h(SvgIcon, {
            name: 'sqlite',
            style: {
              width: 24,
              height: 24,
            },
          }),
          h('span', null, 'sqlite'),
        ]
      )
    },
    value: 'sqlite',
    disabled: true,
  },
]

export const dateTypeOptions: SelectOption[] = [
  {
    label: 'ONLY_DATE',
    value: 'ONLY_DATE',
  },
  {
    label: 'TIME_PACK',
    value: 'TIME_PACK',
  },
]

export const disableTemplateOptions: SelectOption[] = [
  {
    label: 'entity',
    value: 'ENTITY',
  },
  {
    label: 'service',
    value: 'SERVICE',
  },
  {
    label: 'service-impl',
    value: 'SERVICEIMPL',
  },
  {
    label: 'controller',
    value: 'CONTROLLER',
  },
  {
    label: 'mapper',
    value: 'MAPPER',
  },
  {
    label: 'mapper-xml',
    value: 'XML',
  },
]
