import { ReactNode } from "react";
import PageList from "./PageList";

interface RootLayoutProps {
    children: ReactNode;
  }

function Main({children} : RootLayoutProps) {
    return ( 
    <main className="col-start-2 bg-white w-auto">
      <PageList />
      <div className="px-8">
        <h1 className="font-bold my-8">ServiceDiscovery</h1>
        {children}
      </div>
    </main>
     );
}

export default Main;