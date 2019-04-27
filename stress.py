# modified fetch function with semaphore
import json
import random
import asyncio
from aiohttp import ClientSession

d = {
        "name": "FElipe",
        "email": "fefeef@gmail.com",
        "poll": {
            "id": 2
        },
        "nominate": {
            "id": 46
        }
    }

async def fetch(url, session):
    async with session.post(url, data=json.dumps(d)) as response:
        print("SEND")
        return await response.read()


async def bound_fetch(sem, url, session):
    # Getter function with semaphore.
    async with sem:
        await fetch(url, session)


async def run(r):
    url = "http://localhost:3000/poll-api/polls/1/vote"
    tasks = []
    # create instance of Semaphore
    sem = asyncio.Semaphore(1000)

    # Create client session that will ensure we dont open new connection
    # per each request.
    async with ClientSession() as session:
        for i in range(r):
            # pass Semaphore and session to every POST request
            task = asyncio.ensure_future(bound_fetch(sem, url.format(i), session))
            tasks.append(task)

        responses = asyncio.gather(*tasks)
        await responses

number = 1000000
loop = asyncio.get_event_loop()

future = asyncio.ensure_future(run(number))
loop.run_until_complete(future)
