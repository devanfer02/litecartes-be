import { useState } from "react"
// import useFetch from "../../utils/useFetch"

interface Question {
  category_id: number 
  literacy: string 
  answer: string 
}

const questionsList = [
  {
    category_id: 1,
    literacy: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Totam expedita non quidem deserunt, accusantium autem ab ullam rerum optio fugiat tempora placeat delectus. Illo debitis voluptatibus reprehenderit dolorum, quas sequi?',
    answer: 'A'
  },
  {
    category_id: 1,
    literacy: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Totam expedita non quidem deserunt, accusantium autem ab ullam rerum optio fugiat tempora placeat delectus. Illo debitis voluptatibus reprehenderit dolorum, quas sequi?',
    answer: 'A'
  }
]

export default function QuestionList() {
  // const [ users, refetch, error ] = useFetch<User>(import.meta.env.VITE_API_URL + '/user')
  const [ questions ] = useState<Question[]>(questionsList)
  const [ error ] = useState<string | null>(null)

  // useEffect(() => {
  //   refetch()
  // }, [])

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
                        Category ID
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Literacy
                    </th>
                    <th scope="col" className="px-6 py-3">
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
                    <td className="px-6 py-4">
                        { question.literacy }
                    </td>
                    <td className="px-6 py-4 text-center">
                        { question.answer }
                    </td>
                    <td className="px-6 py-4 text-center flex text-white font-semibold">
                        <a href="/questions/edit" className="bg-green-600 hover:bg-green-900 duration-300 ease-in-out mx-1 px-4 py-2 rounded-md">
                          Edit Question
                        </a>
                        <button type="button" className="bg-red-600 hover:bg-red-900 duration-300 ease-in-out mx-1 px-4 py-2 rounded-md">
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