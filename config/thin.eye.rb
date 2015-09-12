RUBY   = 'ruby'
BUNDLE = 'bundle'

path = File.expand_path "../../", __FILE__
PATH = path

Eye.load "../apps/thin_farm.rb"

Eye.config do
  # logger "/tmp/eye.log"
  logger "#{path}/tmp/eye.log"
end
Eye.app 'thin-farm' do
  working_dir "#{PATH}/processes"
  env "RACK_ENV" => "production"

  stop_on_delete true # this option means, when we change pids and load config,
                      # deleted processes will be stops

  trigger :flapping, :times => 3, :within => 20.seconds
  check :memory, :below => 60.megabytes, :every => 30.seconds, :times => 5
  start_timeout 30.seconds

  group :web do
    chain action: :restart, grace: 3.seconds
    chain action: :start,   grace: 1.seconds

    (5555..5560).each do |port|
      thin self, port
    end
  end

end


# Eye.app 'go-loadbal' do
#   working_dir "#{PATH}/processes"
#
#   group :reverse_proxy do
#     chain action: :restart, grace: 1.seconds
#     chain action: :start,   grace: 0.5.seconds
#
#     # go_loadbal self, 8000
#   end
# end
