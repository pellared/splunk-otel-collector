import ServiceSecurityType from "../models/ServiceSecurityType.ts";

export const exampleVal: ServiceSecurityType[] = [
  {
    port: "3306/tcp",
    state: "open",
    host: "localhost",
    service: "mysql",
    version: "MySQL 5.5.64-MariaDB-1~trusty",
    vulnerabilities: [
      {
        cve: "CVE-2019-10131",
        exploit_id: "EXP-2019-10131",
        url: "https://nvd.nist.gov/vuln/detail/CVE-2019-10131",
        cvss: "7.2"
      },
      {
        cve: "CVE-2016-6663",
        exploit_id: "EXP-2016-6663",
        url: "https://nvd.nist.gov/vuln/detail/CVE-2016-6663",
        cvss: "6.8"
      }
    ]
  },
  {
    port: "8080/tcp",
    state: "open",
    host: "localhost",
    service: "http",
    version: "Apache httpd 2.4.57 ((Debian))",
    vulnerabilities: [
      {
        cve: "CVE-2023-27522",
        exploit_id: "EXP-2023-27522",
        url: "https://nvd.nist.gov/vuln/detail/CVE-2023-27522",
        cvss: "9.8"
      },
      {
        cve: "CVE-2023-27522",
        exploit_id: "EXP-2023-27522",
        url: "https://nvd.nist.gov/vuln/detail/CVE-2023-27522",
        cvss: "9.8"
      },
    ]
  },
  {
    port: "8081/tcp",
    state: "open",
    host: "localhost",
    service: "http",
    version: "Apache httpd 2.4.62 ((Debian))",
    vulnerabilities: []
  }
];

export default exampleVal;
