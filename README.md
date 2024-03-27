# robot
Re-designed exercise from the Exercism platform.

## Run tests

```shell
➜  robot git:(main) go test -race -v
=== RUN   TestNameValid
=== PAUSE TestNameValid
=== RUN   TestNameSticks
=== PAUSE TestNameSticks
=== RUN   TestSuccessiveRobotsHaveDifferentNames
=== PAUSE TestSuccessiveRobotsHaveDifferentNames
=== RUN   TestResetName
=== PAUSE TestResetName
=== RUN   TestMultipleNames
=== PAUSE TestMultipleNames
=== RUN   TestCollisions
=== PAUSE TestCollisions
=== CONT  TestNameValid
=== CONT  TestNameSticks
=== CONT  TestResetName
=== CONT  TestMultipleNames
--- PASS: TestResetName (0.00s)
--- PASS: TestNameSticks (0.00s)
--- PASS: TestNameValid (0.00s)
=== CONT  TestSuccessiveRobotsHaveDifferentNames
=== CONT  TestCollisions
--- PASS: TestSuccessiveRobotsHaveDifferentNames (0.00s)
--- PASS: TestMultipleNames (0.01s)
--- PASS: TestCollisions (22.94s)
PASS
ok  	github.com/qba73/robot	24.794s
```
