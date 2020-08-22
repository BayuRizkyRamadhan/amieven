Please install testify first before continuing this section

```
go get github.com/stretchr/testify
```


Case : 

POS 

we need to have a external order service to calls when an item is sold on the point of sale

Requirements:
- when item sale, the POS record product name and amount to deduct
  - amount to deduct cannot be < 1
- when the POS already record product name and amount to deduct, call order service to
  - send product name and amount to deduct