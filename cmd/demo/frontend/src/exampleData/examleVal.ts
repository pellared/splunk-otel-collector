import ServiceSecurityType from "../models/ServiceSecurityType.ts";

const exampleVal: ServiceSecurityType[] = [
    {
      port: "3306/tcp",
      state: "open",
      service: "mysql",
      version: "MySQL 5.5.64-MariaDB-1~trusty",
      vulnerabilities: ["123", "qwe", "aaa"]
    },
    {
      port: "8080/tcp",
      state: "open",
      service: "http",
      version: "Apache httpd 2.4.57 ((Debian))",
      vulnerabilities: ["123", "qwe", "aaa"]
    },
    {
        port: "8081/tcp",
        state: "open",
        service: "http",
        version: "Apache httpd 2.4.62 ((Debian))",
        vulnerabilities: []
    }
  ]

export default exampleVal;