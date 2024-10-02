# Time based Disk Usage

## Key features

- Find directories with defined size increment in time
- Save report as CSV

### Find directories with defined size increment in time

Find directories with more than 300 MB in last 24 hours

```bash
tdu -path /home -min 300 -time 24
```

Find directories with more than 300 MB in last 24 hours and save report as CSV

```bash
tdu -path /home -min 300 -time 24 -csv /root/home-report.csv
```
