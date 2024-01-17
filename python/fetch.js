fetch(`http://54.255.34.229:8060/tasks/${uid}`)
.then(res => {
  return res.json()
})
.then(data => {
  setTask(data.data)
})