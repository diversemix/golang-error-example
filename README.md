# golang-error-example
Demonstrate how to use and implement error handling in golang

# correct-type
How do you check that the error is the correct type?

In this example of reading a CSV file there are three tests:
- Reading a good CSV file
- Reading a missing file
- Reading a file which is not CSV

The problem is, if the file is missing in the third test how can you get the 
test to fail by distinguishing between a "File not found" error and a "Bad CSV"
error?

