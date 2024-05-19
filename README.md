# This is a compiler for my own language

### Example
```
func main() {
    let ans = fact(5);
    print(ans);
}

func fact(num) {
    if num < 1 {
        return 1;
    }

    return fact(num-1) * num;
}
```

### Target code is JavaScript, so it needs node environment to run the compiled code