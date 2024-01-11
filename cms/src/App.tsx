import { Suspense, lazy } from "react"
import { BrowserRouter, Routes, Route } from "react-router-dom"
import ProtectedLayout from "./components/ProtectedLayout"
import ProtectedPages from "./pages/auth/ProtectedRoutes"
import Loading from "./pages/Loading"
import AuthProvider from "./pages/auth/AuthProvider"
import NotFound from "./pages/NotFound"

const Login             = lazy(() => import('./pages/auth/Login'))
const Register          = lazy(() => import('./pages/auth/Register'))

// protected pages
const Dashboard         = lazy(() => import("./pages/Dashboard"))
const UserList          = lazy(() => import('./pages/user/UserList'))
const QuestionList      = lazy(() => import('./pages/question/QuestionList'))
const TaskList          = lazy(() => import('./pages/task/TaskList'))
const AddQuestion       = lazy(() => import('./pages/question/AddQuestion'))
const EditQuestion      = lazy(() => import('./pages/question/EditQuestion'))

const protectedPages = [
  {
    path: '/dashboard',
    render: <Dashboard/>,
    title: 'Home'
  },
  {
    path: '/users',
    render: <UserList/>,
    title: 'Users'
  },
  {
    path: '/questions',
    render: <QuestionList/>,
    title: 'Questions'
  },
  {
    path: '/questions/add',
    render: <AddQuestion/>,
    title: 'Add Question'
  },
  {
    path: '/questions/edit',
    render: <EditQuestion/>,
    title: 'Edit Question'
  },
  {
    path: '/tasks',
    render: <TaskList/>,
    title: 'Tasks'
  },
]

const pages = [
  {
    path: '/admin/auth/login',
    render: <Login/>,
    title: 'Admin Login'
  },
  {
    path: '/admin/auth/register',
    render: <Register/>,
    title: 'Admin Register'
  },
  {
    path: '*',
    render: <NotFound/>,
    title: 'Not Found'
  }
]

export default function App() {
  return (
    <>
      <BrowserRouter>
        <AuthProvider>
          <Routes>
            { protectedPages.map((page, index) => (
              <Route element={<ProtectedPages/>}>
                <Route 
                  key={index}
                  path={page.path}
                  element={
                    <ProtectedLayout pageTitle={page.title}>
                      <Suspense fallback={<Loading/>} >
                        {page.render}
                      </Suspense>
                    </ProtectedLayout>
                  }
                />
              </Route>
            ))}
            { pages.map((page, index) => (
              <Route 
                key={index}
                path={page.path}
                element={
                  <Suspense fallback={<Loading/>}>
                    { page.render}
                  </Suspense>
                }
              />
            ))}
          </Routes>
        </AuthProvider>
      </BrowserRouter>
    </>
  )
}