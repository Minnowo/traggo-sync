
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


## Example usage


### Sync all missing tags
```sh
ts sync -i http://localhost:3030 -u admin -p admin -t http://localhost:3031 --missing-tags
```

### Sync all missing tags and timespans
```sh
ts sync -i http://localhost:3030 -u admin -p admin -t http://localhost:3031 --missing-tags --missing-timespans
```

### Sync missing timespans later than date
```sh
ts sync -i http://localhost:3030 -u admin -p admin -t http://localhost:3031 --missing-timespans --time-start "2025-01-05 EST"
```

### Sync missing timespans between two dates
```sh
ts sync -i http://localhost:3030 -u admin -p admin -t http://localhost:3031 --missing-timespans --time-start "01 jan 2025 EST" --time-end "03 jan 2025 EST"
```

### Sync to more than one servers
```sh
ts sync -i http://localhost:3030 -u admin -p admin -t http://localhost:3031 -t http://localhost:3032 -t http://localhost:3033 --missing-tags
```

### Sync with different credentials
```sh
ts sync -i http://serverA.com -u adminA -p adminA -t http://serverB.com -u adminB -p adminB
```

