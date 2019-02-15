# mockbug
Gomock usability issue

Example of confusing output:

```
--- FAIL: TestNilVsEmpty (0.00s)
    example_test.go:24: Unexpected call to *mock_example.MockThing.Apply([0xc0000900a0]) at /usr/local/google/home/briankennedy/golang/work/src/github.com/briantkennedy/mockbug/pkg/mock_example/mock_thing.go:39 because: 
        Expected call at /usr/local/google/home/briankennedy/golang/work/src/github.com/briantkennedy/mockbug/pkg/example/example_test.go:19 doesn't match the argument at index 0.
        Got: &{[]}
        Want: is equal to &{[]}
    asm_amd64.s:522: missing call(s) to *mock_example.MockThing.Apply(is equal to &{[]}) /usr/local/google/home/briankennedy/golang/work/src/github.com/briantkennedy/mockbug/pkg/example/example_test.go:19
    asm_amd64.s:522: aborting test due to missing call(s)
FAIL
FAIL	github.com/briantkennedy/mockbug/pkg/example	0.021s
?   	github.com/briantkennedy/mockbug/pkg/mock_example	[no test files]
```
