# cn

Usage: `cn [-d] [-h HEADERFILE] LABEL FILE`

A handy little utility to give you the index (1-based) of a label/column name in a header of a CSV. Using the `-d` flag, only the data under the label/column name is returned.

For example:

```
$ cat data.csv
a,b,c
1,0,0
0,1,0
0,0,1

$ cn b data.csv
2

$ cn -d b data.csv
0
1
0
```

You can also use stdin:

```
$ cat data.csv | cn b -
2

$ cat data.csv | cn -d b -
0
1
0
```

This works well with other tools, e.g.:

```
$ cat data2.csv
a,b,c
1,2,3
4,5,6
7,8,9

$ cat data2.csv | cn -d b - | awk -F, '{s+=$1}END{print s}'
15
```

You can even provide the headers from another file, which works well when you are filtering data (coming in from `grep` via stdin, for example) and still want to select a column. For example, let's say you have the following data. You want to sum the `metric_2` column for all rows with `type == 'a'`:

```
$ cat data3.csv
type,metric_1,metric_2,metric_3
a,1,2,3
b,4,5,6
a,7,8,9

$ cat data3.csv | grep a | cn -h data3.csv -d metric_2 - | awk -F, '{s+=$1}END{print s}'
10
```

## Installation

You can install easily with homebrew:

```
brew tap jefferickson/cn
brew install cn
```

Or from source:

```
go get github.com/jefferickson/cn
go install github.com/jefferickson/cn
```
