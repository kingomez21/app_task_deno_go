
type Person = {
    id: string
    firstname: string
    lastname: string
    email: string
    address: string
}

type ListProps = {
    people: Person[]
}

export default function ListView({people}: ListProps) {
    return (
        <ul role="list" className="divide-y divide-gray-100">
        {people.map((person) => (
            <li key={person.id} className="flex justify-between gap-x-6 py-5">
            <div className="flex min-w-0 gap-x-4">
                <img alt="" src="https://cdn-icons-png.flaticon.com/512/6858/6858504.png" className="h-12 w-12 flex-none rounded-full bg-gray-50" />
                <div className="min-w-0 flex-auto">
                <p className="text-sm font-semibold leading-6 text-gray-900">{person.firstname+" "+person.lastname}</p>
                <p className="mt-1 truncate text-xs leading-5 text-gray-500">{person.email}</p>
                </div>
            </div>
            </li>
        ))}
    </ul>
    )
    
}