type Props = {
  message: string 
}

export default function FlashError({ message }: Props) {
  return (
    <div className="bg-red-600 p-5 text-center">  
      <h1 className="text-white uppercase text-xl font-bold">
        Error Sending Request 
      </h1>
      <p className="text-white text-lg font-semibold">
        Error Message : { message }
      </p>
    </div>
  )
}