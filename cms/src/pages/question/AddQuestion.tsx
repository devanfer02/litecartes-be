export default function AddQuestion() {
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
      <div className="mt-5 p-5 border border-ltcbrown">
        <div className="mb-5">
          <label htmlFor="category" className="text-xl font-semibold block">
            Category 
          </label>
          <select name="category" id="" className="bg-ltcbrown text-white px-3 py-1 mt-2 rounded-md">
            <option value={1}>Pilihan Ganda</option>
            <option value={2}>Isian</option>
          </select>
        </div>
        <div className="mb-5">
          <label htmlFor="literacy" className="text-xl font-semibold block">
            Literacy Text 
          </label>
          <textarea name="literacy" id="" className="w-full h-[150px] p-2 mt-2 border border-ltcbrown focus:outline-none rounded-md focus:border-sky-500">
          </textarea>
        </div>
        <div className="mb-5">
          <label htmlFor="" className="text-xl font-semibold block">
            Answer (Formatted)
          </label>
          <input type="text" className="border border-ltcbrown px-2 py-2 w-full mt-2 rounded-md focus:outline-none focus:border-sky-500"/>
        </div>
        <div className="mb-5">
          <button type="button" className="border border-ltcbrown text-white bg-ltcbrown px-4 py-2 rounded-lg duration-300 ease-in-out hover:bg-white hover:text-ltcbrown">
            Add Question
          </button>
        </div>
      </div>
    </section>
  )
}