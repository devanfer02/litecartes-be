import { ChangeEvent, useEffect, useState } from "react"
import FlashError from "../../components/FlashError"
import axios from "axios"
import { useNavigate, useParams } from "react-router-dom"
import Input from "../../components/Input"

type Task = {
  level: number
  sign: string
  level_category_id: number 
}


export default function AddTask() {
  const [ task, setTask ] = useState<Task>({level: 0, sign: '', level_category_id: 0})
  const [ error, setError ] = useState<string | null>(null)
  const { uid } = useParams()
  const navigate = useNavigate()

  const updateTask = async () => {
    try {
      const res =  await axios.put(import.meta.env.VITE_API_URL + `/tasks/${uid}`, task)

      if (res.status != 200) {
        throw new Error(res.data.message)
      }

      navigate('/tasks')
    } catch (e) {
      setError((e as Error).toString())
    }
  }

  useEffect(() => {
    const fetchTask = async () => {
      try {
        const res = await axios.get(import.meta.env.VITE_API_URL + `/tasks/${uid}`)
  
        if (res.status != 200) {
          throw new Error(res.data.message)
        }
  
        setTask(res.data.data)
      } catch (e) {
        setError((e as Error).toString()) 
      }
    }

    fetchTask()
  }, [uid])

  return (
    <section className="ml-72 mt-16 mr-24">
      <div className="flex">
        <div className="w-1/2">
          <h1 className="text-2xl font-semibold">
            Edit Task
          </h1>
        </div>
        <div className="w-1/2 flex justify-end ">
        <a href="/questions" className="bg-blue-500 hover:bg-blue-700 duration-300 ease-in-out text-white px-4 py-2 rounded-md">
            Back To Menu
          </a>
        </div>
      </div>
      { error && (
        <FlashError message={error!} />
      )}
      <div className="mt-5 p-5 border border-ltcbrown mb-10">
        <div className="mb-5">
          <label htmlFor="category" className="text-xl font-semibold block">
            Category 
          </label>
          <select name="category" id="" className="bg-ltcbrown text-white px-3 py-1 mt-2 rounded-md">
            <option 
              onClick={() => setTask({...task, level_category_id: 1})}
              className="bg-ltccrem text-ltcbrown"
            >
              Litera Cookie 
            </option>
            <option
              onClick={() => setTask({...task, level_category_id: 2})}
              className="bg-ltccrem text-ltcbrown"
            >
              Voyager
            </option>
            <option
              onClick={() => setTask({...task, level_category_id: 3})}
              className="bg-ltccrem text-ltcbrown"
            >
              Luminary 
            </option>
            <option
              onClick={() => setTask({...task, level_category_id: 4})}
              className="bg-ltccrem text-ltcbrown"
            >
              Mastermind 
            </option>
          </select>
        </div>
        <Input 
          label="Level" 
          type="number" 
          onChange={(e: ChangeEvent<HTMLInputElement>) => setTask({...task, level: parseInt(e.target.value)})}
          value={task.level }
        />
        <Input 
          label="Sign" 
          type="text" 
          onChange={(e: ChangeEvent<HTMLInputElement>) => setTask({...task, sign: e.target.value})}
          value={task.sign}
        />
        <div className="mb-5">
          <button type="button" onClick={updateTask} className="border border-ltcbrown text-white bg-ltcbrown px-4 py-2 rounded-lg duration-300 ease-in-out hover:bg-white hover:text-ltcbrown">
            Update Task
          </button>
        </div>
      </div>
    </section>
  )
}