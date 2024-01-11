import axios from "axios"
import { useEffect, useState } from "react"

interface Cursor {
  prev_cursor: string 
  next_cursor: string 
}

export default function useFetch<T>(apiUrl: string): [
  T[],
  ()=> Promise<void>,
  unknown,
  boolean,
  Cursor|null 
] {
  const [ cursor, setCursor ] = useState<Cursor | null>(null)
  const [ data, setData ] = useState<T[]>([])
  const [ error, setError ] = useState<unknown>(null)
  const [ loading, setLoading ] = useState(true)

  const getData = async() => {
    try {
      const res = await axios.get(apiUrl)

      setData(res.data.data)
      setLoading(false)
      setCursor(res.data.pagination)
    } catch (e) {
      setError(e)
      setLoading(false)
      console.log(e)
    }
  }

  useEffect(() => {
    getData()
  }, [apiUrl])

  return [ data, getData, error, loading, cursor]
}