import { useEffect, useState } from 'react'

type Todo = {
  id: number
  task: string
}

function App() {
  const [todos, setTodos] = useState<Todo[]>([])

  useEffect(() => {
    fetch('http://localhost:8080/todos')
      .then(res => res.json())
      .then(data => setTodos(data))
  }, [])

  return (
    <div>
      <h1>Todo List</h1>
      <ul>
        {todos.map(todo => (
          <li key={todo.id}>{todo.task}</li>
        ))}
      </ul>
    </div>
  )
}

export default App