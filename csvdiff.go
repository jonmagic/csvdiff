package main

import (
  "encoding/csv"
  "fmt"
  "io"
  "log"
  "os"
)

// Usage instructions to print if incorrect options or -h or --help are passed.
const Usage = "usage: csvdiff new_revision.csv old_revision.csv"

// The column index to use as the identifier (matching) column. This should be
// made into an option eventually.
const identifierColumn = 0

// Extracts filenames from command line arguments. Should be renamed to
// argumentParser and take other options as well.
func revisionFilenames() (filenameX string, filenameY string, err error) {
  err = fmt.Errorf(Usage)

  if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
    err = fmt.Errorf(Usage)
    return "", "", err
  }
  if len(os.Args) == 3 {
    return os.Args[1], os.Args[2], nil
  }

  return "", "", err
}

// Takes two csv files and starts the diffing operation. Goes through revisionX
// comparing each row against every row in revisionY and then repeats the
// process in reverse.
func csvDiff(fileX, fileY io.Reader) (err error) {
  revisionXRows, err := csv.NewReader(fileX).ReadAll()
  revisionYRows, err := csv.NewReader(fileY).ReadAll()

  for _, revisionXRow := range revisionXRows {
    compareRowAgainstRows(revisionXRow, revisionYRows, true)
  }

  for _, revisionYRow := range revisionYRows {
    compareRowAgainstRows(revisionYRow, revisionXRows, false)
  }

  return
}

// Takes a single row in one revision and compares every row in the other
// revision against that row. If changesAndAdditions is true it only records
// changes and additions, if it is false it only records removals.
func compareRowAgainstRows(row []string, rows [][]string, changesAndAdditions bool) {
  match := false
  // Ideally this algorithm would be more efficient by removing rows once they have been matched.
  for _, comparisonRow := range rows {
    if row[identifierColumn] == comparisonRow[identifierColumn] {
      match = true
      changed := false
      for i := 1; i < len(row); i++ {
        if row[i] != comparisonRow[i] {
          changed = true
        }
      }
      if changed && changesAndAdditions {
        fmt.Println("Changed: ", comparisonRow, row)
      }
      return
    }
  }
  if !match {
    if changesAndAdditions {
      fmt.Println("Added: ", row)
    } else {
      fmt.Println("Removed: ", row)
    }
  }
}

// Application starts here.
func main() {
  filenameX, filenameY, err := revisionFilenames()

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fileX := os.Stdin
  if filenameX != "" {
    if fileX, err = os.Open(filenameX); err != nil {
      log.Fatal(err)
    }
    defer fileX.Close()
  }

  fileY := os.Stdin
  if filenameY != "" {
    if fileY, err = os.Open(filenameY); err != nil {
      log.Fatal(err)
    }
    defer fileY.Close()
  }

  if err = csvDiff(fileX, fileY); err != nil {
    log.Fatal(err)
  }
}
