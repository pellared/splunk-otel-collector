import { useState } from 'react';
import Obj from '../models/Obj.ts';
import Row from './Row';
import TableHeader from './TableHeader';

interface TableProps {
  data: Obj[];
}

function Table({ data }: TableProps) {
  const [sortConfig, setSortConfig] = useState<{ key: keyof Obj; direction: 'asc' | 'desc' } | null>(null);

  const sortedData = [...data];
  if (sortConfig !== null) {
    sortedData.sort((a, b) => {
      if (a[sortConfig.key] < b[sortConfig.key]) {
        return sortConfig.direction === 'asc' ? -1 : 1;
      }
      if (a[sortConfig.key] > b[sortConfig.key]) {
        return sortConfig.direction === 'asc' ? 1 : -1;
      }
      return 0;
    });
  }

  const handleSort = (key: keyof Obj) => {
    let direction: 'asc' | 'desc' = 'asc';
    if (sortConfig && sortConfig.key === key && sortConfig.direction === 'asc') {
      direction = 'desc';
    }
    setSortConfig({ key, direction });
  };

  return (
    <div className="overflow-x-auto">
      <table className="min-w-full bg-white border border-gray-300">
        <thead>
          <TableHeader sortConfig={sortConfig} handleSort={handleSort} />
        </thead>
        <tbody>
          {sortedData.map((item, index) => (
            <Row key={index} {...item} />
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default Table;