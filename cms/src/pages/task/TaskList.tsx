import axios from "axios"
import useFetch from "../../utils/useFetch"

type Task = {
  uid: string 
  level: number
  sign: string
  level_category_id: number 
}

const category: string[] = [
  'Litera Cookie',
  'Voyager',
  'Luminary',
  'Mastermind'
]

export default function TaskList() {
  const [ tasks, refetch, error ] = useFetch<Task>(import.meta.env.VITE_API_URL + '/tasks')

  const deleteTask = async (uid: string) => {
    try {
      const res = await axios.delete(import.meta.env.VITE_API_URL + `/tasks/${uid}`)

      if (res.status != 200)  { 
        throw Error("question not found")
      }

      refetch()
    } catch(e) {
      console.log(e as Error)
    }
  }

  return (
    <section className="ml-72 mt-16 mr-24">
      <div className="flex">
        <div className="w-1/2">
          <h1 className="text-2xl font-semibold">Tasks</h1>
        </div>
        <div className="w-1/2 flex justify-end">
          <a href="/tasks/add" className="bg-blue-500 hover:bg-blue-700 duration-300 ease-in-out text-white px-4 py-2 rounded-md">
            Add Task 
          </a>
        </div>
      </div>
      <div className="relative overflow-x-auto mt-5">
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
                        UID
                    </th>
                    <th scope="col" className="px-6 py-3 text-center">
                        Level
                    </th>
                    <th scope="col" className="px-6 py-3 text-center">
                        Sign 
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Level Category
                    </th>
                    <th scope="col" className="px-6 py-3 text-center">
                      Action
                    </th>
                </tr>
            </thead>
            <tbody>
              { tasks.map((task, index) => (
                <tr className="border-b bg-ltccrem">
                    <th scope="row" className="px-6 py-4 font-medium text-ltcbrown whitespace-nowrap ">
                        { index + 1 }
                    </th>
                    <td className="px-6 py-4">
                        { task.uid }
                    </td>
                    <td className="px-6 py-4 text-center">
                        { task.level }
                    </td>
                    <td className="px-6 py-4 text-center"> 
                      { task.sign }
                    </td>
                    <td className="px-6 py-4">
                        { category[task.level_category_id-1] }
                    </td>
                    <td className="px-6 py-4 text-center flex text-white font-semibold justify-center items-center">
                        <a href={`/tasks/edit/${task.uid}`} className="bg-green-600 hover:bg-green-900 duration-300 ease-in-out mx-1 px-4 py-2 rounded-md">
                          Edit Task
                        </a>
                        <button 
                          type="button" 
                          className="bg-red-600 hover:bg-red-900 duration-300 ease-in-out mx-1 px-4 py-2 rounded-md"
                          onClick={() => deleteTask(task.uid)}
                        >
                          Delete Task
                        </button>
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