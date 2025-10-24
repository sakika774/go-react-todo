import { useEffect, useState } from 'react'

type Todo = {
  id: number
  task: string
}

function App() {
  const [todos, setTodos] = useState<Todo[]>([])
  const [newTask, setNewTask] = useState('')

  const fetchTodos = () => {
    fetch('http://localhost:8080/todos')
      .then(res => res.json())
      .then(data => setTodos(data))
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
          <li key={todo.id}>{todo.task}</li>
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