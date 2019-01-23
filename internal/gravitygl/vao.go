package gravitygl

// VertexArrayObject ...
type VertexArrayObject struct {
	ID GLVertexArray
}

// // NewVertexArrayObject ...
// func NewVertexArrayObject(buffers ...Buffer) *VertexArrayObject {
// 	vao := CreateVertexArray()
// 	BindVertexArray(vao)
// 	for i, buffer := range buffers {
// 		b := CreateBuffer()
// 		t := buffer.Target()
// 		BindBuffer(t, b)
// 		BufferData(t, buffer.Usage(), buffer.Data())
// 		fmt.Println("BufferID:", b, "i:", i)
// 		if t == ELEMENT_ARRAY_BUFFER {
// 			VertexAttribPointer(uint32(i), 3, buffer.GLType(), false, 0, 0)
// 		} else {
// 			VertexAttribPointer(uint32(i), 3, buffer.GLType(), false, 0, 0)
// 			EnableVertexAttribArray(uint32(i))
// 		}
// 	}

// 	BindVertexArray(GLNullObj)
// 	return &VertexArrayObject{
// 		ID: vao,
// 	}
// }
