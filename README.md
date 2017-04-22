# ET

The [et gem](https://rubygems.org/gems/et), rewritten in Go.

## Features

* [x] et init
* [x] et list
* [ ] et get slug
* [ ] et submit

## et init

Prompts the user for a username and API key and creates a `.et` file in your `$HOME` directory.

## et list

curl example. API uses basic auth.

```no-highlight
curl -u $USERNAME:$TOKEN --request GET https://learn.launchacademy.com/lessons.json
```

## et get slug

Get the lesson in json format.

```no-highlight
curl -u $USERNAME:$TOKEN --request GET https://learn.launchacademy.com/lessons/slug.json
```

```json
{
  "lesson": {
    "slug": "...",
    "title": "...",
    "body": "...",
    "archive_url": "..."
  }
}
```

Use the `archive_url` to make another GET request for the lesson archive data.

```no-highlight
curl --request GET https://horizon-production.s3.amazonaws.com/downloads/lessons/archive/need-the-ingredients/need-the-ingredients.tar.gz
```

Then, uncompress.
