import { ReactNode } from "react";

interface RootLayoutProps {
  children: ReactNode;
}

function RootLayout({ children }: RootLayoutProps) {
  return (
    <>
      {children}
    </>
  );
}

export default RootLayout;
