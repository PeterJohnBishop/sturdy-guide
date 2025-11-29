import aiohttp
import asyncio

class APIClientAsync:
    def __init__(self, base_url, api_key):
        self.base_url = base_url
        self.headers = {"Authorization": api_key}

    async def get(self, endpoint, params=None):
        async with aiohttp.ClientSession(headers=self.headers) as session:
            async with session.get(f"{self.base_url}/{endpoint}", params=params) as resp:
                return await resp.json()

    async def post(self, endpoint, data=None, json=None):
        async with aiohttp.ClientSession(headers=self.headers) as session:
            async with session.post(f"{self.base_url}/{endpoint}", data=data, json=json) as resp:
                return await resp.json()

    async def put(self, endpoint, json):
        async with aiohttp.ClientSession(headers=self.headers) as session:
            async with session.put(f"{self.base_url}/{endpoint}", json=json) as resp:
                return await resp.json()

    async def delete(self, endpoint):
        async with aiohttp.ClientSession(headers=self.headers) as session:
            async with session.delete(f"{self.base_url}/{endpoint}") as resp:
                return await resp.json()


async def main():
    api = APIClientAsync("https://api.example.com", "YOUR_API_KEY")
    data = await api.get("/api", params={"limit": 10})
    print(data)

if __name__ == "__main__":
    asyncio.run(main())