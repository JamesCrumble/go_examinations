package parts

import (
	"fmt"
)

type Vector2 struct {
	X int
	Y int
}

func (vector *Vector2) ToVec3(z int) Vector3 {
	return Vector3{vector.X, vector.Y, z}

}

// TODO: implement type Vector3 struct {Vector2 Z int}
type Vector3 struct {
	X int
	Y int
	Z int
}

// THIS METHOD IS PART OF Vector3 struct
func (vector *Vector3) ToVec2() Vector2 {
	return Vector2{X: vector.X, Y: vector.Y}
}

// TODO implement working with implicit type for Vector3 of Vector2
func Vec3_to_vec2(vec3 *Vector3) Vector2 {
	return Vector2{vec3.X, vec3.Y}
}

func Vec2_to_vec3(vec2 *Vector2, z int) Vector3 {
	return Vector3{vec2.X, vec2.Y, z}
}

func StructsPartExecute() {
	printInfo("STRUCTS PART")

	var vector2 Vector2 = Vector2{2, 3}

	fmt.Println("Vector3 from convertion function v2 to v3 =>", Vec2_to_vec3(&vector2, 4))
	fmt.Println("Vector3 converted from Vector2 itself =>", vector2.ToVec3(4))

	printInfo("STRUCTS PART")
}
