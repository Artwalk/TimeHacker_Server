require 'sinatra'
require 'json'
require 'fileutils'
#require 'time'

get '/' do 
	content_type :json
  dict = readData
  return dict
end

post '/feedback' do
	data = params[:data]
	saveData("{\"%s\" : %s}," % [Time.now, data])
end

$fName = "post.txt"


# r w file
def saveData(data)
	# coding: utf-8
	f = File.open($fName, "a")
	f.puts data
	f.close
end

def readData
	File.open($fName, "r") do |f|
		s = "[ %s ]" % f.read().slice!(0..-3)
	end
end