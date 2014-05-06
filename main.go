package main

import (
  "github.com/deckarep/golang-set"
  "fmt"
  "log"
  "bufio"
  "os"
  "flag"
)

func makeSetFromFile(fileName string) (mapset.Set,error) {
  file, err := os.Open(fileName)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  linesSet := mapset.NewSet()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    linesSet.Add(scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    return nil,err
  }

  return linesSet, nil
}

func usage() {
  fmt.Fprintf(os.Stderr, "%s: <file 1> <file 2>:\n", os.Args[0])
  fmt.Println("  ",os.Args[0],"prints a single number to standard out reflecting the jaccard similarity between the lines in <file 1> and <file 2>")
}

func main() {
  flag.Parse()

  if flag.NArg() != 2 {
    usage()
    return
  }

  s1, err := makeSetFromFile(flag.Arg(0))
  if err != nil {
    log.Panic(err)
  }
  s2, err := makeSetFromFile(flag.Arg(1))
  if err != nil {
    log.Panic(err)
  }

  union := s1.Union(s2)
  intersection := s1.Intersect(s2)

  fmt.Println(float64(intersection.Cardinality()) / float64(union.Cardinality()))
}
