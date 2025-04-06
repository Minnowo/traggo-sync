
<h1 align="center">traggo/sync</h1>

Traggo Sync is a commandline tool for syncing data from one [Traggo Server](https://github.com/traggo/server) to another.

Currently supports the following actions:

| Action              | Description                                                           |
|---------------------|-----------------------------------------------------------------------|
| Create missing tags | Creates only tags that are missing on the target server.              |
| Create all tags     | Creates missing tags and updates the color if it doesn't match.       |
| Missing timespans   | Creates timespans that are missing on the target server.              |
| Replace timespans   | Deletes all timespans on the target server, and then creates missing. |
| Dashboards          | Copies all dashboards to the target server.                           |

