/** @type {import('tailwindcss').Config} */

const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  content: ["./static/**/*.{css,js}", "./views/**/*.html"],
  theme: {
    extend: {
	  fontFamily: {
        sans: ['Inter var', ...defaultTheme.fontFamily.sans],
		orbitron: ['Orbitron', 'sans-serif'],
      },
	},
  },
  plugins: [],
}
