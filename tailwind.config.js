/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      height: {
        86: '86%'
      }
    }
  },
  plugins: [require('tailwind-scrollbar')]
}
