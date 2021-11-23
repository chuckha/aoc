
There is a new question every day. The question contains an example. The example
should be parsed and usable. We should minimize network requests.

A Question can be viewed online. But the submission and input happens on the terminal.

This program needs to download the input for a specified day and store it on disk.
When asked for the input for that specific day it can read from disk instead of
making a network call. It can force update the input data.

Input: given a year, day, level
Output: example input, real input

This program can submit answers and tell me if it's right or wrong.

Here is an example

aoc -year 2015 -day 1 -level 1 -example
*writes example output to stdout*










# Usage








`aoc` stdout the input for today (most recent)
aoc | go run main.go | aoc -submit
aoc -year 2015 -day 1 -order 1

/cmd/day1
/internal/lib yay generic data structures
/main.go
/scripts/finished.sh (moves main.go to correct day)


`aoc -example` solves today's (most recent) program with example input.

`aoc -year 2015 -day 1 -order 1 -example` solves the 1st problem of the 1st day of the year 2015 with the example input.

`aoc -submit` solves and submits today's program.

`aoc -year 2015 -day 1 -order 1 -submit` solves and submits the 1st problem of the 1st day of the year 2015 with real input.

aoc will execute specific code against the associated input data. if the input data is not in the local cache, aoc will go fetch and store it locally to reduce network requests.

