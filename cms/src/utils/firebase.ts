import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

const firebaseConfig = {
  apiKey: "AIzaSyB-kslb-7-siOlOtV5ccfjiFMZx8keoxQk",
  authDomain: "litecartes.firebaseapp.com",
  projectId: "litecartes",
  storageBucket: "litecartes.appspot.com",
  messagingSenderId: "637638186888",
  appId: "1:637638186888:web:e36193db2bc0963ab30c3d",
  measurementId: "G-76Z6E4DYEE"
};

const app = initializeApp(firebaseConfig)
const auth = getAuth(app)

export { auth, app }
