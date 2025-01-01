<!--markdownlint-disable md010-->
# 📐 Vector math module

Features for doing math with both 2D and 3D vectors.

## Types

```go
type Vector2D struct {
	X float64
	Y float64
}

type Vector3D struct {
	X float64
	Y float64
	Z float64
}
```

## Usage

You can create a 2D vector or a 3D vector using the `NewVector2D()` and `NewVector3D()` functions. Both take their X and Y (and Z for 3D) values as arguments.

For example:

```go
my2dVector := NewVector2D(20, 40) // { X: 20, Y: 40 }

my3dVector := NewVector2D(20, 40, 10) // { X: 20, Y: 40, Z: 10 }
```

Then, you can do math with them. For now we support adding and subtracting vectors. Basically, given two vectors:

```go
vector1 := {
    X: 20,
    Y: 20
}
// and
vector2 := {
    X: 10,
    Y: 15
}
```

For each value, the operation will be done with it's counterpart from the other vector. In other words:

```go
// vector1 has X 20 - vector2 has X 10
// subtract will do this:
X: 20 - 10 = 10

// vector1 has Y 20 - vector2 has Y 15
Y: 20 - 15 = 5
```

So, subtracting vector2 from vector1 would give us the following `Vector2D`:

```go
{
    X: 10,
    Y: 5
}
```

That said, you can add or subtract 2D / 3D vectors as follows. All functions follow the same structure, take one vector as the 1st argument, vector to add/subtract as the 2nd argument, and returns a new 2D / 3D vector:

```go
vector3 := AddVector2D(vector1, vector2)
// vector3 = { X: 30, Y: 25 }

vector4 := SubtractVector2D(vector1, vector2)
// vector4 = { X: 10, Y: 5 }

vector5 := NewVector3D(20, 20, 20)
vector6 := NewVector3D(5, 10, 10)

vector7 := AddVector3D(vector5, vector6)
// vector7 = { X: 25, Y: 30, Z: 30 }

vector8 := SubtractVector3D(vector5, vector6)
// vector8 = { X: 15, Y: 10, Z: 10 }
```
