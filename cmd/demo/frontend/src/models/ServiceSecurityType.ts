import VulnerabilityType from "./VulnerabilityType";

export default interface ServiceSecurityType{
    port :string;
    service :string;
    version :string;
    vulnerabilities :VulnerabilityType[];
}
