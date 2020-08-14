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

func plot(st vec2, pct float) float {
	return smoothstep(pct-0.02, pct, st.y) - smoothstep(pct, pct+0.02, st.y)
}

func Fragment(position vec4, texCoord vec2, clr vec4) vec4 {
	PI := 3.14159265359

	st := position.xy / Resolution
	// Flip to match BoS
	st.y = 1 - st.y

	// 1
	//y := st.x

	// 2
	//y := pow(st.x,5.0);

	// 3
	//y := step(0.5, st.x)

	// 4
	//y := smoothstep(0.1, 0.9, st.x)

	// 5
	//y := sin(st.x)

	// Add time (u_time) to x before computing the sin. Internalize that motion along x.
	//y := sin(st.x+Time)

	//Multiply x by PI before computing the sin. Note how the two phases shrink so each cycle repeats every 2 integers.
	//y := sin(PI*st.x+Time)

	//Multiply time (u_time) by x before computing the sin. See how the frequency between phases becomes more and more compressed. Note that u_time may have already become very large, making the graph hard to read.
	//y := sin(Time*st.x)

	//Add 1.0 to sin(x). See how all the wave is displaced up and now all values are between 0.0 and 2.0.
	//y := sin(st.x) + 1.0

	//Multiply sin(x) by 2.0. See how the amplitude doubles in size.
	//y := 2*sin(st.x)

	//Compute the absolute value (abs()) of sin(x). It looks like the trace of a bouncing ball.
	//y := abs(sin(2*PI*st.x))

	//Extract just the fraction part (fract()) of the resultant of sin(x).
	//y := fract(sin(2*PI*st.x))

	// 13
	y := clamp(st.x, 0.3, 0.7)

	color := vec3(y)

	pct := plot(st, y)
	color = (1-pct)*color + pct*vec3(0, 1, 0)
	return vec4(color, 1.0)
}
