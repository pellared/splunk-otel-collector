import { ReactNode } from "react";
import Header from "../components/root/Header";
import Nav from "../components/root/Nav";
import Main from "../components/root/Main";

interface RootLayoutProps {
  children: ReactNode;
}

function RootLayout({ children }: RootLayoutProps) {
  return (
    <div className="grid grid-rows-[auto_1fr] grid-cols-[auto_1fr] min-h-screen overflow-hidden">
      <Nav />
      <Header />
      <Main>{children}</Main>
    </div> 
  );
}

export default RootLayout;

