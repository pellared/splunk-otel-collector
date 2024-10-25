import ServiceDiscoveryType from "../models/ServiceDiscoveryType.ts"

const example: ServiceDiscoveryType[] = [
    {
      port: "3306/tcp",
      service: "mysql",
      version: "MySQL 5.5.64-MariaDB-1~trusty"
    },
    {
      port: "8080/tcp",
      service: "http",
      version: "Apache httpd 2.4.57 ((Debian))"
    },
    {
        port: "8081/tcp",
        service: "http",
        version: "Apache httpd 2.4.62 ((Debian))"
    }
  ]

export default example;