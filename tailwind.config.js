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
      red: "#F94144",
      orange: "#F3722C",
      yellow_orange: "#F8961E",
      yellow: "#F9C74F",
      green: "#90BE6D",
      turquoise: "#43AA8B",
      blue: "#577590",
    },
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
