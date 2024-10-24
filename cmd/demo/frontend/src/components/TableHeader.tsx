import Obj from "../models/Obj";

interface TableHeaderProps {
  sortConfig: { key: keyof Obj; direction: 'asc' | 'desc' } | null;
  handleSort: (key: keyof Obj) => void;
}

function TableHeader({ sortConfig, handleSort }: TableHeaderProps) {
  const getSortDirection = (key: keyof Obj) => {
    if (!sortConfig || sortConfig.key !== key) return '⯁';
    return sortConfig.direction === 'asc' ? '▲' : '▼';
  };

  return (
    <tr className="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
      <th className="py-3 px-6 text-center border border-gray-300 cursor-pointer" onClick={() => handleSort('port')}>
        Port {getSortDirection('port')}
      </th>
      <th className="py-3 px-6 text-center border border-gray-300 cursor-pointer" onClick={() => handleSort('state')}>
        State {getSortDirection('state')}
      </th>
      <th className="py-3 px-6 text-center border border-gray-300 cursor-pointer" onClick={() => handleSort('service')}>
        Service {getSortDirection('service')}
      </th>
      <th className="py-3 px-6 text-center border border-gray-300 cursor-pointer" onClick={() => handleSort('version')}>
        Version {getSortDirection('version')}
      </th>
    </tr>
  );
}

export default TableHeader;
