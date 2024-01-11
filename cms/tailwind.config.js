/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'ltccrem': '#FFDBB9',
        'ltcbrown': '#8C3310'
      }
    },
  },
  plugins: [],
  
}

