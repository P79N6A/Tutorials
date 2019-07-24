from flask import Flask, request
import json
app = Flask(__name__)


@app.route('/', methods=["POST"])
def hello_world():
    data = request.get_data()
    print(data)
    return json.dumps({"message": "success", "code": 200})

if __name__ == '__main__':
   app.run()
