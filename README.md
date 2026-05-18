# Go-Constant

A tiny, generic Go library for defining typed constant sets with safe string parsing.

Go has no built-in enum type. The common workaround — `iota` or typed strings — gives you no way to safely parse an unknown string into a valid constant. Go-Constant fills that gap with a single generic type that works with any value that implements `fmt.Stringer`.

## Requirements

Go 1.18+ (generics)

## Installation

```bash
go get github.com/your-username/Go-Constant
```

## Usage

### String-based constants

```go
type Direction string

func (d Direction) String() string { return string(d) }

const (
    North Direction = "north"
    South Direction = "south"
    East  Direction = "east"
    West  Direction = "west"
)

var Directions = constant.NewConstSet(North, North, South, East, West)

func main() {
    d, err := Directions.Parse("east")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(d) // east

    _, err = Directions.Parse("up")
    fmt.Println(err) // "up" is not a valid constant
}
```

### Struct-based constants

```go
type Color struct {
    Name string
    Hex  string
}

func (c Color) String() string { return c.Name }

var (
    Red   = Color{Name: "red",   Hex: "#FF0000"}
    Green = Color{Name: "green", Hex: "#00FF00"}
    Blue  = Color{Name: "blue",  Hex: "#0000FF"}

    Colors = constant.NewConstSet(Red, Red, Green, Blue)
)

func main() {
    c, err := Colors.Parse("green")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(c.Hex) // #00FF00
}
```

## API

### `NewConstSet[V fmt.Stringer](defaultVal V, members ...V) ConstSet[V]`

Creates a new constant set. The first argument is the default value returned when `Parse` fails.

### `(ConstSet[V]) Parse(value string) (V, error)`

Converts a string to a member of the set by matching against each member's `String()` value. Returns the default value and a descriptive error if no match is found.

### `(ConstSet[V]) Members() []V`

Returns all members of the set.

### `(ConstSet[V]) Default() V`

Returns the default value of the set.

## Why not `iota`?

| Feature | `iota` | Go-Constant |
|---|---|---|
| Type-safe | Yes | Yes |
| String parsing | Manual switch | Built-in |
| Works with structs | No | Yes |
| List all values | Manual | `Members()` |

## Contributing

Contributions are welcome. Open an issue or submit a pull request.

## License

MIT
