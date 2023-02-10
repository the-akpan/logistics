/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["index.html", "src/**/*.{js,jsx}"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Plus Jakarta Sans", "sans-serif"],
      },
      colors: {
        primary: "#4e148c",
        secondary: "#ff6201",
      },
    },
  },
  plugins: [],
};
