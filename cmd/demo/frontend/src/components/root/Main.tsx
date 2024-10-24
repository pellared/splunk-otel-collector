import { ReactNode } from "react";
import PageList from "./PageList";

interface RootLayoutProps {
    children: ReactNode;
  }

function Main({children} : RootLayoutProps) {
    return ( 
        <main>
            <PageList />
            {children}
        </main>
     );
}

export default Main;