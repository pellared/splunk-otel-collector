import VulnerabilityType from "./VulnerabilityType";

export default interface ServiceSecurityType{
    port: string;
    state: string;
    host: string;
    service: string;
    version: string;
    vulnerabilities :VulnerabilityType[];
}
