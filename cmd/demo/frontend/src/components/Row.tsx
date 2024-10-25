
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

export default Row;