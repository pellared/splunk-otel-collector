import { NavLink } from "react-router-dom";

interface PageListProps {
  currentPath: string;
}

function PageList({ currentPath }: PageListProps) {
  return (
    <div className="my-1 pl-6 flex-row flex border-b border-gray-300">
      <NavLink to="/deployed-integrations" className={`block px-1 ${currentPath === '/deployed-integrations' ? 'font-bold' : 'text-gray-800 font-medium'} my-2`}>
        Deployed Integrations
      </NavLink>
      <NavLink to="/available-integrations" className={`block px-3 ${currentPath === '/available-integrations' ? 'font-bold' : 'text-gray-800 font-medium'} my-2`}>
        Available Integrations
      </NavLink>
      <NavLink to="/example" className={`block px-3 ${currentPath === '/data-tools' ? 'font-bold' : 'text-gray-800 font-medium'} my-2`}>
        Data Tools
      </NavLink>
      <NavLink to="/" className={`block px-4 ${currentPath === '/' ? 'font-bold' : 'text-gray-800 font-medium'} my-2`}>
        Service Discovery
      </NavLink>
      <NavLink to="/ServiceSecurity" className={`block px-3 ${currentPath === '/ServiceSecurity' ? 'font-bold' : 'text-gray-800 font-medium'} my-2`}>
        Service Security
      </NavLink>
    </div>
  );
}

export default PageList;