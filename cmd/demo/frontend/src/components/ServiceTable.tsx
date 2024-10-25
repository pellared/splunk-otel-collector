import React, { useMemo, useState } from 'react';
import ServiceSecurityType from '../models/ServiceSecurityType';
import VulnerabilityType from '../models/VulnerabilityType';

interface TableProps {
  data: ServiceSecurityType[];
}

function ServiceTable({ data }: TableProps) {
  const [expandedRows, setExpandedRows] = useState<Set<number>>(new Set());
  const [sortConfig, setSortConfig] = useState<{ key: keyof ServiceSecurityType; direction: 'asc' | 'desc' } | null>(null);
  const [vulnSortConfig, setVulnSortConfig] = useState<{ [index: number]: { key: keyof VulnerabilityType; direction: 'asc' | 'desc' } }>({});

  const sortedData = useMemo(() => {
    const sorted = [...data];
    if (sortConfig) {
      sorted.sort((a, b) => {
        const aValue = a[sortConfig.key];
        const bValue = b[sortConfig.key];

        if (typeof aValue === 'string' && typeof bValue === 'string') {
          return sortConfig.direction === 'asc' ? aValue.localeCompare(bValue) : bValue.localeCompare(aValue);
        } else if (typeof aValue === 'number' && typeof bValue === 'number') {
          return sortConfig.direction === 'asc' ? aValue - bValue : bValue - aValue;
        } else if (Array.isArray(aValue) && Array.isArray(bValue)) {
          const aCount = aValue.length;
          const bCount = bValue.length; 
          return sortConfig.direction === 'asc' ? aCount - bCount : bCount - aCount;
        }
        return 0;
      });
    }

    sorted.forEach((item, index) => {
      const vulnConfig = vulnSortConfig[index];
      if (vulnConfig) {
        item.vulnerabilities.sort((a, b) => {
          const aValue = a[vulnConfig.key];
          const bValue = b[vulnConfig.key];

          if (typeof aValue === 'string' && typeof bValue === 'string') {
            return vulnConfig.direction === 'asc' ? aValue.localeCompare(bValue) : bValue.localeCompare(aValue);
          } else if (typeof aValue === 'number' && typeof bValue === 'number') {
            return vulnConfig.direction === 'asc' ? aValue - bValue : bValue - aValue;
          }
          return 0;
        });
      }
    });

    return sorted;
  }, [data, sortConfig, vulnSortConfig]);

  const handleExpandClick = (index: number) => {
    const updatedRows = new Set(expandedRows);
    if (updatedRows.has(index)) {
      updatedRows.delete(index);
    } else {
      updatedRows.add(index);
    }
    setExpandedRows(updatedRows);
  };

  const handleSort = (key: keyof ServiceSecurityType) => {
    let direction: 'asc' | 'desc' = 'asc';
    if (sortConfig?.key === key && sortConfig.direction === 'asc') {
      direction = 'desc';
    }
    setSortConfig({ key, direction });
  };

  const handleVulnSort = (index: number, key: keyof VulnerabilityType) => {
    const vulnConfig = vulnSortConfig[index] || { key, direction: 'asc' };
    let direction: 'asc' | 'desc' = 'asc';

    if (vulnConfig.key === key && vulnConfig.direction === 'asc') {
      direction = 'desc';
    }

    setVulnSortConfig(prevConfig => ({
      ...prevConfig,
      [index]: { key, direction }
    }));
  };

  const getRowBackgroundColor = (vulnCount: number) => {
    if (vulnCount === 0) return 'bg-white';
    if (vulnCount < 3) return 'bg-yellow-100';
    if (vulnCount < 6) return 'bg-yellow-300';
    return 'bg-red-400';
  };

  return (
    <div className="overflow-x-auto">
      <table className="min-w-full bg-white border border-gray-300">
        <thead>
          <tr className="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
            <th className="py-3 px-6 text-center border border-gray-300 cursor-pointer" onClick={() => handleSort('port')}>
              Port {sortConfig?.key === 'port' ? (sortConfig.direction === 'asc' ? '▲' : '▼') : '⯁'}
            </th>
            <th className="py-3 px-6 text-center border border-gray-300 cursor-pointer" onClick={() => handleSort('state')}>
              State {sortConfig?.key === 'state' ? (sortConfig.direction === 'asc' ? '▲' : '▼') : '⯁'}
            </th>
            <th className="py-3 px-6 text-center border border-gray-300 cursor-pointer" onClick={() => handleSort('host')}>
              Host {sortConfig?.key === 'host' ? (sortConfig.direction === 'asc' ? '▲' : '▼') : '⯁'}
            </th>
            <th className="py-3 px-6 text-center border border-gray-300 cursor-pointer" onClick={() => handleSort('service')}>
              Service {sortConfig?.key === 'service' ? (sortConfig.direction === 'asc' ? '▲' : '▼') : '⯁'}
            </th>
            <th className="py-3 px-6 text-center border border-gray-300 cursor-pointer" onClick={() => handleSort('version')}>
              Version {sortConfig?.key === 'version' ? (sortConfig.direction === 'asc' ? '▲' : '▼') : '⯁'}
            </th>
            <th className="py-3 px-6 text-center border border-gray-300 cursor-pointer" onClick={() => handleSort('vulnerabilities')}>
              Vulnerabilities {sortConfig?.key === 'vulnerabilities' ? (sortConfig.direction === 'asc' ? '▲' : '▼') : '⯁'}
            </th>
          </tr>
        </thead>
        <tbody>
          {sortedData.map((item: ServiceSecurityType, index: number) => {
            const isExpanded = expandedRows.has(index);
            const vulnCount = item.vulnerabilities.length;

            return (
              <React.Fragment key={index}>
                <tr className={`border-b border-gray-300 cursor-pointer ${getRowBackgroundColor(vulnCount)}`} onClick={() => vulnCount > 0 && handleExpandClick(index)}>
                  <td className="py-3 px-6 text-center border border-gray-300">{item.port}</td>
                  <td className="py-3 px-6 text-center border border-gray-300">{item.service}</td>
                  <td className="py-3 px-6 text-center border border-gray-300">{item.version}</td>
                  <td className="py-3 px-6 text-center border border-gray-300">{vulnCount > 0 ? vulnCount : 'No Vulnerabilities'}</td>
                </tr>
                {isExpanded && vulnCount > 0 && (
                  <tr className="border-b border-gray-300">
                    <td colSpan={4} className="p-4">
                      <table className="min-w-full bg-gray-100 border border-gray-300">
                        <thead>
                          <tr>
                            <th className="py-2 px-4 text-center border border-gray-300 cursor-pointer" onClick={() => handleVulnSort(index, 'CVE')}>
                              CVE {vulnSortConfig[index]?.key === 'CVE' ? (vulnSortConfig[index].direction === 'asc' ? '▲' : '▼') : '⯁'}
                            </th>
                            <th className="py-2 px-4 text-center border border-gray-300 cursor-pointer" onClick={() => handleVulnSort(index, 'ExploitID')}>
                              ExploitID {vulnSortConfig[index]?.key === 'ExploitID' ? (vulnSortConfig[index].direction === 'asc' ? '▲' : '▼') : '⯁'}
                            </th>
                            <th className="py-2 px-4 text-center border border-gray-300 cursor-pointer" onClick={() => handleVulnSort(index, 'URL')}>
                              URL {vulnSortConfig[index]?.key === 'URL' ? (vulnSortConfig[index].direction === 'asc' ? '▲' : '▼') : '⯁'}
                            </th>
                            <th className="py-2 px-4 text-center border border-gray-300 cursor-pointer" onClick={() => handleVulnSort(index, 'CVSS')}>
                              CVSS {vulnSortConfig[index]?.key === 'CVSS' ? (vulnSortConfig[index].direction === 'asc' ? '▲' : '▼') : '⯁'}
                            </th>
                          </tr>
                        </thead>
                        <tbody>
                          {item.vulnerabilities.map((vuln: VulnerabilityType, vulnIndex: number) => (
                            <tr key={vulnIndex} className="border-b border-gray-300">
                              <td className="py-2 px-4 text-center border border-gray-300">{vuln.CVE}</td>
                              <td className="py-2 px-4 text-center border border-gray-300">{vuln.ExploitID}</td>
                              <td className="py-2 px-4 text-center border border-gray-300">
                                <a href={vuln.URL} target="_blank" rel="noopener noreferrer" className="text-blue-600 underline">
                                  {vuln.URL}
                                </a>
                              </td>
                              <td className="py-2 px-4 text-center border border-gray-300">{vuln.CVSS}</td>
                            </tr>
                          ))}
                        </tbody>
                      </table>
                    </td>
                  </tr>
                )}
              </React.Fragment>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}

export default ServiceTable;