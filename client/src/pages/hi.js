import h from 'react-hyperscript'
import { Link } from "react-router-dom"
import Todo from '../components/Todo'

export default () => {
  return h('div', [
    h(Link, { to: '/' }, 'go back'),
    h(Todo),
    h(Todo),
    h(Todo)
  ])
}
