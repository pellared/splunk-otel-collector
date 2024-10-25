import ServiceDiscovery from "../pages/ServiceDiscovery"
import ServiceSecurity from "../pages/ServiceSecurity"

export const menuItems = [
    {
      path: "/",
      element: <ServiceDiscovery />,
    },
    {
      path: "/ServiceSecurity",
      element: <ServiceSecurity />,
    },
  ];