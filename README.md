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
    server:
        port: 6000 # HTTP port
        health: "/health" # health check api name

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
    ```

5. Start by `./ossinsight-plugin`.

## API

### API Using Step

- `Client` start to connect the API Endpoint by `RAW WebSocket` protocol.
- `Server` will waiting `Client` to send params.
- `Client` send params (API has different params struct).
- `Server` will endless send message to `Client` until connection closed.

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

    E.g.:

    ```json
    {
        "samplingRate": 3,
        "eventType": "PushEvent",
        "repoName": "pingcap-inc/ossinsight-plugin",
        "userName": "Icemap"
    }
    ```

- Result same as GitHub [/events](https://docs.github.com/en/rest/activity/events) API.

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
        "msgList": [
            {event entity},
            {event entity}
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
                "id": "23434997759",
                "type": "IssuesEvent",
                "actor": {
                    "id": 35833812,
                    "login": "pengYYYYY",
                    "display_login": "pengYYYYY",
                    "gravatar_id": "",
                    "url": "https://api.github.com/users/pengYYYYY",
                    "avatar_url": "https://avatars.githubusercontent.com/u/35833812?"
                },
                "repo": {
                    "id": 438886855,
                    "name": "Tencent/tdesign-vue-next",
                    "url": "https://api.github.com/repos/Tencent/tdesign-vue-next"
                },
                "payload": {
                    "push_id": 0,
                    "size": 0,
                    "distinct_size": 0,
                    "ref": "",
                    "head": "",
                    "before": "",
                    "commits": null
                },
                "public": true,
                "created_at": "2022-08-15T08:50:02Z"
            },
            {
                "id": "23434997392",
                "type": "CreateEvent",
                "actor": {
                    "id": 74590681,
                    "login": "naortalmor1",
                    "display_login": "naortalmor1",
                    "gravatar_id": "",
                    "url": "https://api.github.com/users/naortalmor1",
                    "avatar_url": "https://avatars.githubusercontent.com/u/74590681?"
                },
                "repo": {
                    "id": 348281615,
                    "name": "aquasecurity/trivy-plugin-aqua",
                    "url": "https://api.github.com/repos/aquasecurity/trivy-plugin-aqua"
                },
                "payload": {
                    "push_id": 0,
                    "size": 0,
                    "distinct_size": 0,
                    "ref": "plugin-update-v0.59.0",
                    "head": "",
                    "before": "",
                    "commits": null
                },
                "public": true,
                "created_at": "2022-08-15T08:50:01Z"
            }
        ],
        "typeMap": {
            "CreateEvent": 7,
            "IssueCommentEvent": 4,
            "IssuesEvent": 1,
            "PullRequestEvent": 3,
            "PullRequestReviewEvent": 1,
            "PushEvent": 18,
            "WatchEvent": 2
        }
    }
    ```
