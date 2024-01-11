import { useContext } from "react";
import { Outlet, Navigate } from "react-router-dom";
import { AuthContext } from "./AuthProvider";
import Loading from "../Loading";


export default function ProtectedPages() {  
  const { user, loading } = useContext(AuthContext)!

  if (loading) {
    return <Loading/>
  }

  if (user) {
    return <Outlet/>
  }

  return <Navigate to="/admin/auth/login"/>
}