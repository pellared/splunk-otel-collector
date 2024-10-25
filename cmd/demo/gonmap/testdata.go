package main

const example = `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE nmaprun>
<?xml-stylesheet href="file:///usr/bin/../share/nmap/nmap.xsl" type="text/xsl"?>
<!-- Nmap 7.80 scan initiated Wed Oct 23 21:03:12 2024 as: nmap -sV -&#45;script default,vuln,brute -oX - localhost -->
<nmaprun scanner="nmap" args="nmap -sV -&#45;script default,vuln,brute -oX - localhost" start="1729710192" startstr="Wed Oct 23 21:03:12 2024" version="7.80" xmloutputversion="1.04">
<scaninfo type="connect" protocol="tcp" numservices="1000" services="1,3-4,6-7,9,13,17,19-26,30,32-33,37,42-43,49,53,70,79-85,88-90,99-100,106,109-111,113,119,125,135,139,143-144,146,161,163,179,199,211-212,222,254-256,259,264,280,301,306,311,340,366,389,406-407,416-417,425,427,443-445,458,464-465,481,497,500,512-515,524,541,543-545,548,554-555,563,587,593,616-617,625,631,636,646,648,666-668,683,687,691,700,705,711,714,720,722,726,749,765,777,783,787,800-801,808,843,873,880,888,898,900-903,911-912,981,987,990,992-993,995,999-1002,1007,1009-1011,1021-1100,1102,1104-1108,1110-1114,1117,1119,1121-1124,1126,1130-1132,1137-1138,1141,1145,1147-1149,1151-1152,1154,1163-1166,1169,1174-1175,1183,1185-1187,1192,1198-1199,1201,1213,1216-1218,1233-1234,1236,1244,1247-1248,1259,1271-1272,1277,1287,1296,1300-1301,1309-1311,1322,1328,1334,1352,1417,1433-1434,1443,1455,1461,1494,1500-1501,1503,1521,1524,1533,1556,1580,1583,1594,1600,1641,1658,1666,1687-1688,1700,1717-1721,1723,1755,1761,1782-1783,1801,1805,1812,1839-1840,1862-1864,1875,1900,1914,1935,1947,1971-1972,1974,1984,1998-2010,2013,2020-2022,2030,2033-2035,2038,2040-2043,2045-2049,2065,2068,2099-2100,2103,2105-2107,2111,2119,2121,2126,2135,2144,2160-2161,2170,2179,2190-2191,2196,2200,2222,2251,2260,2288,2301,2323,2366,2381-2383,2393-2394,2399,2401,2492,2500,2522,2525,2557,2601-2602,2604-2605,2607-2608,2638,2701-2702,2710,2717-2718,2725,2800,2809,2811,2869,2875,2909-2910,2920,2967-2968,2998,3000-3001,3003,3005-3007,3011,3013,3017,3030-3031,3052,3071,3077,3128,3168,3211,3221,3260-3261,3268-3269,3283,3300-3301,3306,3322-3325,3333,3351,3367,3369-3372,3389-3390,3404,3476,3493,3517,3527,3546,3551,3580,3659,3689-3690,3703,3737,3766,3784,3800-3801,3809,3814,3826-3828,3851,3869,3871,3878,3880,3889,3905,3914,3918,3920,3945,3971,3986,3995,3998,4000-4006,4045,4111,4125-4126,4129,4224,4242,4279,4321,4343,4443-4446,4449,4550,4567,4662,4848,4899-4900,4998,5000-5004,5009,5030,5033,5050-5051,5054,5060-5061,5080,5087,5100-5102,5120,5190,5200,5214,5221-5222,5225-5226,5269,5280,5298,5357,5405,5414,5431-5432,5440,5500,5510,5544,5550,5555,5560,5566,5631,5633,5666,5678-5679,5718,5730,5800-5802,5810-5811,5815,5822,5825,5850,5859,5862,5877,5900-5904,5906-5907,5910-5911,5915,5922,5925,5950,5952,5959-5963,5987-5989,5998-6007,6009,6025,6059,6100-6101,6106,6112,6123,6129,6156,6346,6389,6502,6510,6543,6547,6565-6567,6580,6646,6666-6669,6689,6692,6699,6779,6788-6789,6792,6839,6881,6901,6969,7000-7002,7004,7007,7019,7025,7070,7100,7103,7106,7200-7201,7402,7435,7443,7496,7512,7625,7627,7676,7741,7777-7778,7800,7911,7920-7921,7937-7938,7999-8002,8007-8011,8021-8022,8031,8042,8045,8080-8090,8093,8099-8100,8180-8181,8192-8194,8200,8222,8254,8290-8292,8300,8333,8383,8400,8402,8443,8500,8600,8649,8651-8652,8654,8701,8800,8873,8888,8899,8994,9000-9003,9009-9011,9040,9050,9071,9080-9081,9090-9091,9099-9103,9110-9111,9200,9207,9220,9290,9415,9418,9485,9500,9502-9503,9535,9575,9593-9595,9618,9666,9876-9878,9898,9900,9917,9929,9943-9944,9968,9998-10004,10009-10010,10012,10024-10025,10082,10180,10215,10243,10566,10616-10617,10621,10626,10628-10629,10778,11110-11111,11967,12000,12174,12265,12345,13456,13722,13782-13783,14000,14238,14441-14442,15000,15002-15004,15660,15742,16000-16001,16012,16016,16018,16080,16113,16992-16993,17877,17988,18040,18101,18988,19101,19283,19315,19350,19780,19801,19842,20000,20005,20031,20221-20222,20828,21571,22939,23502,24444,24800,25734-25735,26214,27000,27352-27353,27355-27356,27715,28201,30000,30718,30951,31038,31337,32768-32785,33354,33899,34571-34573,35500,38292,40193,40911,41511,42510,44176,44442-44443,44501,45100,48080,49152-49161,49163,49165,49167,49175-49176,49400,49999-50003,50006,50300,50389,50500,50636,50800,51103,51493,52673,52822,52848,52869,54045,54328,55055-55056,55555,55600,56737-56738,57294,57797,58080,60020,60443,61532,61900,62078,63331,64623,64680,65000,65129,65389"/>
<verbose level="0"/>
<debugging level="0"/>
<host starttime="1729710203" endtime="1729710234"><status state="up" reason="conn-refused" reason_ttl="0"/>
<address addr="127.0.0.1" addrtype="ipv4"/>
<hostnames>
<hostname name="localhost" type="user"/>
<hostname name="localhost" type="PTR"/>
</hostnames>
<ports><extraports state="closed" count="997">
<extrareasons reason="conn-refused" count="997"/>
</extraports>
<port protocol="tcp" portid="3306"><state state="open" reason="syn-ack" reason_ttl="0"/><service name="mysql" product="MySQL" version="5.5.64-MariaDB-1~trusty" method="probed" conf="10"><cpe>cpe:/a:mysql:mysql:5.5.64-mariadb-1~trusty</cpe></service><script id="clamav-exec" output="ERROR: Script execution failed (use -d to debug)"/><script id="mysql-brute" output="&#xa;  Accounts: No valid accounts found&#xa;  Statistics: Performed 0 guesses in 1 seconds, average tps: 0.0&#xa;  ERROR: The service seems to have failed or is heavily firewalled..."><elem key="Accounts">No valid accounts found</elem>
<elem key="Statistics">Performed 0 guesses in 1 seconds, average tps: 0.0</elem>
<elem key="ERROR">The service seems to have failed or is heavily firewalled...</elem>
</script><script id="mysql-enum" output="&#xa;  Valid usernames: &#xa;    root:&lt;empty&gt; - Valid credentials&#xa;    netadmin:&lt;empty&gt; - Valid credentials&#xa;    guest:&lt;empty&gt; - Valid credentials&#xa;    user:&lt;empty&gt; - Valid credentials&#xa;    web:&lt;empty&gt; - Valid credentials&#xa;    sysadmin:&lt;empty&gt; - Valid credentials&#xa;    administrator:&lt;empty&gt; - Valid credentials&#xa;    webadmin:&lt;empty&gt; - Valid credentials&#xa;    admin:&lt;empty&gt; - Valid credentials&#xa;    test:&lt;empty&gt; - Valid credentials&#xa;  Statistics: Performed 10 guesses in 5 seconds, average tps: 2.0"><table key="Valid usernames">
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">root</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">netadmin</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">guest</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">user</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">web</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">sysadmin</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">administrator</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">webadmin</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">admin</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">test</elem>
</table>
</table>
<elem key="Statistics">Performed 10 guesses in 5 seconds, average tps: 2.0</elem>
</script><script id="mysql-info" output="&#xa;  Protocol: 10&#xa;  Version: 5.5.64-MariaDB-1~trusty&#xa;  Thread ID: 98055&#xa;  Capabilities flags: 63487&#xa;  Some Capabilities: Speaks41ProtocolOld, Support41Auth, ConnectWithDatabase, SupportsTransactions, IgnoreSigpipes, InteractiveClient, Speaks41ProtocolNew, DontAllowDatabaseTableColumn, SupportsLoadDataLocal, IgnoreSpaceBeforeParenthesis, FoundRows, ODBCClient, LongPassword, LongColumnFlag, SupportsCompression, SupportsAuthPlugins, SupportsMultipleResults, SupportsMultipleStatments&#xa;  Status: Autocommit&#xa;  Salt: etWs\ToT5JUJVPn[\;&#xa;  Auth Plugin Name: mysql_native_password"><elem key="Protocol">10</elem>
<elem key="Version">5.5.64-MariaDB-1~trusty</elem>
<elem key="Thread ID">98055</elem>
<elem key="Capabilities flags">63487</elem>
<table key="Some Capabilities">
<elem>Speaks41ProtocolOld</elem>
<elem>Support41Auth</elem>
<elem>ConnectWithDatabase</elem>
<elem>SupportsTransactions</elem>
<elem>IgnoreSigpipes</elem>
<elem>InteractiveClient</elem>
<elem>Speaks41ProtocolNew</elem>
<elem>DontAllowDatabaseTableColumn</elem>
<elem>SupportsLoadDataLocal</elem>
<elem>IgnoreSpaceBeforeParenthesis</elem>
<elem>FoundRows</elem>
<elem>ODBCClient</elem>
<elem>LongPassword</elem>
<elem>LongColumnFlag</elem>
<elem>SupportsCompression</elem>
<elem>SupportsAuthPlugins</elem>
<elem>SupportsMultipleResults</elem>
<elem>SupportsMultipleStatments</elem>
</table>
<elem key="Status">Autocommit</elem>
<elem key="Salt">et:Ws\ToT5JUJVPn[\;</elem>
<elem key="Auth Plugin Name">mysql_native_password</elem>
</script></port>
<port protocol="tcp" portid="8080"><state state="open" reason="syn-ack" reason_ttl="0"/><service name="http" product="Apache httpd" version="2.4.57" extrainfo="(Debian)" method="probed" conf="10"><cpe>cpe:/a:apache:http_server:2.4.57</cpe></service><script id="citrix-brute-xml" output="FAILED: No domain specified (use ntdomain argument)"/><script id="clamav-exec" output="ERROR: Script execution failed (use -d to debug)"/><script id="http-brute" output="  &#xa;  Path &quot;/&quot; does not require authentication"/><script id="http-csrf" output="&#xa;Spidering limited to: maxdepth=3; maxpagecount=20; withinhost=localhost&#xa;  Found the following possible CSRF vulnerabilities: &#xa;    &#xa;    Path: http://localhost:8080/&#xa;    Form id: &#xa;    Form action: add_task.php&#xa;"/><script id="http-dombased-xss" output="Couldn&apos;t find any DOM based XSS."/><script id="http-enum" output="&#xa;  /phpinfo.php: Possible information file&#xa;"/><script id="http-open-proxy" output="Proxy might be redirecting requests"/><script id="http-server-header" output="Apache/2.4.57 (Debian)"><elem>Apache/2.4.57 (Debian)</elem>
</script><script id="http-stored-xss" output="Couldn&apos;t find any stored XSS vulnerabilities."/><script id="http-title" output="Todo List"><elem key="title">Todo List</elem>
</script><script id="http-vuln-cve2017-1001000" output="ERROR: Script execution failed (use -d to debug)"/><script id="vulners" output="&#xa;  cpe:/a:apache:http_server:2.4.57: &#xa;    &#x9;95499236-C9FE-56A6-9D7D-E943A24B633A&#x9;10.0&#x9;https://vulners.com/githubexploit/95499236-C9FE-56A6-9D7D-E943A24B633A&#x9;*EXPLOIT*&#xa;    &#x9;2C119FFA-ECE0-5E14-A4A4-354A2C38071A&#x9;10.0&#x9;https://vulners.com/githubexploit/2C119FFA-ECE0-5E14-A4A4-354A2C38071A&#x9;*EXPLOIT*&#xa;    &#x9;CVE-2024-38476&#x9;9.8&#x9;https://vulners.com/cve/CVE-2024-38476&#xa;    &#x9;CVE-2024-38474&#x9;9.8&#x9;https://vulners.com/cve/CVE-2024-38474&#xa;    &#x9;A5425A79-9D81-513A-9CC5-549D6321897C&#x9;9.8&#x9;https://vulners.com/githubexploit/A5425A79-9D81-513A-9CC5-549D6321897C&#x9;*EXPLOIT*&#xa;    &#x9;CVE-2024-38475&#x9;9.1&#x9;https://vulners.com/cve/CVE-2024-38475&#xa;    &#x9;0486EBEE-F207-570A-9AD8-33269E72220A&#x9;9.1&#x9;https://vulners.com/githubexploit/0486EBEE-F207-570A-9AD8-33269E72220A&#x9;*EXPLOIT*&#xa;    &#x9;B0A9E5E8-7CCC-5984-9922-A89F11D6BF38&#x9;8.2&#x9;https://vulners.com/githubexploit/B0A9E5E8-7CCC-5984-9922-A89F11D6BF38&#x9;*EXPLOIT*&#xa;    &#x9;CVE-2024-38473&#x9;8.1&#x9;https://vulners.com/cve/CVE-2024-38473&#xa;    &#x9;249A954E-0189-5182-AE95-31C866A057E1&#x9;8.1&#x9;https://vulners.com/githubexploit/249A954E-0189-5182-AE95-31C866A057E1&#x9;*EXPLOIT*&#xa;    &#x9;23079A70-8B37-56D2-9D37-F638EBF7F8B5&#x9;8.1&#x9;https://vulners.com/githubexploit/23079A70-8B37-56D2-9D37-F638EBF7F8B5&#x9;*EXPLOIT*&#xa;    &#x9;F7F6E599-CEF4-5E03-8E10-FE18C4101E38&#x9;7.5&#x9;https://vulners.com/githubexploit/F7F6E599-CEF4-5E03-8E10-FE18C4101E38&#x9;*EXPLOIT*&#xa;    &#x9;E5C174E5-D6E8-56E0-8403-D287DE52EB3F&#x9;7.5&#x9;https://vulners.com/githubexploit/E5C174E5-D6E8-56E0-8403-D287DE52EB3F&#x9;*EXPLOIT*&#xa;    &#x9;DB6E1BBD-08B1-574D-A351-7D6BB9898A4A&#x9;7.5&#x9;https://vulners.com/githubexploit/DB6E1BBD-08B1-574D-A351-7D6BB9898A4A&#x9;*EXPLOIT*&#xa;    &#x9;CVE-2024-40898&#x9;7.5&#x9;https://vulners.com/cve/CVE-2024-40898&#xa;    &#x9;CVE-2024-39573&#x9;7.5&#x9;https://vulners.com/cve/CVE-2024-39573&#xa;    &#x9;CVE-2024-38477&#x9;7.5&#x9;https://vulners.com/cve/CVE-2024-38477&#xa;    &#x9;CVE-2024-38472&#x9;7.5&#x9;https://vulners.com/cve/CVE-2024-38472&#xa;    &#x9;CVE-2024-27316&#x9;7.5&#x9;https://vulners.com/cve/CVE-2024-27316&#xa;    &#x9;CVE-2023-43622&#x9;7.5&#x9;https://vulners.com/cve/CVE-2023-43622&#xa;    &#x9;CVE-2023-31122&#x9;7.5&#x9;https://vulners.com/cve/CVE-2023-31122&#xa;    &#x9;CDC791CD-A414-5ABE-A897-7CFA3C2D3D29&#x9;7.5&#x9;https://vulners.com/githubexploit/CDC791CD-A414-5ABE-A897-7CFA3C2D3D29&#x9;*EXPLOIT*&#xa;    &#x9;C9A1C0C1-B6E3-5955-A4F1-DEA0E505B14B&#x9;7.5&#x9;https://vulners.com/githubexploit/C9A1C0C1-B6E3-5955-A4F1-DEA0E505B14B&#x9;*EXPLOIT*&#xa;    &#x9;BD3652A9-D066-57BA-9943-4E34970463B9&#x9;7.5&#x9;https://vulners.com/githubexploit/BD3652A9-D066-57BA-9943-4E34970463B9&#x9;*EXPLOIT*&#xa;    &#x9;B5E74010-A082-5ECE-AB37-623A5B33FE7D&#x9;7.5&#x9;https://vulners.com/githubexploit/B5E74010-A082-5ECE-AB37-623A5B33FE7D&#x9;*EXPLOIT*&#xa;    &#x9;B0208442-6E17-5772-B12D-B5BE30FA5540&#x9;7.5&#x9;https://vulners.com/githubexploit/B0208442-6E17-5772-B12D-B5BE30FA5540&#x9;*EXPLOIT*&#xa;    &#x9;A820A056-9F91-5059-B0BC-8D92C7A31A52&#x9;7.5&#x9;https://vulners.com/githubexploit/A820A056-9F91-5059-B0BC-8D92C7A31A52&#x9;*EXPLOIT*&#xa;    &#x9;9814661A-35A4-5DB7-BB25-A1040F365C81&#x9;7.5&#x9;https://vulners.com/githubexploit/9814661A-35A4-5DB7-BB25-A1040F365C81&#x9;*EXPLOIT*&#xa;    &#x9;5A864BCC-B490-5532-83AB-2E4109BB3C31&#x9;7.5&#x9;https://vulners.com/githubexploit/5A864BCC-B490-5532-83AB-2E4109BB3C31&#x9;*EXPLOIT*&#xa;    &#x9;45D138AD-BEC6-552A-91EA-8816914CA7F4&#x9;7.5&#x9;https://vulners.com/githubexploit/45D138AD-BEC6-552A-91EA-8816914CA7F4&#x9;*EXPLOIT*&#xa;    &#x9;17C6AD2A-8469-56C8-BBBE-1764D0DF1680&#x9;7.5&#x9;https://vulners.com/githubexploit/17C6AD2A-8469-56C8-BBBE-1764D0DF1680&#x9;*EXPLOIT*&#xa;    &#x9;CVE-2023-45802&#x9;5.9&#x9;https://vulners.com/cve/CVE-2023-45802&#xa;    &#x9;CVE-2024-39884&#x9;0.0&#x9;https://vulners.com/cve/CVE-2024-39884&#xa;    &#x9;CVE-2024-36387&#x9;0.0&#x9;https://vulners.com/cve/CVE-2024-36387&#xa;    &#x9;CVE-2024-24795&#x9;0.0&#x9;https://vulners.com/cve/CVE-2024-24795&#xa;    &#x9;CVE-2023-38709&#x9;0.0&#x9;https://vulners.com/cve/CVE-2023-38709"><table key="cpe:/a:apache:http_server:2.4.57">
<table>
<elem key="type">githubexploit</elem>
<elem key="id">95499236-C9FE-56A6-9D7D-E943A24B633A</elem>
<elem key="cvss">10.0</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">2C119FFA-ECE0-5E14-A4A4-354A2C38071A</elem>
<elem key="cvss">10.0</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-38476</elem>
<elem key="cvss">9.8</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-38474</elem>
<elem key="cvss">9.8</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">A5425A79-9D81-513A-9CC5-549D6321897C</elem>
<elem key="cvss">9.8</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-38475</elem>
<elem key="cvss">9.1</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">0486EBEE-F207-570A-9AD8-33269E72220A</elem>
<elem key="cvss">9.1</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">B0A9E5E8-7CCC-5984-9922-A89F11D6BF38</elem>
<elem key="cvss">8.2</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-38473</elem>
<elem key="cvss">8.1</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">249A954E-0189-5182-AE95-31C866A057E1</elem>
<elem key="cvss">8.1</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">23079A70-8B37-56D2-9D37-F638EBF7F8B5</elem>
<elem key="cvss">8.1</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">F7F6E599-CEF4-5E03-8E10-FE18C4101E38</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">E5C174E5-D6E8-56E0-8403-D287DE52EB3F</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">DB6E1BBD-08B1-574D-A351-7D6BB9898A4A</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-40898</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-39573</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-38477</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-38472</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-27316</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2023-43622</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2023-31122</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">CDC791CD-A414-5ABE-A897-7CFA3C2D3D29</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">C9A1C0C1-B6E3-5955-A4F1-DEA0E505B14B</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">BD3652A9-D066-57BA-9943-4E34970463B9</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">B5E74010-A082-5ECE-AB37-623A5B33FE7D</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">B0208442-6E17-5772-B12D-B5BE30FA5540</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">A820A056-9F91-5059-B0BC-8D92C7A31A52</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">9814661A-35A4-5DB7-BB25-A1040F365C81</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">5A864BCC-B490-5532-83AB-2E4109BB3C31</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">45D138AD-BEC6-552A-91EA-8816914CA7F4</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">githubexploit</elem>
<elem key="id">17C6AD2A-8469-56C8-BBBE-1764D0DF1680</elem>
<elem key="cvss">7.5</elem>
<elem key="is_exploit">true</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2023-45802</elem>
<elem key="cvss">5.9</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-39884</elem>
<elem key="cvss">0.0</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-36387</elem>
<elem key="cvss">0.0</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2024-24795</elem>
<elem key="cvss">0.0</elem>
<elem key="is_exploit">false</elem>
</table>
<table>
<elem key="type">cve</elem>
<elem key="id">CVE-2023-38709</elem>
<elem key="cvss">0.0</elem>
<elem key="is_exploit">false</elem>
</table>
</table>
</script></port>
<port protocol="tcp" portid="8081"><state state="open" reason="syn-ack" reason_ttl="0"/><service name="http" product="Apache httpd" version="2.4.62" extrainfo="(Debian)" method="probed" conf="10"><cpe>cpe:/a:apache:http_server:2.4.62</cpe></service><script id="clamav-exec" output="ERROR: Script execution failed (use -d to debug)"/><script id="http-brute" output="  &#xa;  Path &quot;/&quot; does not require authentication"/><script id="http-csrf" output="&#xa;Spidering limited to: maxdepth=3; maxpagecount=20; withinhost=localhost&#xa;  Found the following possible CSRF vulnerabilities: &#xa;    &#xa;    Path: http://localhost:8081/doc/html/index.html&#xa;    Form id: &#xa;    Form action: search.html&#xa;"/><script id="http-dombased-xss" output="Couldn&apos;t find any DOM based XSS."/><script id="http-enum" output="&#xa;  /robots.txt: Robots file&#xa;  /README: Interesting, a readme.&#xa;"/><script id="http-fileupload-exploiter" output="&#xa;  &#xa;    Couldn&apos;t find a file-type field.&#xa;  &#xa;    Couldn&apos;t find a file-type field."><table>
<elem>Couldn&apos;t find a file-type field.</elem>
</table>
<table>
<elem>Couldn&apos;t find a file-type field.</elem>
</table>
</script><script id="http-form-brute" output="&#xa;  Accounts: No valid accounts found&#xa;  Statistics: Performed 24 guesses in 2 seconds, average tps: 12.0&#xa;  ERROR: The service seems to have failed or is heavily firewalled..."><elem key="Accounts">No valid accounts found</elem>
<elem key="Statistics">Performed 24 guesses in 2 seconds, average tps: 12.0</elem>
<elem key="ERROR">The service seems to have failed or is heavily firewalled...</elem>
</script><script id="http-robots.txt" output="1 disallowed entry &#xa;/"/><script id="http-server-header" output="Apache/2.4.62 (Debian)"><elem>Apache/2.4.62 (Debian)</elem>
</script><script id="http-sql-injection" output="&#xa;  Possible sqli for queries:&#xa;    http://localhost:8081/js/dist/error_report.js?v=5.2.1%27%20OR%20sqlspider&#xa;    http://localhost:8081/js/dist/ajax.js?v=5.2.1%27%20OR%20sqlspider&#xa;    http://localhost:8081/js/messages.php?v=5.2.1%27%20OR%20sqlspider&amp;lang=en&amp;l=en&#xa;    http://localhost:8081/js/messages.php?v=5.2.1&amp;lang=en%27%20OR%20sqlspider&amp;l=en&#xa;    http://localhost:8081/js/messages.php?v=5.2.1&amp;lang=en&amp;l=en%27%20OR%20sqlspider&#xa;    http://localhost:8081/js/vendor/tracekit.js?v=5.2.1%27%20OR%20sqlspider&#xa;    http://localhost:8081/js/dist/functions.js?v=5.2.1%27%20OR%20sqlspider&#xa;    http://localhost:8081/js/dist/error_report.js?v=5.2.1%27%20OR%20sqlspider&#xa;    http://localhost:8081/js/dist/ajax.js?v=5.2.1%27%20OR%20sqlspider&#xa;    http://localhost:8081/js/messages.php?v=5.2.1%27%20OR%20sqlspider&amp;lang=en&amp;l=en&#xa;    http://localhost:8081/js/messages.php?v=5.2.1&amp;lang=en%27%20OR%20sqlspider&amp;l=en&#xa;    http://localhost:8081/js/messages.php?v=5.2.1&amp;lang=en&amp;l=en%27%20OR%20sqlspider&#xa;    http://localhost:8081/js/vendor/tracekit.js?v=5.2.1%27%20OR%20sqlspider&#xa;    http://localhost:8081/js/dist/functions.js?v=5.2.1%27%20OR%20sqlspider&#xa;"/><script id="http-stored-xss" output="Couldn&apos;t find any stored XSS vulnerabilities."/><script id="http-title" output="phpMyAdmin"><elem key="title">phpMyAdmin</elem>
</script><script id="http-vuln-cve2017-1001000" output="ERROR: Script execution failed (use -d to debug)"/></port>
</ports>
<times srtt="60" rttvar="51" to="100000"/>
</host>
<postscript><script id="creds-summary" output="&#xa;  127.0.0.1: &#xa;    3306/mysql: &#xa;      root:&lt;empty&gt; - Valid credentials&#xa;      netadmin:&lt;empty&gt; - Valid credentials&#xa;      guest:&lt;empty&gt; - Valid credentials&#xa;      user:&lt;empty&gt; - Valid credentials&#xa;      web:&lt;empty&gt; - Valid credentials&#xa;      sysadmin:&lt;empty&gt; - Valid credentials&#xa;      administrator:&lt;empty&gt; - Valid credentials&#xa;      webadmin:&lt;empty&gt; - Valid credentials&#xa;      admin:&lt;empty&gt; - Valid credentials&#xa;      test:&lt;empty&gt; - Valid credentials"><table key="127.0.0.1">
<table key="3306/mysql">
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">root</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">netadmin</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">guest</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">user</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">web</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">sysadmin</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">administrator</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">webadmin</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">admin</elem>
</table>
<table>
<elem key="password">&lt;empty&gt;</elem>
<elem key="state">Valid credentials</elem>
<elem key="username">test</elem>
</table>
</table>
</table>
</script></postscript><runstats><finished time="1729710234" timestr="Wed Oct 23 21:03:54 2024" elapsed="41.76" summary="Nmap done at Wed Oct 23 21:03:54 2024; 1 IP address (1 host up) scanned in 41.76 seconds" exit="success"/><hosts up="1" down="0" total="1"/>
</runstats>
</nmaprun>
`
