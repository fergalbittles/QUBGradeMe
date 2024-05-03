from flask import Flask
from flask import request
from flask import Response
import json
import sum

app = Flask(__name__)

@app.route('/')

def calculateTotal():
    marks = {}

    marks["Mark 1"] = request.args.get('mark_1')
    marks["Mark 2"] = request.args.get('mark_2')
    marks["Mark 3"] = request.args.get('mark_3')
    marks["Mark 4"] = request.args.get('mark_4')
    marks["Mark 5"] = request.args.get('mark_5')

    # Missing parameter error
    for key, value in marks.items():
        if not value:
            msg = key + " value is missing"
            return errorResponse(msg, 400)

    # Invalid input error
    for key, value in marks.items():
        try:
            marks[key] = int(value)
        except ValueError:
            msg = "You must provide a valid integer for " + key 
            return errorResponse(msg, 400)
        
        if marks[key] < 0:
            msg = "You must provide a non-negative integer for " + key
            return errorResponse(msg, 400)

        if marks[key] > 100:
            msg = "You cannot exceed 100 marks for " + key
            return errorResponse(msg, 400)

    # Perform calculation
    result = sum.totalMarks(marks)

    # Ensure the calculation was successful
    if result == False:
        msg = "Error occured while calculating total marks, ensure that valid input was entered"
        return errorResponse(msg, 400)

    # Return response
    r = {
        "error": False,
        "string": "Total Marks Acquired = " + str(result),
        "result": result,
    }
    reply = json.dumps(r)

    response = Response(response=reply, status=200, mimetype='application/json')
    response.headers["Content-Type"]="application/json"
    response.headers["Access-Control-Allow-Origin"]="*"

    return response


def errorResponse(msg, statusCode):
    r = {
            "error": True,
            "string": msg,
    }
    reply = json.dumps(r)

    response = Response(response=reply, status=statusCode, mimetype='application/json')
    response.headers["Content-Type"]="application/json"
    response.headers["Access-Control-Allow-Origin"]="*"

    return response


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)