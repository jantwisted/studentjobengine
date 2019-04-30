from flask import Flask
from flask import Flask, flash, redirect, render_template, request, session, abort
import os
import requests
import json
import jwt

app = Flask(__name__)
 
@app.route('/')
def home():
    if not session.get('logged_in'):
        return render_template('login.html')
    else:
        return "Hello Boss!  <a href='/logout'>Logout</a>"
 
@app.route('/login', methods=['POST'])
def do_admin_login():
    user_details = {}
    user_name = request.form['username']
    user_password = request.form['password']
    user_details['UserName'] = user_name
    user_details['Password'] = user_password
    url = 'http://127.0.0.1:8080/users/login'
    body = user_details
    headers = {'content-type': 'application/json'}
    response = requests.post(url, data=json.dumps(body), headers=headers)
    response_dict = json.loads(response.text)
    try:
        claim = jwt.decode(response_dict['JWT'], algorithms=['HS256'], verify=False)
        if claim['user'] == user_name:
            session['logged_in'] = True
        else:
            flash('wrong password!')
    except jwt.DecodeError:
        flash('wrong credentials!')
    return home()

@app.route("/logout")
def logout():
    session['logged_in'] = False
    return home()

 
if __name__ == "__main__":
    app.secret_key = os.urandom(12)
    app.run(debug=True,host='0.0.0.0', port=4000)
