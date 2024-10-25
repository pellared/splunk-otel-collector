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
        CVE: "CVE-2019-10131",
        ExploitID: "EXP-2019-10131",
        URL: "https://nvd.nist.gov/vuln/detail/CVE-2019-10131",
        CVSS: "7.2"
      },
      {
        CVE: "CVE-2016-6663",
        ExploitID: "EXP-2016-6663",
        URL: "https://nvd.nist.gov/vuln/detail/CVE-2016-6663",
        CVSS: "6.8"
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
        CVE: "CVE-2023-27522",
        ExploitID: "EXP-2023-27522",
        URL: "https://nvd.nist.gov/vuln/detail/CVE-2023-27522",
        CVSS: "9.8"
      },
      {
        CVE: "CVE-2021-42013",
        ExploitID: "EXP-2021-42013",
        URL: "https://nvd.nist.gov/vuln/detail/CVE-2021-42013",
        CVSS: "7.5"
      }
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
