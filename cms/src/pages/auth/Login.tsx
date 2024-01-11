import { ChangeEvent, useContext, useState } from 'react'
import { AuthContext } from './AuthProvider'

interface Payload {
  email: string 
  password: string 
}

export default function Login() {
  const [ payload, setPayload ] = useState<Payload>({email:'', password: ''})
  const [ error , setError ] = useState<string | null>(null)
  const { loginUser } = useContext(AuthContext)!

  const handleOnChange = (e: ChangeEvent<HTMLInputElement>, key: string) => {
    setPayload({...payload, [key]: e.target.value})
  }

  const login = () => {
    loginUser(payload.email, payload.password)
      .then(() => {
        window.location.replace('/dashboard')
      })
      .catch(() => {
        setError('Failed to login')
      })
  } 
  

  return (
    <section className="flex items-center justify-center h-screen bg-ltccrem">
      <div className="border border-ltcbrown p-10 bg-white">
        <div className="text-lg font-semibold">
          Dashboard Login
        </div>
        <div className="my-2">
          <label htmlFor="email" className="block mb-2 text-sm font-medium text-gray-800">
            Email
          </label>
          <input 
            type="email" 
            id="email" 
            className="bg-gray-50 border border-gray-900 text-gray-900 text-sm rounded-sm focus:outline-none p-[0.5em]"
            onChange={(e) => handleOnChange(e, 'email')}
          />
        </div>
        <div className="mb-5">
          <label htmlFor="passsword" className="block mb-2 text-sm font-medium text-gray-800">
            Password
          </label>
          <input 
            type="password" 
            id="passsword" 
            className="bg-gray-50 border border-gray-900 text-gray-900 text-sm rounded-sm focus:outline-none p-[0.5em]"
            onChange={(e) => handleOnChange(e, 'password')}
          />
        </div>
        { error && (
          <div className='my-2 bg-red-600 py-2 px-4 rounded-md'>
            <h3 className='text-white'>
              {error}
            </h3>
          </div>
        )}
        <div className="mb-5">
          <button 
            className="bg-ltcbrown text-white px-4 py-2 rounded-md border border-ltcbrown duration-300 ease-in-out hover:bg-white hover:text-ltcbrown"
            onClick={login}
            type='button'
          >
            Sign In
          </button>
        </div>
      </div>
    </section>
  )
}