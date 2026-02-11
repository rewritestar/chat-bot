import "./globals.css";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ko">
      <body className="bg-blue-200 flex justify-center items-center">
        {children}
      </body>
    </html>
  );
}
