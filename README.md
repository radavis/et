# ET

The [et gem](https://rubygems.org/gems/et), rewritten in Go.

## Features

* [x] et init
* [ ] et list
* [ ] et get slug
* [ ] et submit

## et init

Prompts the user for a username and API key and creates a `.et` file in your `$HOME` directory.

## et list

curl example. API uses basic auth.

```no-highlight
curl -u $USERNAME:$TOKEN --request GET https://learn.launchacademy.com/lessons.json
```
