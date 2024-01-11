import { ChangeEvent, useContext, useState } from 'react'
import axios from 'axios'
import { AuthContext } from './AuthProvider'
import { updateProfile } from 'firebase/auth'

interface Payload {
  email: string 
  password: string 
  username: string
}

export default function Register() {
  const [ payload, setPayload ] = useState<Payload>({email:'', password: '', username: ''})
  const [ error , setError ] = useState<string | null>(null)
  const { registerUser } = useContext(AuthContext)!

  const handleOnChange = (e: ChangeEvent<HTMLInputElement>, key: string) => {
    setPayload({...payload, [key]: e.target.value})
  }

  const register = async () => {
    try {
      const userCredential = await registerUser(payload.email, payload.password)
      const user = userCredential.user 

      await updateProfile(user, {
        displayName: payload.username 
      })

      await axios.post(import.meta.env.VITE_API_URL + `/user/${user.uid}`)

      console.log(user)

      window.location.replace('/dashboard')
    } catch (e) {
      setError('Failed to register')
      console.log(e)
    }

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
        <div className="my-2">
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
        <div className="mt-2 mb-5">
          <label htmlFor="username" className="block mb-2 text-sm font-medium text-gray-800">
            Username 
          </label>
          <input 
            type="username" 
            id="username" 
            className="bg-gray-50 border border-gray-900 text-gray-900 text-sm rounded-sm focus:outline-none p-[0.5em]"
            onChange={(e) => handleOnChange(e, 'username')}
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
            onClick={register}
            type='button'
          >
            Sign In
          </button>
        </div>
      </div>
    </section>
  )
}