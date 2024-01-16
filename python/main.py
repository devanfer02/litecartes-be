from flask import Flask, jsonify, request 
from vertexai.preview.generative_models import GenerativeModel
from vertexai.language_models import TextGenerationModel
from dotenv import load_dotenv
import os 
import vertexai

os.environ["GOOGLE_APPLICATION_CREDENTIALS"] = 'config/litecartes-gcloud.json'
load_dotenv()

PROJECT_ID = os.getenv('PROJECT_ID')
LOCATION = os.getenv('LOCATION')
CODE_CHAT_MODEL = os.getenv('CODE_CHAT_MODEL')
EXPECTED_ATTR = ['literacy', 'user_response']

app = Flask(__name__)

vertexai.init(project=PROJECT_ID, location=LOCATION)

params = {
    "max_output_tokens": 512,
    "temperature": 0.9,
    "top_p": 1,
    "top_k": 40
}

model = TextGenerationModel.from_pretrained("text-bison@001")

@app.route('/uraian', methods=['POST'])
def get_uraian_response():

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

    response = model.predict(f"give me rating from 1 to 10 of the user response : {user_response} to the literacy : {literacy}", **params)

    return jsonify({
        "response": response.text
    }), 200

if __name__ == "__main__" :
    app.run()