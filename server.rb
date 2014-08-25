require 'sinatra'

get '/' do 
	"<p>This is for TimeHacker Server!<p>%s" % readData
end

post '/feedback' do
	data = params[:data]
	saveData(data)
end

$fName = "post.txt"

def saveData(data)
	# coding: utf-8 
	f = File.open($fName, "a")
	f.puts data
	f.close
end

def readData
	File.open($fName, "r") { |f|
		s = f.read
	}
	
end