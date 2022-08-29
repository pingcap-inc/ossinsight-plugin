# OSSInsight Plugin

## Mechanism

1. Pull GitHub event by loop.
2. Filter the duplicate event.
3. Product event messages to `Pulsar`.
4. Consume event messages from `Plusar`.
5. Provide WebSocket API to the front end.

## Usage

1. Generate a new read-only token at [GitHub](https://github.com/settings/tokens).
2. Build binary file: `make build` (local) or `make build-server` (amd64 x86 linux).
3. Create a free pulsar cluster at [StreamNative Cloud](https://console.streamnative.cloud/organizations), and download key file.
4. Configure this file, and save to one of those path `./config.yaml`, `./config/config.yaml` or `../config/config.yaml` (relative path by binary file):

    ```yaml
    api:
        version: 2

    server:
        port: 6000 # HTTP port
        health: "/health" # health check api name
        syncEvent: "/sync-event" # event sync api name

    log:
        level: debug
        file: test.log
        format: json # json/text

    # redis 7+
    redis:
        host: localhost:6379
        password: ""
        db: 0

    pulsar:
        host: pulsar+ssl://XXX.XXX.XXX:XXXX # pulsar host
        audience: XXX:XX:XXXXX:XXXXX:XXXX # pulsar audience
        keypath: /XXX/XXX/XXX.json # private key path, download form StreamNative cloud
        producer:
            topic: XXXXX # send topic name
            retry: 3 # send message retry times
        consumer:
            topic: XXXXX # consume topic name
            name: events-consumer # consumer name
            concurrency: 5 # consumer numbers in a single ossinsight-plugin

    github:
        loop:
            timeout: 5000 # github events fetch loop timeout (ms)
            break: 1000 # github events fetch loop break (ms)
        tokens:
            - ghp_XXXXXXXXXXXXXXXXXXXXXXXXX # github read-only token
   
    interval:
        yearCount: 1m # year count fetch interval
        dayCount: 1m # year count fetch interval
        daily: 1m
        retry: 3
        retryWait: 10000 # ms

    tidb:
        host: localhost # tidb host
        port: 4000 # tidb port
        user: root # tidb user
        password: "" # tidb password
        db: gharchive_dev # tidb database
        sql: # some sql
        eventsDaily: |
            SELECT event_day, COUNT(*) AS events
            FROM github_events ge
            WHERE
            type = 'PullRequestEvent'
            AND event_year = YEAR(NOW())
            GROUP BY event_day
            ORDER BY event_day;
        prToday: |
            SELECT action, COUNT(action) as event_num
            FROM github_events
            WHERE type = 'PullRequestEvent'
            AND ((action = 'closed' AND pr_merged = true) OR action = 'opened')
            AND event_day = DATE(NOW())
            GROUP BY action, event_day;
        prDeveloperToday: |
            SELECT COUNT(DISTINCT actor_id) as actor_num
            FROM github_events
            WHERE type = 'PullRequestEvent'
            AND ((action = 'closed' AND pr_merged = true) OR action = 'opened')
            AND event_day = DATE(NOW())
        prThisYear: |
            SELECT action, COUNT(action) as event_num
            FROM github_events
            WHERE type = 'PullRequestEvent'
            AND ((action = 'closed' AND pr_merged = true) OR action = 'opened')
            AND event_year = YEAR(NOW())
            GROUP BY action
        prDeveloperThisYear: |
            SELECT COUNT(DISTINCT actor_id) as actor_num
            FROM github_events
            WHERE type = 'PullRequestEvent'
            AND ((action = 'closed' AND pr_merged = true) OR action = 'opened')
            AND event_year = YEAR(NOW())
    ```

5. Start by `./ossinsight-plugin`.

## API

### API Using Step

- `Client` start to connect the API Endpoint by `RAW WebSocket` protocol.
- `Server` will waiting `Client` to send params.
- `Client` send params (API has different params struct).
- `Server` will send a initial message by ***ONLY ONCE***.
- `Server` will endless send message to `Client` until connection closed.

### Initial Message

The initial message will has `firstMessageTag` tag, and it will be true.

E.g.:

```json
{
    "firstMessageTag": true,
    "apiVersion": 2,
    "eventMap": {
        "2022-01-01": "34",
        "2022-01-02": "41",
        "2022-01-03": "100",
        "2022-01-04": "322",
        "2022-01-05": "326",
        "2022-01-06": "391",
        "2022-01-07": "279",
        "2022-01-08": "120",
        "2022-01-09": "39",
        "2022-01-10": "331",
        "2022-01-11": "355"
    },
    "yearCountMap": {
        "dev": "2091",
        "merge": "28225",
        "open": "34078"
    },
    "dayCountMap": {
        "dev": "2",
        "merge": "65",
        "open": "227"
    }
}
```

### Sampling

Sampling message and return to client.

- Endpoint: `/sampling`, e.g.: `wss://localhost:6000/sampling`.
- Params:

    | Name | Required | Type | Description |
    | :- | :- | :- | :- |
    | `samplingRate` | Yes | int | Sampling rate. It means that N events are received that satisfy the conditions but only one of them is returned to the front end. If you want all of them, you need set it to `1`. |
    | `eventType` | No | string | Specify the event type you want to see. If you don't set it, all event types will be returned. |
    | `repoName` | No | string | Specify the repo name you want to see. If you don't set it, all repo names will be returned. |
    | `userName` | No | string | Specify the user name you want to see. If you don't set it, all user names will be returned. |
    | `filter` | No | string list | Specify the fields you want to see. If you don't set it, all fields will be returned. (This param will change result struct)|
    
    E.g.:

    ```json
    {
        "samplingRate": 3,
        "eventType": "PushEvent",
        "repoName": "pingcap-inc/ossinsight-plugin",
        "userName": "Icemap"
    }
    ```

- Result (the `{event entity}` is same as [/events](https://docs.github.com/en/rest/activity/events) API):

    ```json
    {
        "event": {event entity},
        "payload": {
            "devDay": 0,
            "devYear": 0,
            "merge": 0,
            "open": 0,
            "pr": 0
        }
    }
    ```

    E.g.:

    ```json
    {
        "event": {
            "type": "DeleteEvent",
            "public": true,
            "payload": {
                "ref": "dependabot/npm_and_yarn/netlify-cli-11.3.0",
                "ref_type": "branch",
                "pusher_type": "user"
            },
            "repo": {
                "id": 526513549,
                "name": "davy39/ntn-boilerplate",
                "url": "https://api.github.com/repos/davy39/ntn-boilerplate"
            },
            "actor": {
                "login": "dependabot[bot]",
                "id": 49699333,
                "avatar_url": "https://avatars.githubusercontent.com/u/49699333?",
                "gravatar_id": "",
                "url": "https://api.github.com/users/dependabot[bot]"
            },
            "created_at": "2022-08-29T02:06:15Z",
            "id": "23683284276"
        },
        "payload": {
            "devDay": 0,
            "devYear": 0,
            "merge": 0,
            "open": 0,
            "pr": 0
        }

    }
    ```

- With filter result will return a map, such as:
    
    Params:

    ```json
    {
        "samplingRate": 1,
        "filter": [
            "event.id", "event.type", "event.actor.login", "event.actor.avatar_url", "event.payload.push_id", "payload.merge"
        ]
    }
    ```

    Result:

    ```json
    {
        "event.actor.avatar_url": "https://avatars.githubusercontent.com/u/30280995?",
        "event.actor.login": "alanhaledc",
        "event.id": "23683417557",
        "event.payload.push_id": 10857043797,
        "event.type": "PushEvent",
        "payload.merge": 0
    }
    ```

- If you want a list:

    Params:

    ```json
    {
        "samplingRate": 1,
        "filter": [
            "event.id", "event.type", "event.actor.login", "event.actor.avatar_url", "event.payload.push_id", "payload.merge"
        ],
        "returnType": "list"
    }
    ```

    Result:

    ```json
    ["23683424409","PullRequestEvent","dependabot[bot]","https://avatars.githubusercontent.com/u/49699333?",null,0]
    ```

### Loop

Return to client by fixed time.

- Endpoint: `/loop`, e.g.: `wss://localhost:6000/loop`.
- Params:

    | Name | Required | Type | Description |
    | :- | :- | :- | :- |
    | `loopTime` | Yes | int | Fixed loop time. `Server` will send message to `Client` by per loop time. Unit is millsecond, and must be larger than `500`. |
    | `eventType` | No | string | Specify the event type you want to see. If you don't set it, all event types will be returned. |
    | `repoName` | No | string | Specify the repo name you want to see. If you don't set it, all repo names will be returned. |
    | `userName` | No | string | Specify the user name you want to see. If you don't set it, all user names will be returned. |
    | `detail` | No | bool | Whether send event details to `Client`, default is `false`. |

    E.g.:

    ```json
    {
        "loopTime": 1000,
        "eventType": "PushEvent",
        "repoName": "pingcap-inc/ossinsight-plugin",
        "userName": "Icemap",
        "detail": false
    }
    ```

- Result (the `{event entity}` is same as [/events](https://docs.github.com/en/rest/activity/events) API):

    ```json
    {
        "msgList": [{
                "event": {event entity},
                "payload": {
                    "devDay": 0,
                    "devYear": 0,
                    "merge": 0,
                    "open": 0,
                    "pr": 0
                }
            }, {
                "event": {event entity},
                "payload": {
                    "devDay": 0,
                    "devYear": 0,
                    "merge": 0,
                    "open": 0,
                    "pr": 0
                }
            }
        ],
        "typeMap": {
            "{event type}": {event num}
        }
    }
    ```

    E.g.:

    ```json
    {
        "msgList": [
            {
                "event": {
                    "type": "PushEvent",
                    "public": true,
                    "payload": {
                        "push_id": 10857075261,
                        "size": 3,
                        "distinct_size": 1,
                        "ref": "refs/heads/dependabot/npm_and_yarn/typescript-4.8.2",
                        "head": "cbc96e9571671ea0f1db18915b6e10c5263ad39d",
                        "before": "55366fe5ae3befe1e8ee0ebd8d698fb3b600c7d2",
                        "commits": [
                            {
                                "sha": "d57bad671ad2bbc41f36207063585caca1c5d00e",
                                "author": {
                                    "email": "49699333+dependabot[bot]@users.noreply.github.com",
                                    "name": "dependabot[bot]"
                                },
                                "message": "chore(deps-dev): bump @types/node from 18.7.9 to 18.7.13 (#338)",
                                "distinct": false,
                                "url": "https://api.github.com/repos/iobroker-community-adapters/ioBroker.ring/commits/d57bad671ad2bbc41f36207063585caca1c5d00e"
                            },
                            {
                                "sha": "55739c215a89e4a961dfe8ca5a63d9e8ff9cc528",
                                "author": {
                                    "email": "49699333+dependabot[bot]@users.noreply.github.com",
                                    "name": "dependabot[bot]"
                                },
                                "message": "chore(deps-dev): bump @typescript-eslint/eslint-plugin (#341)",
                                "distinct": false,
                                "url": "https://api.github.com/repos/iobroker-community-adapters/ioBroker.ring/commits/55739c215a89e4a961dfe8ca5a63d9e8ff9cc528"
                            },
                            {
                                "sha": "cbc96e9571671ea0f1db18915b6e10c5263ad39d",
                                "author": {
                                    "email": "49699333+dependabot[bot]@users.noreply.github.com",
                                    "name": "dependabot[bot]"
                                },
                                "message": "chore(deps-dev): bump typescript from 4.7.4 to 4.8.2\n\nBumps [typescript](https://github.com/Microsoft/TypeScript) from 4.7.4 to 4.8.2.\n- [Release notes](https://github.com/Microsoft/TypeScript/releases)\n- [Commits](https://github.com/Microsoft/TypeScript/compare/v4.7.4...v4.8.2)\n\n---\nupdated-dependencies:\n- dependency-name: typescript\n  dependency-type: direct:development\n  update-type: version-update:semver-minor\n...\n\nSigned-off-by: dependabot[bot] \u003csupport@github.com\u003e",
                                "distinct": true,
                                "url": "https://api.github.com/repos/iobroker-community-adapters/ioBroker.ring/commits/cbc96e9571671ea0f1db18915b6e10c5263ad39d"
                            }
                        ]
                    },
                    "repo": {
                        "id": 161008550,
                        "name": "iobroker-community-adapters/ioBroker.ring",
                        "url": "https://api.github.com/repos/iobroker-community-adapters/ioBroker.ring"
                    },
                    "actor": {
                        "login": "dependabot[bot]",
                        "id": 49699333,
                        "avatar_url": "https://avatars.githubusercontent.com/u/49699333?",
                        "gravatar_id": "",
                        "url": "https://api.github.com/users/dependabot[bot]"
                    },
                    "org": {
                        "login": "iobroker-community-adapters",
                        "id": 46189175,
                        "avatar_url": "https://avatars.githubusercontent.com/u/46189175?",
                        "url": "https://api.github.com/orgs/iobroker-community-adapters"
                    },
                    "created_at": "2022-08-29T02:25:05Z",
                    "id": "23683480363"
                },
                "payload": {
                    "devDay": 0,
                    "devYear": 0,
                    "merge": 0,
                    "open": 0,
                    "pr": 0
                }
            },
            {
                "event": {
                    "type": "CreateEvent",
                    "public": true,
                    "payload": {
                        "ref": "dependabot/npm_and_yarn/eslint-plugin-react-7.31.1",
                        "ref_type": "branch",
                        "master_branch": "master",
                        "description": "Helpers for React Navigation",
                        "pusher_type": "user"
                    },
                    "repo": {
                        "id": 201448541,
                        "name": "WrathChaos/react-navigation-helpers",
                        "url": "https://api.github.com/repos/WrathChaos/react-navigation-helpers"
                    },
                    "actor": {
                        "login": "dependabot[bot]",
                        "id": 49699333,
                        "avatar_url": "https://avatars.githubusercontent.com/u/49699333?",
                        "gravatar_id": "",
                        "url": "https://api.github.com/users/dependabot[bot]"
                    },
                    "created_at": "2022-08-29T02:25:05Z",
                    "id": "23683480358"
                },
                "payload": {
                    "devDay": 0,
                    "devYear": 0,
                    "merge": 0,
                    "open": 0,
                    "pr": 0
                }
            }
        ],
        "typeMap": {
            "CreateEvent": 1,
            "PushEvent": 1
        }
    }
    ```
