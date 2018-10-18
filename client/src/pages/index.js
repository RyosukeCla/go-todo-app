import h from 'react-hyperscript'
import React from 'react'

export default class extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      todos: [],
      currentText: ''
    }
  }

  async componentWillMount() {
    const data = await fetch('/api/todo/get', { method: 'GET' })
    const todos = await data.json()
    this.setState({
      todos
    })
  }

  handleInputForm(e) {
    this.setState({
      currentText: e.target.value
    })
  }

  async handleClick() {
    const data = await fetch('/api/todo/add', {
      method: 'POST',
      body: JSON.stringify({
        text: this.state.currentText
      }),
      headers: {
        'Content-Type': 'application/json'
      }
    })

    const newTodo = await data.json()

    this.setState({
      currentText: '',
      todos: [...this.state.todos, newTodo]
    })
  }

  async handleCheck(todo) {
    const data = await fetch('/api/todo/check', {
      method: 'POST',
      body: JSON.stringify({
        id: todo.id
      }),
      headers: {
        'Content-Type': 'application/json'
      }
    })

    const updateTodo = await data.json()

    this.setState({
      todos: this.state.todos.map(todo => todo.id === updateTodo.id ? updateTodo : todo)
    })
  }

  async handleDone(todos) {
    await fetch('/api/todo/done', {
      method: 'POST',
      body: JSON.stringify({
        ids: todos.map(todo => todo.id)
      }),
      headers: {
        'Content-Type': 'application/json'
      }
    })

    this.setState({
      todos: this.state.todos.filter(todo => !todos.some(_todo => _todo.id === todo.id))
    })
  }

  render() {
    const Todos = this.state.todos.map(todo => {
      return h('div', [
        h('input', { type: 'checkbox', id: todo.id, checked: todo.checked, onChange: () => this.handleCheck(todo) }),
        h('label', { htmlFor: todo.id }, todo.text)
      ])
    })

    const checkedTodos = this.state.todos.filter(todo => todo.checked)
    const areSomeChecked = checkedTodos.length > 0
    const isAddButtonShow = this.state.currentText.length > 0

    return h('div', [
      h('h1', 'simple todo app'),
      h('p', 'built with react & go. no design, just it.'),
      h('div', [
        h('input', { value: this.state.currentText, onChange: (e) => this.handleInputForm(e) }),
        isAddButtonShow ? h('button', { onClick: () => this.handleClick() }, 'add') : undefined,
        areSomeChecked ? h('button', { onClick: () => this.handleDone(checkedTodos) }, 'done') : undefined,
      ]),
      h('div', [
        ...Todos
      ]),
    ])
  }
}
