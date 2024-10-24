import { GoHome } from "react-icons/go";
import { FaSitemap } from "react-icons/fa";
import { IoNewspaperOutline } from "react-icons/io5";
import { RiRobot2Line } from "react-icons/ri";
import { MdNotificationsNone } from "react-icons/md";
import { CiGrid41 } from "react-icons/ci";
import { FaRuler } from "react-icons/fa6";
import { FaDatabase } from "react-icons/fa";
import { IoSettingsOutline } from "react-icons/io5";
import { HiIdentification } from "react-icons/hi2";
import { BiShapeSquare } from "react-icons/bi";
import { IoIosArrowForward } from "react-icons/io";
import { MdKeyboardDoubleArrowLeft } from "react-icons/md";

function Nav() {
    return (
        <nav className="row-span-2 w-64 bg-gray-50 h-screen flex flex-col justify-start border-r border-gray-300">
            {/* Left section with logo */}
            <div className="flex items-center space-x-2 border-b border-gray-300">
                {/* Logo placeholder */}
                <img src="logo.png" alt="Logo" className="w-14 h-14" />
                <p className="text-gray-900">GDI Cloud Integrations Sandbox</p>
            </div>
            <div>
                {/* First block of links */}
                <div className="space-y-2">
                    <a className="text-gray-800 font-medium ml-5 mt-5 flex items-center"> <GoHome /> Home</a>
                    <a className="text-gray-800 font-medium m-5 flex items-center"> <BiShapeSquare /> APM</a>
                    <a className="text-gray-800 font-medium m-5 flex items-center"> <FaSitemap /> Infrastructure</a>
                    <a className="text-gray-800 font-medium m-5 flex items-center"> <IoNewspaperOutline /> Log Observer</a>
                    <a className="text-gray-800 font-medium m-5 flex items-center"> <HiIdentification /> RUM</a>
                    <a className="text-gray-800 font-medium m-5 flex items-center"> <RiRobot2Line />Synthetics</a>
                </div>
                <hr className="mt-5"/>
                {/* Second block of links */}
                <div className="mt-8 space-y-2">
                    <a className="text-gray-800 font-medium ml-5 flex items-center"> <MdNotificationsNone />Detectors & SLOs</a>
                    <a className="text-gray-800 font-medium m-5 flex items-center"> <CiGrid41 /> Dashboards</a>
                    <a className="text-gray-800 font-medium m-5 flex items-center"> <FaRuler /> Metric Finder</a>
                    <a className="text-gray-800 font-medium m-5 flex items-center"> <FaDatabase /> Data Management</a>
                </div>
                <hr className="mt-5"/>
            </div>

            {/* Settings link */}
            <div>
                <a className="text-gray-800 font-medium m-5 flex items-center"> <IoSettingsOutline /> Settings <IoIosArrowForward /></a>
            </div>
            {/* bottom arrow */}
            <hr className="mt-32"/>
            <div>
                <a className="text-gray-800 font-medium m-5 flex items-center"> <MdKeyboardDoubleArrowLeft /></a>
            </div>
        </nav>
    );
}

export default Nav;