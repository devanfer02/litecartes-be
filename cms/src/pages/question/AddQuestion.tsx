import axios from "axios"
import { useNavigate } from "react-router-dom"
import { ChangeEvent, useState } from "react"
import FlashError from "../../components/FlashError"
import Input from "../../components/Input"

interface Question {
  category_id: string 
  task_uid: string | null
  literacy: string
  answer: string 
}

export default function AddQuestion() {
  const [ question, setQuestion ] = useState<Question>({category_id: '', task_uid:'', literacy: '', answer: ''})
  const [ error ,setError ] = useState<string | null>(null)
  const navigate = useNavigate()

  const addQuestion = async () => {
    try {
    let payload = {}

    if (question.task_uid == '') {
      payload = {
        category_id: question.category_id,
        literacy: question.literacy,
        answer: question.answer 
      }
    } else {
      payload = question 
    }
    
      const res = await axios.post(import.meta.env.VITE_API_URL + '/questions', payload)

      if (res.status != 200) {
        setError(res.data.message)
        throw new Error(res.data.message)
      }

      navigate('/questions')
      
    } catch(e) {
      setError((e as Error).toString())
      
    }
  }

  return (
    <section className="ml-72 mt-16 mr-24">
      <div className="flex">
        <div className="w-1/2">
          <h1 className="text-2xl font-semibold">
            Add Question
          </h1>
        </div>
        <div className="w-1/2 flex justify-end ">
        <a href="/questions" className="bg-blue-500 hover:bg-blue-700 duration-300 ease-in-out text-white px-4 py-2 rounded-md">
            Back To Menu
          </a>
        </div>
      </div>
      { error && (
        <FlashError message={error} />
      )}
      <div className="mt-5 p-5 border border-ltcbrown mb-10">
        <div className="mb-5">
          <label htmlFor="category" className="text-xl font-semibold block">
            Category 
          </label>
          <select name="category" id="" className="bg-ltcbrown text-white px-3 py-1 mt-2 rounded-md">
            <option 
              onClick={() => setQuestion({...question, category_id: 'LTC-APP-generated1'})}
              className="bg-ltccrem text-ltcbrown"
            >
              Pilihan Ganda
            </option>
            <option
              onClick={() => setQuestion({...question, category_id: 'LTC-APP-generated2'})}
              className="bg-ltccrem text-ltcbrown"
            >
              Isian
            </option>
            <option
              onClick={() => setQuestion({...question, category_id: 'LTC-APP-generated3'})}
              className="bg-ltccrem text-ltcbrown"
            >
              Ulasan Panjang
            </option>
          </select>
        </div>
        <Input 
          label="Task UID" 
          type="text" 
          onChange={(e: ChangeEvent<HTMLInputElement>) => setQuestion({...question, task_uid: e.target.value})}
          value={question.task_uid!}
        />
        <div className="mb-5">
          <label htmlFor="literacy" className="text-xl font-semibold block">
            Literacy Text 
          </label>
          <textarea 
            name="literacy" 
            id="" 
            className="w-full h-[150px] p-2 mt-2 border border-ltcbrown focus:outline-none rounded-md focus:border-sky-500"
            value={question.literacy}
            onChange={(e: ChangeEvent<HTMLTextAreaElement>) => setQuestion({...question, literacy: e.target.value})}
          >
          </textarea>
        </div>
        <Input 
          label="Answer (Formatted)" 
          type="text" 
          onChange={(e: ChangeEvent<HTMLInputElement>) => setQuestion({...question, answer: e.target.value})}
          value={ question.answer }
        />
        <div className="mb-5">
          <button type="button" onClick={addQuestion} className="border border-ltcbrown text-white bg-ltcbrown px-4 py-2 rounded-lg duration-300 ease-in-out hover:bg-white hover:text-ltcbrown">
            Add Question
          </button>
        </div>
      </div>
    </section>
  )
}