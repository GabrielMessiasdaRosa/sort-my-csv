# sort-my-csv

## Instalation and usage

### if you have snap support on your system, you can install it from snap store

[![sort-my-csv](https://snapcraft.io/sort-my-csv/badge.svg)](https://snapcraft.io/sort-my-csv)

```bash
sudo snap install sort-my-csv --edge
```

otherwise you can download the binary from [here](https://raw.githubusercontent.com/GabrielMessiasdaRosa/sort-my-csv/master/bin/sort-my-csv)

### Usage

```bash
Usage of /snap/sort-my-csv/1/sort-my-csv:
  --sort string
     ➜ column title
  --order string
     ➜ asc | desc
  --input string
     ➜ input file path
  --output string
     ➜ output file path
  --help
     ➜ show this help
```

#### example

```bash
sort-my-csv --sort=<column_title> --order=<asc|desc> ./file_to_be_sorted.csv ./sorted_file_output.csv
```
