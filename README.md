# go-bored
Go wrapper for [Bored API](https://www.boredapi.com).

## Installation
```
go get -u github.com/trubitsyn/go-bored/v1
```
## Usage
```go
package main

import (
    "fmt"
    boredapi "github.com/trubitsyn/go-bored/v1"
)

func main() {
	activity, err := boredapi.RandomActivity()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(activity.Activity)
}
```

## LICENSE
```
Copyright 2019 Nikola Trubitsyn

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```