# ltsv2tsv
ltsv2tsv will convert ltsv into tsv(csv).

## Example

```golang
package main

import (
        "bytes"
        "fmt"
        "github.com/morix1500/go-ltsv2tsv"
)

var ltsv_str = `
firstname:Donald        lastname:Trump  age:71
firstname:Barack        lastname:Obama  age:56
firstname:George        lastname:Bush   age:71
`

func main() {
        reader := ltsv2tsv.NewConverter(bytes.NewBufferString(ltsv_str))
        records, err := reader.Converter()
        if err != nil {
                panic(err)
        }
        for i := 0; i < len(records); i++ {
                fmt.Println(records[i])
        }
}
```

```bash
# Output: [][]string
[firstname lastname age]
[Donald Trump 71]
[Barack Obama 56]
[George Bush 71]
```

## Installation
This package can be installed with the go get command:

```bash
go get github.com/morix1500/go-ltsv2tsv
```

## License
Please see the LICENSE file for details.  
<https://github.com/morix1500/go-ltsv2tsv/blob/master/LICENSE>

## Author
Shota Omori(Morix)  
<https://github.com/morix1500>
