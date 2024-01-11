import { ReactNode, createContext, useEffect, useState } from "react";
import { auth } from "../../utils/firebase";
import { 
  User, 
  UserCredential, 
  createUserWithEmailAndPassword, 
  onAuthStateChanged, 
  signInWithEmailAndPassword,
  signOut
} from "firebase/auth";

interface Props {
  children: ReactNode
}

interface Auth {
  user: User | null 
  loading: boolean 
  loginUser: (email: string, password: string) => Promise<UserCredential> 
  logoutUser: () => Promise<void>
  registerUser: (email: string, password: string) => Promise<UserCredential>
}

export const AuthContext = createContext<Auth | null>(null)

export const AuthProvider = ({children}: Props) => {
  const [ currentUser, setCurrentUser ] = useState<User | null>(null)
  const [ loading, setLoading ] = useState<boolean>(true)

  const loginUser = (email: string, password: string) => {
    setLoading(true)
    return signInWithEmailAndPassword(auth, email, password)
  }

  const registerUser = (email: string, password: string) => {
    setLoading(true)
    return createUserWithEmailAndPassword(auth, email, password)
  }

  const logoutUser = () => {
    setLoading(true)
    return signOut(auth)
  }

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, (user) => {
      setCurrentUser(user)
      setLoading(false)
    })

    return () => {
      unsubscribe()
    }
  }, [])

  const authValue: Auth = {
    user: currentUser, 
    loading, 
    loginUser, 
    logoutUser,
    registerUser 
  }

  return (
    <AuthContext.Provider value={authValue}>
      {children}
    </AuthContext.Provider>
  )
}

export default AuthProvider