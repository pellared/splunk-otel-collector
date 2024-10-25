import ServiceSecurityType from "../models/ServiceSecurityType";

const empty: ServiceSecurityType[] = [
    {
        port: "",
        state: "",
        host: "",
        service: "",
        version: "",
      vulnerabilities: []
    },
]

export default empty;