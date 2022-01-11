# Writing Untestable Go Code

This repo contains some example (micro)service. It should expose some issues when tyring to put too much logic into one file or to mix things that should not be mixed together.

## Service ACs

Try to write a program in Go that:

- receives a request
- does some basic authentication
- validates the request
- transforms the request
- calls an external API
- sends a response back to the caller

Cover the code with unit tests!

##  My learnings

When writing code:

- separate data from behaviour
- introduce separate layers for handling business logic and representation
- divide code into services with meaningful name and well-defined interface
- make services small, let them take care of one or only few things
- try to pass only the data that are needed to the services, do not pass all the application around
- try using TDD

When writing unit tests:

- concentrate on the tasks done by the service under test, do not test behaviour that is out of the scope
- as the tested code, unit tests should be small

Try to make your application to look like more like a graph, not like a list:

Instead of

```
   s1
  /|\
 / | \
/  |  \
s2 s3  s4
```

... do

```
 s1
 /\
/  \
s2 s3
   /\
  /  \
 s4  s5
```
