# uptime

Tell how long the system has been running. (Aim for having good compatibility with original uptime implementation.)

```
$ ./uptime --help
 uptime gives a one line display of the following information.  The current time, how long the
system has been running, how many users are currently logged on, and the system load averages
for the past 1, 5, and 15 minutes.

Usage:
  uptime [flags]

Flags:
  -h, --help     help for uptime
  -p, --pretty   show uptime in pretty format
  -s, --since    system up since
```

## Example Usage

- Basic usage (Currently, pseudo user counter is implemented...)

```
$ ./uptime
 16:30:05 up  5: 3,  1 user,  load average: 4.16, 3.74, 3.45
```

- Pretty formatting for uptime

```
$ ./uptime -p
up 5 hours, 9 minutes
```

- System up since

```
$ ./uptime -s
2018-08-17 11:26:15
```
