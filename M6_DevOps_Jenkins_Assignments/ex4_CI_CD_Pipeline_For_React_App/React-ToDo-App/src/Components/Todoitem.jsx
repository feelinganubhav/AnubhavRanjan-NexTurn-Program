import React from 'react'
import UseTodocontext from '../Context/Todocontext';
import { useState } from 'react';

function Todoitem({todo}) {
    const [isTodoEditable, setIsTodoEditable] = useState(false)
    const {updatetodo, deletetodo, togglecomplete} = UseTodocontext()
    const [todoMsg, setTodoMsg] = useState(todo.todo)
    const [targetDate, setTargetDate] = useState(todo.targetDate)

    const editTodo = ()=> {
        updatetodo(todo.id, {...todo, todo: todoMsg, targetDate})
        setIsTodoEditable(false)
    }

    const toggleCompleted = () => {
        togglecomplete(todo.id)
    }

    return (
        <div
            className={`flex border border-black/10 rounded-lg px-3 py-1.5 gap-x-3 shadow-sm shadow-white/50 duration-300  text-black ${todo.completed ? "bg-[#c6e9a7]" : "bg-[#ccbed7]"
                }`}
        >
            <input
                type="checkbox"
                className="cursor-pointer"
                checked={todo.complete}
                onChange={toggleCompleted}
            />
            <input
                type="text"
                className={`border outline-none w-full bg-transparent rounded-lg ${isTodoEditable ? "border-black/10 px-2" : "border-transparent"
                    } ${todo.completed ? "line-through" : ""}`}
                value={todoMsg}
                onChange={(e) => setTodoMsg(e.target.value)}
                readOnly={!isTodoEditable}
            />
            <input
                type="date"
                className={`border outline-none ml-8 w-full bg-transparent rounded-lg ${isTodoEditable ? "border-black/10 px-2" : "border-transparent"
                    } ${todo.completed ? "line-through" : ""}`}
                value={targetDate}
                onChange={(e) => setTargetDate(e.target.value)}
                readOnly={todo.completed}
            />
            {/* Edit, Save Button */}
            <button
                className="inline-flex w-8 h-8 rounded-lg text-sm border border-black/10 justify-center items-center bg-gray-50 hover:bg-gray-100 shrink-0 disabled:opacity-50"
                onClick={() => {
                    if (todo.complete) return;

                    if (isTodoEditable) {
                        editTodo();
                    } else setIsTodoEditable((prev) => !prev);
                }}
                disabled={todo.completed}
            >
                {isTodoEditable ? "📁" : "✏️"}
            </button>
            {/* Delete Todo Button */}
            <button
                className="inline-flex w-8 h-8 rounded-lg text-sm border border-black/10 justify-center items-center bg-gray-50 hover:bg-gray-100 shrink-0"
                onClick={() => deletetodo(todo.id)}
            >
                ❌
            </button>
        </div>
    )
}

export default Todoitem
