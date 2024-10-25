import { useMemo, useState } from 'react';

interface TableProps<T> {
  data: T[];
}

function Table<T extends Record<string, any>>({ data }: TableProps<T>) {
  const [sortConfig, setSortConfig] = useState<{ key: keyof T; direction: 'asc' | 'desc' } | null>(null);

  const headers = useMemo(() => {
    if (data.length === 0) return [];
    return Object.keys(data[0]).map(key => ({
      key: key as keyof T,
      label: key.charAt(0).toUpperCase() + key.slice(1),
    }));
  }, [data]);

  const sortedData = [...data];
  if (sortConfig !== null) {
    sortedData.sort((a, b) => {
      if (a[sortConfig.key] < b[sortConfig.key]) return sortConfig.direction === 'asc' ? -1 : 1;
      if (a[sortConfig.key] > b[sortConfig.key]) return sortConfig.direction === 'asc' ? 1 : -1;
      return 0;
    });
  }

  const handleSort = (key: keyof T) => {
    let direction: 'asc' | 'desc' = 'asc';
    if (sortConfig && sortConfig.key === key && sortConfig.direction === 'asc') {
      direction = 'desc';
    }
    setSortConfig({ key, direction });
  };

  return (
    <div className="h-96 overflow-y-scroll">
      <table className="min-w-full bg-white border border-gray-300">
        <thead>
          <tr className="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
            {headers.map(header => (
              <th
                key={String(header.key)}
                className="py-3 px-6 text-center border border-gray-300 cursor-pointer"
                onClick={() => handleSort(header.key)}
              >
                {header.label} {sortConfig?.key === header.key ? (sortConfig.direction === 'asc' ? '▲' : '▼') : '⯁'}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {sortedData.map((item, index) => (
            <Row key={index} data={item} headers={headers} />
          ))}
        </tbody>
      </table>
    </div>
  );
}

interface RowProps<T> {
  data: T;
  headers: Array<{ key: keyof T; label: string }>;
}

function Row<T>({ data, headers }: RowProps<T>) {
  return (
    <tr className="border-b border-gray-300 hover:bg-gray-100">
      {headers.map(header => (
        <td key={String(header.key)} className="py-3 px-6 border border-gray-300 text-black">
          {renderCellContent(data[header.key])}
        </td>
      ))}
    </tr>
  );
}

function renderCellContent(value: any): React.ReactNode {
  if (Array.isArray(value)) {
    return value.join(', ');
  }
  if (typeof value === 'string' || typeof value === 'number') {
    return value;
  }
  return JSON.stringify(value);
}

export default Table;