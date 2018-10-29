// Copyright 2014 Hajime Hoshi
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

package opengl

// Since js.Object (Program) can't be keys of a map, use integers (programID) instead.

type locationCache struct {
	uniformLocationCache map[programID]map[string]uniformLocation
	attribLocationCache  map[programID]map[string]attribLocation
}

func newLocationCache() *locationCache {
	return &locationCache{
		uniformLocationCache: map[programID]map[string]uniformLocation{},
		attribLocationCache:  map[programID]map[string]attribLocation{},
	}
}

func (c *locationCache) GetUniformLocation(context *Context, p program, location string) uniformLocation {
	id := getProgramID(p)
	if _, ok := c.uniformLocationCache[id]; !ok {
		c.uniformLocationCache[id] = map[string]uniformLocation{}
	}
	l, ok := c.uniformLocationCache[id][location]
	if !ok {
		l = context.getUniformLocationImpl(p, location)
		c.uniformLocationCache[id][location] = l
	}
	return l
}

func (c *locationCache) GetAttribLocation(context *Context, p program, location string) attribLocation {
	id := getProgramID(p)
	if _, ok := c.attribLocationCache[id]; !ok {
		c.attribLocationCache[id] = map[string]attribLocation{}
	}
	l, ok := c.attribLocationCache[id][location]
	if !ok {
		l = context.getAttribLocationImpl(p, location)
		c.attribLocationCache[id][location] = l
	}
	return l
}
