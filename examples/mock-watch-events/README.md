# Use the fake watch API and simulate the watch behaviour sequence

Sometimes we encounter the case where we need to simulate
the watch event behaviour in order to test code that uses
client-go `watch` API.

In this example we will see how to mock watch API events
sequence as part of the unit testing.

Before that 