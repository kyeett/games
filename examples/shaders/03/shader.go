// Copyright 2020 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package main

var Resolution vec2
var Mouse vec2
var Time float

// Uniforms
//func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
//	return vec4(abs(sin(Time)), 0.0, 0.0, 1.0)
//}

// gl_FragCoord
//func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
//	st := position.xy/Resolution
//
//	// Y-coordinate is flipped compared to book of shaders
//	//flipY := 1-st.y
//	return vec4(st.x, st.y, 0.0, 1.0)
//}


/*
Can you tell where the coordinate (0.0, 0.0) is in our canvas?
A: 0,0 is where color is black

What about (1.0, 0.0), (0.0, 1.0), (0.5, 0.5) and (1.0, 1.0)?
A:
(1.0, 0.0) - Color red
(0.0, 1.0) - Color green
(0.5, 0.5) - In the middle of everything
(1.0, 1.0) - Color yellow


Can you figure out how to use u_mouse knowing that the values are in pixels and NOT normalized values? Can you use it to move colors around?
A:
func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	cursor := Mouse/Resolution
	return vec4(cursor.x, cursor.y, 0.0, 1.0)
}

Can you imagine an interesting way of changing this color pattern using u_time and u_mouse coordinates?
A: */
func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	cursor := Mouse/Resolution
	return vec4(cursor.x, cursor.y, abs(sin(Time)), 1.0)
}