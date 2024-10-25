import './App.css'
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { menuItems } from "./data/menuItems.tsx";
import RootLayout from "./layouts/RootLayout.tsx";

function App() {
  const router = createBrowserRouter([
    {
      path: '/',
      element: <RootLayout />,
      children: menuItems,
    },
  ]);

  return <RouterProvider router={router} />;
}

export default App;