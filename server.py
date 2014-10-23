#!/usr/bin/env python
#coding: utf-8

import json
from datetime import datetime

from flask import Flask, request
from flask.ext.restful import Api, Resource
import psycopg2

app = Flask(__name__)

dbname = 'timehackerdb'
dbuser = 'postgres'
dbpasswd = 'postgres'
dbhost = 'localhost'
conn_args = "host='{}' dbname='{}' user='{}' password='{}'".format(dbhost, dbname, dbuser, dbpasswd)
g_conn = None

def get_conn():
    global g_conn
    if g_conn is None or g_conn.closed:
        g_conn = psycopg2.connect(conn_args)
    return g_conn

class Feedbacks(Resource):
    def get(self):
        conn = get_conn()
        cur = conn.cursor()
        cur.execute('select time, data from user_data')
        records = cur.fetchall()
        cur.close()
        return json.dumps([{"time": time.strftime("%Y-%m-%d %H:%M"), "data": data}
                            for time, data in records])


class Feedback(Resource):
    def post(self):
        conn = get_conn()
        cur = conn.cursor()
        cur.execute('insert into user_data(time, data) values(now(), %(data)s);',
                    {"data": request.form["data"]})
        conn.commit()
        cur.close()
        return 'ok'


api = Api(app)
api.add_resource(Feedbacks, '/feedbacks')
api.add_resource(Feedback, '/feedback')


if __name__ == '__main__':
    app.run(host='127.0.0.1', port=8002)

