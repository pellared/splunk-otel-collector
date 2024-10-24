function PageList() {
    return (
      <div className="my-1 pl-6 flex-row flex border-b border-gray-300">
        <a className="block px-1 text-gray-800 font-medium my-2">Deployed integrations</a>
        <a className="block px-3 text-gray-800 font-medium my-2">Available integrations</a>
        <a className="block px-3 text-gray-800 font-medium my-2">Data tools</a>
        <a className="block px-4 text-gray-800 font-bold my-2">Service Discovery</a>
      </div>
    );
  }

export default PageList;