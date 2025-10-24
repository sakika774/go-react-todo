import { useEffect, useState } from 'react'

type Todo = {
  id: number
  task: string
  done: boolean
}

function App() {
  const [todos, setTodos] = useState<Todo[]>([])
  const [newTask, setNewTask] = useState('')

  const fetchTodos = () => {
    fetch('http://localhost:8080/todos')
      .then(res => res.json())
      .then(data => setTodos(data))
  }

  const deleteTodo = (id: number) => {
    fetch(`http://localhost:8080/todos?id=${id}`, { method: 'DELETE' })
      .then(() => fetchTodos())
  }

  const toggleTodo = (id: number) => {
    fetch(`http://localhost:8080/todos/toggle?id=${id}`, { method: 'PATCH' })
      .then(() => fetchTodos())
  }

  useEffect(() => {
    fetchTodos()
  }, [])

  const addTodo = () => {
    if (!newTask) return
    fetch('http://localhost:8080/todos', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ task: newTask })
    })
      .then(res => res.json())
      .then(() => {
        setNewTask('')
        fetchTodos()
      })
  }

  return (
    <div>
      <h1>Todo List</h1>
      <ul>
        {todos.map(todo => (
          <li key={todo.id}>
            <input
              type="checkbox"
              checked={todo.done}
              onChange={() => toggleTodo(todo.id)}
            />
            <span style={{ textDecoration: todo.done ? 'line-through' : 'none' }}>
              {todo.task}
            </span>
            <button onClick={() => deleteTodo(todo.id)}>削除</button>
          </li>
        ))}
      </ul>
      <input
        type="text"
        value={newTask}
        onChange={e => setNewTask(e.target.value)}
        placeholder="新しいTodo"
      />
      <button onClick={addTodo}>追加</button>
    </div>
  )
}

export default App