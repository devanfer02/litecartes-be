import { Link } from "react-router-dom"

export default function Dashboard() {
  return (
    <section className="ml-72 mt-16 mr-24">
      <div>
        <h1 className="text-3xl font-bold text-ltcbrown">
          Litecartes Content Management System 
        </h1>
        <hr className="text-ltcbrown bg-ltcbrown w-full h-1"/>
      </div>
      <div className="mt-5">
        <p className="text-lg">
          Litecartes CMS provides a very easy content management system to monitor items, inserting new items,
          updating new items, and deleting items
        </p>
      </div>
      <div className="mt-4">
        <h4 className="text-lg">Listed Items : </h4>
        <ol>
          <li><Link to="/tasks">Tasks</Link></li>
          <li><Link to="/users">Users</Link></li>
          <li><Link to="/questions">Questions</Link></li>
        </ol>
      </div>
    </section>
  )
}