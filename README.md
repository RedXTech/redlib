# redlib
##### A go library of functions that I like to use.

This so far just includes two function, `redlib.Readln()`, and `redlib.Get()`.

### Readln()
`redlib.Readln()` works the same way that `Console.ReadLine()` works in C#. Typically, you would run a `fmt.Print("What is your name: ")`, and then `name, err := redlib.Readln()`, and everything they type in before pressing enter is saved into the `name` variable, while an error would be saved to `err`. If there isn't an error, `err` will be `nil`.

### Get()
`redlib.Get` works as a very simple `http.Get` wrapper. `site, err := redlib.Get("https://redxtech.ca")` would set `site` to a string of the website passed, and the error handling is the same as `redlib.Readln()`.