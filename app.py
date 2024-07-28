from flask import Flask, request, jsonify
import requests
import json

app = Flask(__name__)

LLM_API_URL = 'http://localhost:11434/api/chat'
MODEL_NAME = 'llama3'

def send_chat_to_llm(messages):
    data = {
        "model": MODEL_NAME,
        "messages": messages,
        "stream": True
    }
    response = requests.post(url=LLM_API_URL, json=data)
    response.raise_for_status()
    return response

def read_llm_response(response):
    output = ''
    for line in response.iter_lines():
        body = json.loads(line)
        if body.get("done") is False:
            message = body.get("message", "")
            content = message.get("content", "")
            output += content
            print(content, end="", flush=True)
        else:
            return output
    return output

@app.route('/chat', methods=['POST'])
def chat():
    data = request.json
    messages = data.get('messages', [])

    try:
        response = send_chat_to_llm(messages)
        ai_response = read_llm_response(response)
        messages.append({"role": "assistant", "content": ai_response})
        return jsonify({"response": ai_response})
    except Exception as e:
        return jsonify({"error": str(e)}), 500

@app.route('/')
def hello_world():
    return 'Hello, World!\n'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=9090)
