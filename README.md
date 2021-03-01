## Running

`docker compose build && docker compose up`

request rewards from reward service:

`curl -i -XPOST localhost:1337/rewards -d '{"campaign_id": 1}'`

request arm from bandit service:

`curl -XPOST localhost:1338/randomize -d '{"unit": "visitor_id:12345", "context": {"campaign_id": 1}}'`