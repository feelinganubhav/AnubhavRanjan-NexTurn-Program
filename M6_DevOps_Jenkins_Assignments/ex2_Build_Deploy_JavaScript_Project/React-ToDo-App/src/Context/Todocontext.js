import { useContext } from "react";
import { createContext } from "react";

export const Todocontext = createContext({
    //ye ek variable hi hai jo ki array hai or wo array objects ka hai
    //todo me mujee kya kya chiaye hoga ek array containing id, uska message and wo complete hua hai ya nahi
    //target dte v add kar skte hai
    todos: [
        {
            id: 1,
            message: "Todo msg",
            completed: false,
            targetdate:"10/10/2024",
        }
    ],

    //functionality
    //to do add karne ke liye msg chaiye hoga or uka id set kar denge baad me
    //abb todo ban gaya to uska id hoga to usko update karne ke liye id lgegega and wapas se todo lagega ki kisse update karna hai
    //delete karne ke liye id lagega
    //complete hai ya nahi uske liye id lagega

    addtodo: (message) => {},
    updatetodo: (id, message) => {},
    deletetodo: (id) => {},
    togglecomplete: (id) =>{},
    targetDate: (id, targetdate) => {}

})

export const TodocontextProvider = Todocontext.Provider

export default function UseTodocontext(){
    // usecntext hai to usko ek context dena padega ki kiske bare me baat kar rahe hai
    return useContext(Todocontext)
}