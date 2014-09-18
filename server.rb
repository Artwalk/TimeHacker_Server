require 'sinatra'
require 'sinatra/json'
require 'json'
require 'pg'

conn = PGconn.open(:dbname => 'TimeHackerDB')

get '/' do
  res = conn.query('SELECT * FROM user_data')

  allFeedbacks = {}
  for r in res
    allFeedbacks[r['time']] = JSON.parse(r['data'])
  end

  json allFeedbacks
end

post '/feedback' do
	data = params[:data]
  conn.exec('INSERT INTO user_data(time, data) VALUES(now(), \'%s\')' % data)
end
