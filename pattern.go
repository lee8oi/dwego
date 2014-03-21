/*
pattern.go is a custom solution for finding command string patterns and determining the
data contained in them. Typical commands expected:

give <item> to <character>
put <item> in <container>
get <item>
get <item> from <object>
use <item> on <object|character>
cast <spell> on <character|enemy>
push <object>
pull <object>
open <object>
lock <object>
unlock <object>
north
south
east
west
<attack> <enemy>
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
