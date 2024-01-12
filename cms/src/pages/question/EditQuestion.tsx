import { ChangeEvent, useEffect, useState } from "react"
import { useParams, useNavigate } from "react-router-dom"
import axios from "axios"

interface Question {
  category_id: string 
  literacy: string 
  answer: string 
}

export default function EditQuestion() {
  const [ question, setQuestion ] = useState<Question>({category_id:'',literacy:'',answer:''})
  const [ error ,setError ] = useState<string | null>(null)
  const { uid } = useParams()
  const navigate = useNavigate()

  useEffect(() => {
    const fetchQuestion = async () => {
      try {
        const res = await axios.get(import.meta.env.VITE_API_URL + `/questions/${uid}`)

        if (res.status != 200) {
          throw new Error 
        }

        setQuestion(res.data.data)
        navigate("/questions")
      } catch (e) {
        setError((e as Error).toString())
      } 
    }

    fetchQuestion()
  }, [uid, question])

  return (
    <section className="ml-72 mt-16 mr-24">
      <div className="flex">
        <div className="w-1/2">
          <h1 className="text-2xl font-semibold">
            Edit Question
          </h1>
        </div>
        <div className="w-1/2 flex justify-end ">
        <a href="/questions" className="bg-blue-500 hover:bg-blue-700 duration-300 ease-in-out text-white px-4 py-2 rounded-md">
            Back To Menu
          </a>
        </div>
      </div>
      { error && (
        <div className="bg-red-600 p-5 text-center">  
          <h1 className="text-white uppercase text-xl font-bold">
            Error Sending Request 
          </h1>
          <p className="text-white text-lg font-semibold">
            Error Message : { error }
          </p>
        </div>
      )}
      <div className="mt-5 p-5 border border-ltcbrown">
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
        <div className="mb-5">
          <label htmlFor="" className="text-xl font-semibold block">
            Answer (Formatted)
          </label>
          <input 
            type="text" 
            onChange={(e: ChangeEvent<HTMLInputElement>) => setQuestion({...question, answer: e.target.value})}
            value={question.answer}
            className="border border-ltcbrown px-2 py-2 w-full mt-2 rounded-md focus:outline-none focus:border-sky-500"
          />
        </div>
        <div className="mb-5">
          <button type="button" onClick={() => {}} className="border border-ltcbrown text-white bg-ltcbrown px-4 py-2 rounded-lg duration-300 ease-in-out hover:bg-white hover:text-ltcbrown">
            Add Question
          </button>
        </div>
      </div>
    </section>
  )
}