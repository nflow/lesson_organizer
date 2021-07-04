const colors = require("tailwindcss/colors");

module.exports = {
  purge: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  darkMode: false, // or 'media' or 'class'
  theme: {
    colors: {
      transparent: "transparent",
      black: colors.black,
      white: colors.white,
      gray: colors.trueGray,
      current: "currentColor",
      red: "#FF595E",
      orange: "#FFCA3A",
      green: "#8AC926",
      blue: "#1982C4",
      purple: "#6A4C93",
    },
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
