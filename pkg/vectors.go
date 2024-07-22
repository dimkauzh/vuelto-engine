package vuelto

type Vector2D struct {
	X float64
	Y float64
}

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

// Creates a new 2D vector with x and y values
func NewVector2D(x, y float64) *Vector2D {
	return &Vector2D{
		X: x,
		Y: y,
	}
}

// Adds two 2D vectors together
func AddVector2D(v1, v2 Vector2D) *Vector2D {
	return &Vector2D{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

// Subtracts two 2D vectors from each other
func SubtractVector2D(v1, v2 Vector2D) *Vector2D {
	return &Vector2D{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

// Creates a new 3D vector with x, y and z values
func NewVector3D(x, y, z float64) *Vector3D {
	return &Vector3D{
		X: x,
		Y: y,
		Z: z,
	}
}

// Adds two 3D vectors together
func AddVector3D(v1, v2 Vector3D) *Vector3D {
	return &Vector3D{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
}

// Subtracts two 3D vectors from each other
func SubtractVector3D(v1, v2 Vector3D) *Vector3D {
	return &Vector3D{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}
