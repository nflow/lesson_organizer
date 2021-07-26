const colors = require("tailwindcss/colors");
const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  purge: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
  darkMode: false, // or 'media' or 'class'
  prefix: "tw-",
  theme: {
    colors: {
      transparent: "transparent",
      black: colors.black,
      white: colors.white,
      gray: colors.trueGray,
      current: "currentColor",
      "red-300": "#F94144",
      "red-400": "#F82529",
      orange: "#F3722C",
      "yellow-orange": "#F8961E",
      yellow: "#F9C74F",
      green: "#90BE6D",
      turquoise: "#43AA8B",
      blue: "#577590",
    },
    extend: {
      fontFamily: {
        logo: ["Permanent Marker", ...defaultTheme.fontFamily.sans],
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
