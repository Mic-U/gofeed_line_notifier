# gofeed_line_notifier

GoFeedLineNotifier gets articles from RSS feed and send it to LINE channel via [LINE Messaging API](https://developers.line.biz/en/reference/messaging-api/).

## Prerequisities

Issue Channel access token from LINE Developpers console.

## Usage

### golang

```
$ export RSS_URL={URL of RSS Feed}
$ export LINE_ACCESS_TOKEN={Your Access Token}
$ make build
$ ./dist/exec
```

### Docker

### Build Image

```
$ docker image build -t gofeed_line_notifier .
$ docker container run -ti --rm gofeed_line_notifier:latest -e RSS_URL={URL of RSS Feed} -e LINE_ACCESS_TOKEN={Your Access Token}
```

### Pull Image

```
$ docker login -u {GITHUB_USER_NAME} -p {GITHUB_ACCESS_TOKEN} docker.pkg.github.com
$ docker pull docker.pkg.github.com/mic-u/gofeed_line_notifier/beta:1.0
```