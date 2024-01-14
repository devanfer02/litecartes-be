import { ChangeEvent } from "react"

type InputProps = {
  label: string 
  type: string 
  onChange: (e: ChangeEvent<HTMLInputElement>) => void 
  value: string | number
}

export default function Input({label, type, onChange, value}: InputProps) {
  return (
    <div className="mb-5">
      <label htmlFor="" className="text-xl font-semibold block">
        { label }
      </label>
      <input 
        type={type}
        onChange={onChange}
        value={value}
        className="border border-ltcbrown px-2 py-2 w-full mt-2 rounded-md focus:outline-none focus:border-sky-500"
      />
    </div>
  )
}