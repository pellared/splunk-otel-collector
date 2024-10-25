import { pageTitles } from '../../data/PageTitles';
import PageList from './PageList';
import { Outlet, useLocation } from 'react-router-dom';

function Main() {
  const location = useLocation();
  const currentPath = location.pathname;

  return ( 
    <main className="col-start-2 bg-white w-auto">
      <PageList currentPath={currentPath} />
      <div className="px-8">
        <h1 className="font-bold my-8">{pageTitles[currentPath] || 'Service Discovery'}</h1>
        <Outlet />
      </div>
    </main>
  );
}

export default Main;