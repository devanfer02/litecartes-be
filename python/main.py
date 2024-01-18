from flask import Flask, jsonify, request
from vertexai.preview.generative_models import GenerativeModel
from vertexai.language_models import TextGenerationModel
from dotenv import load_dotenv
import os 
import requests
import json
import vertexai

os.environ["GOOGLE_APPLICATION_CREDENTIALS"] = 'config/litecartes-gcloud.json'
load_dotenv()

PROJECT_ID = os.getenv('PROJECT_ID')
LOCATION = os.getenv('LOCATION')
API_URL = os.getenv('GO_API_URL')
EXPECTED_ATTR = ['literacy', 'user_response']

vertexai.init(project=PROJECT_ID, location=LOCATION)

params = {
    "max_output_tokens": 512,
    "temperature": 0.9,
    "top_p": 1,
    "top_k": 40
}

model = TextGenerationModel.from_pretrained("text-bison@001")

app = Flask(__name__)

def get_json_literacy() :
    response = model.predict(f"""
        give me a short 1 paragraph literacy text with bahasa Indonesia with a title and topic 
        literacy about latest topic that gen z might be interested into and a question questioning 
        user about the literacy text and give 4 options to make user choose the correct one. 
        dont give options letter A, B, C, D,  just leave them as it is
        only one option is a correct choice and other options is wrong
        formatted like this in json format : 
                             
        title 
                             
        literacy

        question

        4 options                      
                             
        answer
    """, **params
    )

    try : 
        json_response = json.loads(response.text)

        return json_response
    except Exception as e:
        err_response = {
            "error": f"{e}",
            "result_prompt": response.text
        }
        return err_response

@app.route('/create/literacy', methods=['POST'])
def create_new_literacy():

    json_response = get_json_literacy()

    return json_response, 200

@app.route('/create/question/option', methods=['POST']) 
def create_option_question() :

    json_payload = get_json_literacy()

    if json_payload.get('error') is not None : 

        return json_payload, 500

    for index, value in enumerate(json_payload['options']) :
        if value == json_payload['answer'] :
            json_payload['answer'] = index

    json_payload['options'] = "|".join(json_payload['options'])
    json_payload['category_id'] = 'LTC-APP-generated1'
    json_payload['answer'] = str(json_payload['answer'])

    try :
        response = requests.post(f'{API_URL}/questions', json=json_payload)
        json_response = response.json()

        json_response['payload'] = json_payload 

        return json_response
    except Exception as e : 
        err_response = {
            "error": f"{e}",
            "payload": json_payload
        }

        return err_response, 400

@app.route('/uraian', methods=['POST'])
def evaluate_uraian():

    if not request.is_json :
        return jsonify({
            "error": "invalid request, expected json data"
        }), 400

    data = request.get_json()

    for attr in EXPECTED_ATTR :
        if not attr in data :
            return jsonify({
                "error": f"missing expected attribute: {attr}"
            }), 400

    literacy = data[EXPECTED_ATTR[0]]
    user_response = data[EXPECTED_ATTR[1]]

    response = model.predict(f"""
        give me rating from 1 to 10 with Paul-Elder Critical Thinking Framework of the 
        user response : {user_response} 
        to the literacy : {literacy}
        """
    , **params)

    return jsonify({
        "response": response.text
    }), 200

if __name__ == "__main__" :
    app.jinja_env.auto_reload = True
    app.run(debug=True, port=5000, use_reloader=True)