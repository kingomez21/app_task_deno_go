import type { JSX } from "preact/jsx-runtime";
import { InputForms } from "../components/InputForms.tsx";

const NamesForms = [
    "firstname",
    "lastname",
    "email",
    "address"
]

export default function Forms(){

    const submit = (e: JSX.TargetedSubmitEvent<HTMLFormElement>) => {
        e.preventDefault()
    }

    return(
        <form onSubmit={(e) => submit(e)} method="POST" action="http://localhost:4000/users" >   
            {NamesForms.map((value) => (
                <InputForms name={value}  />
            ))}

            <button type="submit" >Crear Usuario</button>
        </form>
    )
}