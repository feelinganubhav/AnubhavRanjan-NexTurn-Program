//URL -> https://to-do-pied-xi.vercel.app/
import React from 'react';
import { useState } from 'react'
import './App.css'

import { useEffect } from 'react'
import { TodocontextProvider } from './Context/Todocontext'
import { Todoform, Todoitem } from './Components'

function App() {
  const [todos, settodos] = useState([])

  const addtodo = (Message) => {
    //ye jo todo aaya wo array ke andar jana chaiye
    //direct vejunga agar to fir wo purana sab delete kar dega so phlae access liye oldtodo ka then usko spread kar diya and then jo nwe todo aaya hai usko daal denge
    settodos((oldtodos) => [...oldtodos, { id: Date.now(), ...Message}])
  }

  const updatetodo = (id, updatedtodo) => {
    settodos((oldtodos) => oldtodos.map((prevtodo) => (prevtodo.id === id ? updatedtodo : prevtodo)))
  }

  //filter iliye use kiye ki agar todo.id id se match kar jata hai to usko filter kar ke hta do nahi to baki sab ko as it is return kar do (filter new array return karta hai)
  const deletetodo = (id) => {
    settodos((oldtodos) => oldtodos.filter((todokaid) => todokaid.id !== id))
  }

  const togglecomplete = (id) => {
    settodos((oldtodos) => oldtodos.map((prevtodo) => (prevtodo.id === id ? { ...prevtodo, completed: !prevtodo.completed } : prevtodo)))
  }

  const targetDate = (id, targetDate) => {
    settodos((oldTodos) =>
      oldTodos.map((todo) => (todo.id === id ? { ...todo, targetDate } : todo))
    );
  }
  //handeling local storage
  //locall storage se jab value lete hai to wo sari string hoti hai to usko jeson me convert karna padega
  //local storage direct access kar sakte hai jab tak react me hai
  useEffect(() => {
    const valuegotfromlocalstorage = JSON.parse(localStorage.getItem("key"))
    if(valuegotfromlocalstorage && valuegotfromlocalstorage.length > 0){
      settodos(valuegotfromlocalstorage)
    }
  }, [])

  useEffect(() => {
    localStorage.setItem("key", JSON.stringify(todos))
  }, [todos])


  return (
    <TodocontextProvider value={{ todos, addtodo, updatetodo, togglecomplete, deletetodo, targetDate}}>
      <div className="bg-[#172842] w-full min-h-screen py-8">
        <div className="w-full max-w-2xl mx-auto bg-[#172233] shadow-lg rounded-lg px-4 py-3 text-white">
          <h1 className="text-2xl font-bold text-center mb-8 mt-2">Manage Your Todos</h1>
          <div className="mb-4">
            <Todoform/>
          </div>
          <div className="flex flex-wrap gap-y-3">
            {/*Loop and Add TodoItem here */}
            {todos.map((todo) => (
              <div key={todo.id}
              className='w-full'
              >
                <Todoitem todo={todo}/>
              </div>
            ))}
          </div>
        </div>
      </div>
    </TodocontextProvider>
  )
}

export default App
