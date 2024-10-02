[![Go](https://github.com/spetr/tdu/actions/workflows/go.yml/badge.svg)](https://github.com/spetr/tdu/actions/workflows/go.yml)
[![License](https://img.shields.io/github/license/spetr/tdu)](https://github.com/spetr/tdu/blob/main/LICENSE)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/spetr/tdu?label=latest%20release)](https://github.com/spetr/tdu/releases/latest)

# Time based Disk Usage

## Key features

- Find directories with defined size increment in time
- Save report as CSV

### Find directories with defined size increment in time

Find directories with more than 300 MB in last 24 hours

```bash
tdu -path /home -min=300 -time=24
```

Find directories with more than 300 MB in last 24 hours and save report as CSV

```bash
tdu -path /home -min=300 -time=24 -csv=/root/home-report.csv
```
