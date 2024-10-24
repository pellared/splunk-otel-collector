import Obj from "../models/Obj.ts"

function Row({ port, state, service, version }: Obj) {
  return (
    <tr className="border-b border-gray-300 hover:bg-gray-100">
      <td className="py-3 px-6 border border-gray-300 text-black">{port}</td>
      <td className="py-3 px-6 border border-gray-300 text-black">{state}</td>
      <td className="py-3 px-6 border border-gray-300 text-black">{service}</td>
      <td className="py-3 px-6 border border-gray-300 text-black">{version}</td>
    </tr>
  );
}

export default Row;