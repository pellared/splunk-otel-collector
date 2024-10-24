import './App.css'
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { menuItems } from "./data/menuItems.tsx";
import RootLayout from "./layouts/RootLayout.tsx";

function App() {

  const router = createBrowserRouter(menuItems);

  return (
    <>
      <RootLayout>
        <RouterProvider router={router}></RouterProvider>
      </RootLayout>
    </>
  )
}

export default App
