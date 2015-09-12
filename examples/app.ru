require 'sinatra'

class Test < Sinatra::Base

  get "/" do
    "
      <h2>
        <a href='/hello'>Hello</a>
      </h2>

      pid: #{Process.pid}
    "
  end

  get '/hello' do
    sleep 0.5
    "Hello World!"
  end
end

run Test.new
