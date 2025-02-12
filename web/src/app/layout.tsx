import { Inter } from "next/font/google";
import { QueryProvider } from "@/providers";
import "./globals.css";

const inter = Inter({
  weight: ["400", "500"],
  style: "normal",
  subsets: ["latin"],
  display: "swap",
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <div
          className={`w-full h-screen pt-[100px] px-8 flex justify-center ${inter.className}`}
        >
          <QueryProvider>{children}</QueryProvider>
        </div>
      </body>
    </html>
  );
}
