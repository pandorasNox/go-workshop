# Arrays, Slices, Maps

## Arrays, Slices

### Arrays
- initilized with fixed length (`capasity`)
- type of elements and length are both part of the array’s type
  - e.g. `var a [5]int`
    - type is `[5]int`, where `5` is part of the type
- set values by index
  - `a[4] = 100`

### Slices
- slices are more like the classic arrays you might know from interpreded languages
- they can grow
- uses array as underlying type + header + some magic
- works mostly the same as arrays
- adding to arrays with `append` function

## Maps
- maps are more like the dicts, a key-value store datatype
- keys are unique within a map
- keys are hashed | hash(key) → hashed_value
- any types is allowed where the comparison operators `==` and `!=` are fully defined
