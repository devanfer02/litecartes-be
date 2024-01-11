import { ReactNode, useEffect } from "react"
import Sidebar from './Sidebar'

type LayoutProps = {
  children: ReactNode
  pageTitle: string 
}

export default function Layout({children, pageTitle}: LayoutProps) {
  useEffect(() => {
    document.title = pageTitle 
  })

  return (
    <>
      <Sidebar/>
      { children }
    </>
  )
}