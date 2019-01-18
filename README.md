# cn

A handle little utility to give you the index (1-based) of a label/column name in a header of a CSV.

For example:

```
$ echo data.csv
a,b,c
1,0,0
0,1,0
0,0,1

$ cn b data.csv
2
```

You can also use stdin:

```
$ cat data.csv | cn b -
2
```

## Installation

You can install easily with homebrew:

```
brew tap jefferickson/cn
brew install cn
```
