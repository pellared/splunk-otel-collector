import { ReactNode } from "react";
import Header from "../components/root/Header";
import Nav from "../components/root/Nav";
import Main from "../components/root/main";

interface RootLayoutProps {
  children: ReactNode;
}

function RootLayout({ children }: RootLayoutProps) {
  return (
    <>
    <Header />
    <Nav />
    <Main>
    { children }
    </Main>
    </>
  );
}

export default RootLayout;
