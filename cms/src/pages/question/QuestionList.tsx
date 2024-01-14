import axios from "axios"
import useFetch from "../../utils/useFetch"
import FlashError from "../../components/FlashError"

type Question = {
  uid: string 
  category_id: number 
  task_uid: string
  literacy: string 
  answer: string 
}

export default function QuestionList() {
  const [ questions, refetch, error ] = useFetch<Question>(import.meta.env.VITE_API_URL + '/questions')

  const deleteQuestion = async (uid: string) => {
    try {
      const res = await axios.delete(import.meta.env.VITE_API_URL + `/questions/${uid}`)

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
          <h1 className="text-2xl font-semibold">Questions</h1>
        </div>
        <div className="w-1/2 flex justify-end">
          <a href="/questions/add" className="bg-blue-500 hover:bg-blue-700 duration-300 ease-in-out text-white px-4 py-2 rounded-md">
            Add Question 
          </a>
        </div>
      </div>
      <div className="relative overflow-x-auto mt-5">
        { error ? (
          <FlashError message={(error as Error).toString()} />
        ) : (
        <table className="w-full text-sm text-left rtl:text-right text-ltc-brown">
            <thead className="text-xs uppercase text-ltccrem bg-ltcbrown">
                <tr>
                    <th scope="col" className="px-6 py-3">
                        No 
                    </th>
                    <th scope="col" className="px-6 py-3 text-center">
                        Category ID
                    </th>
                    <th scope="col" className="px-6 py-3 text-center">
                        Task UID
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Literacy
                    </th>
                    <th scope="col" className="px-6 py-3 text-center">
                        Answer
                    </th>
                    <th scope="col" className="px-6 py-3 text-center">
                      Action
                    </th>
                </tr>
            </thead>
            <tbody>
              { questions.map((question, index) => (
                <tr className="border-b bg-ltccrem">
                    <th scope="row" className="px-6 py-4 font-medium text-ltcbrown whitespace-nowrap ">
                        { index + 1 }
                    </th>
                    <td className="px-6 py-4 text-center">
                        { question.category_id }
                    </td>
                    <td className="px-6 py-4 text-center"> 
                      { question.task_uid }
                    </td>
                    <td className="px-6 py-4">
                        { question.literacy }
                    </td>
                    <td className="px-6 py-4 text-center">
                        { question.answer }
                    </td>
                    <td className="px-6 py-4 text-center flex text-white font-semibold justify-center items-center">
                        <a href={`/questions/edit/${question.uid}`} className="bg-green-600 hover:bg-green-900 duration-300 ease-in-out mx-1 px-4 py-2 rounded-md">
                          Edit Question
                        </a>
                        <button 
                          type="button" 
                          className="bg-red-600 hover:bg-red-900 duration-300 ease-in-out mx-1 px-4 py-2 rounded-md"
                          onClick={() => deleteQuestion(question.uid)}
                        >
                          Delete Question
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