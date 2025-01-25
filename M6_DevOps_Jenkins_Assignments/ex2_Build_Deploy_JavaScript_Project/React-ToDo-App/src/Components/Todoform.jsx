import React from 'react'
import { useState } from 'react'
import UseTodocontext from '../Context/Todocontext'

function Todoform() {
    const [todo, settodo] = useState("")
    const [targetDate, setTargetDate] =  useState("")
    const {addtodo} = UseTodocontext()

    const add = (e) => {
        e.preventDefault()
        if(!todo)  return
        addtodo({ todo, completed: false, targetDate})
        settodo("")
        setTargetDate("")
    }
    return (
        <form onSubmit={add} className="flex flex-wrap items-end">
            <div className='w-full flex flex-wrap gap-x-2'>
            <input
                type="text"
                placeholder="Write Todo..."
                className=" flex-grow border border-black/10 rounded-lg px-3 outline-none duration-150 bg-white/20 py-1.5"
                value={todo}
                onChange={(e) => settodo(e.target.value)}
            />
           
            <input
                type="date"
                placeholder="Write Date..."
                className=" flex-grow border border-black/10 rounded-lg px-3 outline-none duration-150 bg-white/20 py-1.5 "
                value={targetDate}
                onChange={(e) => setTargetDate(e.target.value)}
            />
            <button type="submit" className=" rounded-lg px-6 py-1 bg-green-600 text-white shrink-0 ml-1">
                Add
            </button>
            </div>
            
        </form>
    )
}

export default Todoform
