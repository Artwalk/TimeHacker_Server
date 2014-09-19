require 'sinatra'
require 'sinatra/json'
require 'json'
require 'pg'

set :port, 8001

conn = PGconn.open(:dbname => 'timehackerdb')

get '/feedbacks' do
  res = conn.query('SELECT * FROM user_data')

  feedbacks = {}
  for r in res
    feedbacks[r['time']] = JSON.parse(r['data'])
  end

  json feedbacks
end

post '/feedback' do
	data = params[:data]
  conn.exec('INSERT INTO user_data(time, data) VALUES(now(), \'%s\')' % data)
end
