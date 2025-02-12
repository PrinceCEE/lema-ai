import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        black: "#181D27",
        lightblack: "#535862",
        tableArrow: "#717680",
        paginationBtnBg: "#F9F5FF",
        paginationBtnHoverText: "#7F56D9",
        "red-400": "#F9566A",
      },
      fontWeight: {
        semibold: "600",
        medium: "500",
        normal: "400",
      },
      fontSize: {
        xl: ["60px", "72px"],
        xs: ["12px", "18px"],
        sm: ["14px", "20px"],
        x: ["36px", "43.57px"],
      },
      screens: {
        mobilesm: { max: "639px" },
      },
    },
  },
  plugins: [],
} satisfies Config;
