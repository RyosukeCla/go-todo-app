import IndexPage from './pages/index'
import HiPage from './pages/hi'

export default [
  { exact: true, path: '/', component: IndexPage },
  { path: '/hi', component: HiPage }
]

