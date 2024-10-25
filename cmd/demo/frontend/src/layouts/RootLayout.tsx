import Header from "../components/root/Header";
import Nav from "../components/root/Nav";
import Main from "../components/root/Main";

function RootLayout() {

  return (
    <div className="grid grid-rows-[auto_1fr] grid-cols-[auto_1fr] min-h-screen overflow-hidden">
      <Nav />
      <Header />
      <Main />
    </div>
  );
}

export default RootLayout;
