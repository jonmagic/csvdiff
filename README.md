# csvdiff

In order to compare two csv files and get something more meaningful than normal line diffs you can choose an identifier column in your data that will be used to match rows. If no identifier column is chosen you can of course fall back to the line number as the unique identifier.

Once a row from two revisions is matched you can compare the values in every other cell in that row and *show the changes in a meaningful way*.

For text normal red/green comparisons are handy, but things get much more interesting when you *start comparing cells with numerical values as you can show the change in value*. Date and time cells could show a *change in days, hours, minutes, seconds, or all of the above*. If you can determine the column type more meaningful comparisons can be done.

This repository is my attempt at building a csv differ in Go.

## Usage

```bash
$ go build csvdiff.go

$ ./csvdiff --help
usage: csvdiff new_revision.csv old_revision.csv
exit status 1

$ ./csvdiff revisionX.csv revisionY.csv
Changed:  [1 jonmagic 1000] [1 jonmagic 2500]
Changed:  [2 benbalter 9999] [2 benbalter 8888]
Added:  [4 caged 6200]
Removed:  [3 gjtorikian 8500]
```

## Output Format

The output is currently very limited but eventually I hope to return json that looks like this:

```json
{
  "changed": [
    {
      "rowIndex": 1,
      "cells": [
        {
          "previous_value": 1,
          "value": 1,
          "type": "integer"
        },
        {
          "previous_value": "jonmagic",
          "value": "jonmagic",
          "type": "string"
        },
        {
          "previous_value": 1000,
          "value": 2500,
          "type": "integer",
          "change": 1500
        }
      ]
    },
    {
      "rowIndex": 2,
      "cells": [
        {
          "previous_value": 2,
          "value": 2,
          "type": "integer"
        },
        {
          "previous_value": "benbalter",
          "value": "benbalter",
          "type": "string"
        },
        {
          "previous_value": 9999,
          "value": 8888,
          "type": "integer",
          "change": -1111
        }
      ]
    }
  ],
  "added": [
    {
      "rowIndex": 3,
      "cells": [
        {
          "value": 4,
          "type": "integer"
        },
        {
          "value": "caged",
          "type": "string"
        },
        {
          "value": 6200,
          "type": "integer"
        }
      ]
    }
  ],
  "removed": [
    {
      "rowIndex": 3,
      "cells": [
        {
          "value": 3,
          "type": "integer"
        },
        {
          "value": "gjtorikian",
          "type": "string"
        },
        {
          "value": 8500,
          "type": "integer"
        }
      ]
    }
  ]
}
```

## Contributors

* [jonmagic](https://github.com/jonmagic)
