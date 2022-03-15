# Locks

## References

- [SO](https://stackoverflow.com/questions/19148809/how-to-use-rwmutex)

A mutex from a package like `sync` provides locking and unlocking functionality so that when writing to data structures, we don't have more than one thread doing anything at once.

A struct will usually hold the mutex in some way:

```go
type Cereal struct {
  items []string
  mutex sync.Mutex
}
```

This then allows calling the mutex on that struct when the data is being updated.

```go
func (c *Cereal) Add(name string) {
  c.mutex.Lock()
  defer c.mutex.Unlock()
  append(c.items, name)
}
```

a `Lock` is exclusive. Only one thread can access that data at any time. You can create a semi-exclusive lock via `sync.RWMutex` and `RLock` (with `RUnlock` to unlock).

- `RLock` is a shared _read_ lock.
- When the mutex is `RLocked`'d a call to `Lock` is blocked.
  -- The same vice versa.

Further to `sync` there is `atomic`. This allows updating primtiive values in a safe way, it has methods tied to primitive data types such as `AddInt64` which accepts the variable to update along with the value to add.

## Example

If you have a data store that matches the `map[string]int64` type signature and new keys and items can be added then we should be using an `RLock` to read out the correct value while still allowing access to other values via the same mechanism. Then a `Lock` can be used when initialising a new value.

```go
type Stat struct {
  counters map[string]*int64
  // RWMutex to control locking
  mutex    sync.RWMutex
}

func InitStat() Stat {
  return Stat{counters: make(map[string]*int64)}
}

func (s *Stat) Count(name string) int64 {
  var counter *int64
  if counter = getCounter(name); counter == nil {
      counter = initCounter(name);
  }
  // `atomic` can update the primitive we've received
  return atomic.AddInt64(counter, 1)
}

func (s *Stat) getCounter(name string) *int64 {
  // Shared read lock created to allow access
  // to multiple threads when reading out the
  // counter they require
  s.mutex.RLock()
  defer s.mutex.RUnlock()
  return s.counters[name]
}

func (s *Stat) initCounter(name string) *int64 {
  // Full lock when writing to the `counters` datastore
  // and adding a new counter item with relevant primitive
  s.mutex.Lock()
  defer s.mutex.Unlock()
  counter := s.counters[name]
  if counter == nil {
      value := int64(0)
      counter = &value
      s.counters[name] = counter
  }
  return counter
}
```
