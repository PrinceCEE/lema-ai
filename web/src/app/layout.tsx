import { Inter } from "next/font/google";
import { AppStoreProvider, QueryProvider } from "@/providers";
import "./globals.css";
import { ToastContainer } from "@/components";

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
          className={`w-full h-screen md:pt-[100px] pt-[50px] px-8 flex justify-center ${inter.className}`}
        >
          <AppStoreProvider>
            <ToastContainer />
            <QueryProvider>{children}</QueryProvider>
          </AppStoreProvider>
        </div>
      </body>
    </html>
  );
}
