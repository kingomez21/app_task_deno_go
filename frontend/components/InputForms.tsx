import type { JSX } from "preact/jsx-runtime";

export function InputForms(props: JSX.HTMLAttributes<HTMLInputElement>) {
    return (
        <div>
            <label for="price" className="block text-sm font-medium leading-6 text-gray-900">{props.name}</label>
            <div className="relative mt-2 rounded-md shadow-sm">
                <input 
                    {...props}
                    type="text"
                    className="block w-full rounded-md border-0 py-1.5 pl-7 pr-20 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
            </div>
            
            
        </div>
    )
}