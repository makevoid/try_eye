
def thin(proxy, port)
  name = "thin-#{port}"

  # "-c #{proxy.working_dir}",

  proxy.process(name) do
    pid_file "#{name}.pid"

    start_command "#{BUNDLE} exec rackup -O trying_eye -p #{port} -P #{name}.pid"
    stop_signals [:QUIT, 2.seconds, :TERM, 1.seconds, :KILL]

    stdall "thin.stdall.log"

    check :http, :url => "http://127.0.0.1:#{port}/hello", :pattern => /World/,
                 :every => 5.seconds, :times => [2, 3], :timeout => 1.second
  end
end

def loadbal(ports)
  start_command "#{BUNDLE} exec rackup -O trying_eye -p #{port} -P #{name}.pid"
end
