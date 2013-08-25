# csvdiff

In order to compare two csv files and get something more meaningful than normal line diffs you have to choose an identifier column in your data that will be used to match rows.

Once a row from two revisions is matched you can compare the values in every other cell in that row and show the changes in a meaningful way.

For text normal red/green comparisons are handy, but things get much more interesting when you start comparing cells with numerical values as you can show the change in value. Date and time cells could show a change in days or hours.

This repository is my attempt at building a csv differ in Go.
