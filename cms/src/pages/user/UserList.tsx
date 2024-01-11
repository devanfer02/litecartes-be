import { useEffect, useState } from "react"
import useFetch from "../../utils/useFetch"

interface User {
  username: string
  email: string 
  subscription_id: number 
  total_exp: number
  gems: number
  streaks: number
  last_active: string 
}

export default function UserList() {
  const [ cursorReq, setCursorReq ] = useState("")
  const [ next, setNext ] = useState(true)
  const [ users, refetch, error, loading, cursor ] = useFetch<User>(import.meta.env.VITE_API_URL + `/users?limit=8&cursor=${cursorReq}&next=${next}`)

  useEffect(() => {
    refetch()
  }, [cursorReq, next])

  return (
    <section className="ml-72 mt-16">
      <div>
        <h1 className="text-2xl font-semibold">Users</h1>
      </div>
      <div className="my-4 flex">
        <button 
          className="font-semibold text-lg bg-ltcbrown text-white px-2 rounded-md border border-ltcbrown hover:bg-white hover:text-ltcbrown duration-200 ease-in-out mr-2"
          onClick={() => { setCursorReq(cursor!.prev_cursor); setNext(false)}}
        >
          &lt;
        </button>
        <button 
          className="font-semibold text-lg bg-ltcbrown text-white px-2 rounded-md border border-ltcbrown hover:bg-white hover:text-ltcbrown duration-200 ease-in-out ml-2"
          onClick={() =>{ setCursorReq(cursor!.next_cursor); setNext(true) }}
        >
          &gt;
        </button>
      </div>
      <div className="relative overflow-x-auto mr-24 mt-5">
        { loading && (
          <div className="p-5 text-center">
            <h1 className=" uppercase text-3xl font-bold text-ltcbrown">
              Loading Users Data 
            </h1>
          </div>          
        )}
        { error ? (
          <div className="bg-red-600 p-5 text-center">
            <h1 className="text-white uppercase text-3xl font-bold">
              ERROR DISPLAYING DATA
            </h1>
            <p className="text-white text-lg font-semibold">
              cek logs for more details
            </p>
          </div>
        ) : (
        <table className="w-full text-sm text-left rtl:text-right text-ltc-brown">
            <thead className="text-xs uppercase text-ltccrem bg-ltcbrown">
                <tr>
                    <th scope="col" className="px-6 py-3">
                        No 
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Username
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Email
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Subscription 
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Total Exp
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Gems
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Streaks 
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Last Active 
                    </th>
                </tr>
            </thead>
            <tbody>
              { users.map((user, index) => (
                <tr className="border-b bg-ltccrem">
                    <th scope="row" className="px-6 py-4 font-medium text-ltcbrown whitespace-nowrap ">
                        { index + 1 }
                    </th>
                    <td className="px-6 py-4">
                        { user.username }
                    </td>
                    <td className="px-6 py-4">
                        { user.email }
                    </td>
                    <td className="px-6 py-4">
                        { user.subscription_id === 1 ? 'Free' : 'Paid' }
                    </td>
                    <td className="px-6 py-4">
                        { user.total_exp }
                    </td>
                    <td className="px-6 py-4">
                        { user.gems }
                    </td>
                    <td className="px-6 py-4">
                        { user.streaks }
                    </td>
                    <td className="px-6 py-4">
                        { user.last_active }
                    </td>
                </tr>
              ))}
            </tbody>
        </table>
        )}
      </div>

    </section>
  )
}