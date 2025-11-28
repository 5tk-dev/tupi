package tupi

// import (
// 	"fmt"
// 	"sync"
// )

// var (
// 	once             sync.Once
// 	muFielder        sync.RWMutex
// 	fielderResgistry map[*Fielder[any]]bool
// )

// func initRegistry() {
// 	once.Do(func() {
// 		fielderResgistry = map[*Fielder[any]]bool{}
// 	})
// }

// func RegisterFielder[T any](key string, schema *Fielder[T]) error {
// 	initRegistry()

// 	muFielder.Lock()
// 	defer muFielder.Unlock()
// 	_, ok := fielderResgistry[(*Fielder[any])(schema)]
// 	if ok {
// 		return fmt.Errorf("fielder %q already registered", key)
// 	}
// 	fielderResgistry[(*Fielder[any])(schema)] = true
// 	return nil
// }

// func GetFielder[T any]() *Fielder[T] {
// 	muFielder.RLock()
// 	defer muFielder.RUnlock()
// 	return fielderResgistry[(*Fielder[T])]
// }

// func ValidateSchema[T any](key string, data any) Schema {
// 	f := GetFielder(key)
// 	if f == nil {
// 		return &schema{
// 			errors: []error{
// 				fmt.Errorf("fielder %q is undefined", key),
// 			},
// 		}
// 	}
// 	return f.Decode(data)
// }
