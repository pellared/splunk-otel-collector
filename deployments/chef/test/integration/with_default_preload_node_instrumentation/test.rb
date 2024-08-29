libsplunk_path = '/usr/lib/splunk-instrumentation/libsplunk.so'
node_options = '-r /usr/lib/splunk-instrumentation/splunk-otel-js/node_modules/@splunk/otel/instrument'
resource_attributes = 'splunk.zc.method=splunk-otel-auto-instrumentation-\d+\.\d+\.\d+'

describe package('splunk-otel-auto-instrumentation') do
  it { should be_installed }
end

describe npm('@splunk/otel', path: '/usr/lib/splunk-instrumentation/splunk-otel-js') do
  it { should be_installed }
end

describe file('/etc/ld.so.preload') do
  its('content') { should match /^#{libsplunk_path}$/ }
end

describe file('/usr/lib/systemd/system.conf.d/00-splunk-otel-auto-instrumentation.conf') do
  it { should_not exist }
end

describe file('/usr/lib/splunk-instrumentation/instrumentation.conf') do
  it { should_not exist }
end

describe file('/etc/splunk/zeroconfig/java.conf') do
  it { should_not exist }
end

describe file('/etc/splunk/zeroconfig/dotnet.conf') do
  it { should_not exist }
end

describe file('/etc/splunk/zeroconfig/node.conf') do
  its('content') { should match /^NODE_OPTIONS=#{node_options}$/ }
  its('content') { should match /^OTEL_RESOURCE_ATTRIBUTES=#{resource_attributes}$/ }
  its('content') { should_not match /.*OTEL_SERVICE_NAME.*/ }
  its('content') { should match /^SPLUNK_PROFILER_ENABLED=false$/ }
  its('content') { should match /^SPLUNK_PROFILER_MEMORY_ENABLED=false$/ }
  its('content') { should match /^SPLUNK_METRICS_ENABLED=false$/ }
  its('content') { should_not match /.*OTEL_EXPORTER_OTLP_ENDPOINT.*/ }
  its('content') { should_not match /.*OTEL_EXPORTER_OTLP_PROTOCOL.*/ }
  its('content') { should_not match /.*OTEL_METRICS_EXPORTER.*/ }
  its('content') { should_not match /.*OTEL_LOGS_EXPORTER.*/ }
end

describe service('splunk-otel-collector') do
  it { should be_enabled }
  it { should be_running }
end

describe service('td-agent') do
  it { should_not be_enabled }
  it { should_not be_running }
end
