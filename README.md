# This is a compiler for my own language

### Target code is JavaScript, so it needs node environment to run the compiled code

### Example
```swift
func main() {
    let ans = fact(5);
    println(ans);
}

func fact(num) {
    if num < 1 {
        return 1;
    }

    return fact(num-1) * num;
}
```

### Keyword

- **func**:    declare a function
- **let**:     declare a immutable variable
- **var**:     declare a mutable variable
- **return**:  return value of function
- **if**:      if statement
- **for**:     for loop
- **println**: convert to console.log in js directly

### Tips

- The entry of this language is main function
- Every expression should be end with a semicolon