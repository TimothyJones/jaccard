Jaccard
-------

Simple go program to determine the [Jaccard Index](http://en.wikipedia.org/wiki/Jaccard_index) between the lines in two files

Given two files, it prints the jaccard index to standard out, measuring the line-by-line similarity.

[![Build Status](https://travis-ci.org/TimothyJones/jaccard.svg?branch=master)](https://travis-ci.org/TimothyJones/jaccard)

Uses Ralph Caraveo's [golang-set](https://github.com/deckarep/golang-set) to back the sets.

## Example usage

    $ jaccard README.md README.md
    1


    $ head README.md > top-of-readme
    $ jaccard README.md top-of-readme
    0.6363636363636364    
