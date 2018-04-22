# VooT DB
## Asynchronous Key Value Database

> A simple value backup system for `golang`.


## Usage:

```go
package main
import (
	"voot"
)

var Key   = "The value you want to store"

var store = voot.VooT{Name:"DatabaseName",Data:&Key}
var DB    = voot.NewDB(&store)
func main(){
	defer DB.SaveAndClose()
	// All your code goes here
}
```


## Why and when VooT?

**First of all this database is not an actual database!**
It syncs the value of a given key. And stores it into pure json format in a file as the database name given in `voot.VooT{}`.
A single database can contain only a single value.

I started to code `VooT` for storing a huge amount (`~300mb`) of values into a pure json formatted file.

I run a project (`web server`) where I have 1m users! And it's tough to put and query this amount of data and traffic on a `NoSQL` like database.

So, I used `VooT` and I used a global variable with all user data. And I used insert/ query/ update/ delete in the runtime as a `golang` variable. I didn't need to rely on a database server or whateve. It's pure golang's speed.

It made my application `5x` faster and stable.

## What Does VooT Do?

VooT syncs the value of the given key. Saves it in filesystem (in goroutine). So it doesn't make your program slow to save the database!
It saves the value in every `0.5s`
Also on the `SaveAndClose()` event!
Also on `^C`/`Os Interrupt` event!

So you don't have a risk to loss data even on unexpected Exit event!


> For contact: AnikHasibul@outlook.com
