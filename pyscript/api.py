import requests
import sys
import json

class APIClient:
    def __init__(self, base_url, api_key):
        self.base_url = base_url
        self.headers = {
            "Authorization": f"Bearer {api_key}",
            "Content-Type": "application/json"
        }

    def post(self, endpoint, payload):
        url = f"{self.base_url}/{endpoint}"
        response = requests.post(url, json=payload, headers=self.headers)
        return self._handle_response(response)

    def get(self, endpoint, params=None):
        url = f"{self.base_url}/{endpoint}"
        response = requests.get(url, headers=self.headers, params=params)
        return self._handle_response(response)

    def put(self, endpoint, payload):
        url = f"{self.base_url}/{endpoint}"
        response = requests.put(url, json=payload, headers=self.headers)
        return self._handle_response(response)

    def delete(self, endpoint):
        url = f"{self.base_url}/{endpoint}"
        response = requests.delete(url, headers=self.headers)
        return self._handle_response(response)

    def _handle_response(self, response):
        try:
            response.raise_for_status()
            return {
                "status": response.status_code,
                "data": response.json() if "application/json" in response.headers.get("Content-Type", "") else response.text
            }
        except requests.exceptions.RequestException as e:
            return {"error": str(e), "response": response.text}

def main():
    raw = sys.stdin.read()
    data = json.loads(raw)
    name = data.get("name")
    print(f"Hello, {name}!")
    # json.dumps(result) send to go

if __name__ == "__main__":
    BASE_URL = "https://api.example.com/v1"
    API_KEY = "YOUR_API_KEY_HERE"

    api = APIClient(BASE_URL, API_KEY)

    print("\nPOST")
    new_item = api.post("items", {
        "name": "Example Item",
        "price": 29.99
    })
    print(new_item)

    print("\nGET")
    fetched = api.get("items", params={
        "limit": 10,
        "sort": "price",
        "order": "desc"
    })
    print(fetched)

    print("\nPUT")
    updated = api.put("items/123", {
        "name": "Updated Name",
        "price": 39.99
    })
    print(updated)

    print("\nDELETE")
    deleted = api.delete("items/123")
    print(deleted)
