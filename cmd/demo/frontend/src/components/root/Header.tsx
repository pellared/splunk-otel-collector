import { BsFillLightningChargeFill } from "react-icons/bs";
import { FaSearch, FaPlus, FaRegBookmark} from "react-icons/fa";

function Header() {
    return (
      <header className="flex justify-between h-14 items-center bg-gray-100 border-b border-gray-300 w-auto">     
        {/* Center section */}
        <div className="text-left ml-5 h-14">
          <p className="text-gray-600">Data Management</p>
          <p className="text-lg text-gray-800">Service Discovery</p>
        </div>
        
        {/* Right section with button */}
        <div className="flex items-center h-14 w-4/6 space-x-4">
          <div className="flex-row border-l h-14 items-center border-gray-300 flex justify-between px-5">
          <p className="text-gray-600 w-full pr-2 flex flex-row whitespace-nowrap items-center font-semibold">
  <BsFillLightningChargeFill className="mr-2" /> 
  NEW! Splunk's new look
</p>
            <button className="bg-blue-500 text-white px-3 py-1 rounded-md flex-shrink-0">Learn more</button>
          </div>
          <div className="border-x px-5 h-14 items-center flex border-gray-300">
            <p className="text-gray-800 w-full whitespace-nowrap font-bold">2774 days left in trial</p>
          </div>
          <div className="flex text-gray-600 flex-row justify-around w-1/5 ">
          <FaSearch />
          <FaPlus />
          <FaRegBookmark />
          </div>
        </div>
      </header>
    );
  }
  
export default Header;