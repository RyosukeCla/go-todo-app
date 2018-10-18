import h from 'react-hyperscript'
import { Link } from "react-router-dom"

export default () => {
  return h('div', [
    h(Link, { to: '/hi' }, 'to hi'),
    h('div', 'hello, world!'),
  ])
}
