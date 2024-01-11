import { Icon } from '@iconify/react'
import { useContext } from 'react'
import { AuthContext } from '../pages/auth/AuthProvider'
import { NavLink } from 'react-router-dom'

const navs = [
  {
    name: 'Home',
    link: '/dashboard',
    icon: 'mdi:home'
  },
  {
    name: 'Users',
    link: '/users',
    icon: 'mdi:account-group'
  },
  {
    name: 'Questions',
    link: '/questions',
    icon: 'mdi:message-question'
  },
  {
    name: 'Task',
    link: '/tasks',
    icon: 'mdi:list-box'
  }
]

export default function Sidebar() {
  const { logoutUser } = useContext(AuthContext)!

  const handleSignOut = async () => {
    await logoutUser()
  }

  return (
    <nav
      className="sidebar fixed top-0 bottom-0 lg:left-0 p-2 w-[220px] overflow-y-auto text-center bg-ltcbrown"
    >
      <div className="text-gray-100 text-xl">
        <div className="p-2.5 mt-1 flex items-center">
          <h1 className="font-bold text-ltccrem text-xl ml-3">Litecartes CMS</h1>
          <i
            className="bi bi-x cursor-pointer ml-28 lg:hidden"
          ></i>
        </div>
      </div>
      <hr className="my-2 bg-gray-600 h-[1px]"/>
      { navs.map((nav, index) => (
        <NavLink 
          className="p-2.5 mt-3 flex items-center rounded-md px-4 duration-300 cursor-pointer hover:bg-ltccrem hover:text-ltcbrown text-ltccrem group sidebar"
          key={index}
          to={nav.link}
        >
          <Icon icon={nav.icon} className='w-[24px] h-[24px]'/>
          <span className="text-[15px] ml-4 font-bold">{ nav.name }</span>
        </NavLink>
      ))}
      <hr className="my-4 bg-gray-600 h-[1px]"/>
      <div
        className="p-2.5 mt-3 flex items-center rounded-md px-4 duration-300 cursor-pointer hover:text-ltcbrown text-ltccrem group hover:bg-ltccrem"
        onClick={handleSignOut}
      >
        <Icon icon={'mdi:logout'} className='w-[24px] h-[24px]' />
        <span className="text-[15px] ml-4 group-hover:text-ltcbrown font-bold">Logout</span>
      </div>
    </nav>
  )
}