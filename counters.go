
package counters

type Counter interface {
    Add(interface{})
    Get() string
    GetRaw() interface{}
    Name() string
    Set(interface{})
}
