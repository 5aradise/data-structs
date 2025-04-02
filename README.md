# Data Structs

Well-tested data structures in Go.

### Calculating the variant number and variant description

The number of the student record book: 18
Variant: 18 % 4 = 2

| Remainder | Initial implementation of the list  | Second realization of the list |
|-----------|-------------------------------------|--------------------------------|
| 2         | list based on built-in arrays/lists | circular singly linked list    |

### Build application and run tests

- Install Go 1.24.1+ : [go.dev](https://go.dev/dl/)

- Clone this repository:

```
git clone https://github.com/5aradise/data-structs.git
```

- Run tests:

```
go test -v ./... -cover
```

- Build app:

```
go run cmd/main.go
```

### Commit on which the CI tests have crashed

[83a509cca9d88f94db0e1de6e9d1dc3c4d0de6e0](https://github.com/5aradise/data-structs/commit/83a509cca9d88f94db0e1de6e9d1dc3c4d0de6e0)

### Conclusions

In my opinion, in this case, the tests were really appropriate and convenient, because the project implementation was logically complex and had many edge cases, and thanks to the previously written test based on a simpler and more understandable model, the further implementation of the more complex version was well controlled by a large number of tests, which made it possible not only to be sure that the written implementation was working, but also to speed up the search for problems and shortcomings.
