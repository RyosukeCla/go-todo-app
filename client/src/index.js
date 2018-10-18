import ReactDOM from 'react-dom'
import h from 'react-hyperscript'
import { BrowserRouter as Router, Route, Link } from "react-router-dom"
import routes from './routes'

ReactDOM.render(
  h(Router, [
    h('div', [
      ...routes.map(route => h(Route, route))
    ])
  ]),
  document.getElementById('root')
)
