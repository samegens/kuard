/*
Copyright 2017 The KUAR Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package htmlutils

import (
	"fmt"
	"hash/fnv"
	"strings"
)

// ColorFromString returns a CSS color string by hashing the incoming string.
// Used to visually identify different versions.
func ColorFromString(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	hue := h.Sum32() % 360
	return fmt.Sprintf("hsl(%d,100%%,50%%)", hue)
}

// VersionNameColor parses the version string and returns a CSS color
// that matches the version name (blue, green, purple).
func VersionNameColor(version string) string {
	v := strings.ToLower(version)
	if strings.Contains(v, "blue") {
		return "#0066CC"
	} else if strings.Contains(v, "green") {
		return "#00AA00"
	} else if strings.Contains(v, "purple") {
		return "#9933FF"
	}
	return "#666666" // fallback gray
}
