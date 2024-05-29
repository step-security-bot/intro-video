/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./internal/template/**/*.templ', './public/components/**/*.js'],
  theme: {
    extend: {
      fontFamily: {
        sans: ['var(--font-geist-sans)'],
        mono: ['var(--font-geist-mono)'],
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
