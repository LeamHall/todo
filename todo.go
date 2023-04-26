// name    :	todo.go
// version :	0.0.1
// date    :	20230424
// author  :	Leam Hall
// desc    :	The ToDo thing

package todo

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "time"
)

// item Struct is a ToDo item
type item struct {
    Task        string
    Done        bool
    CreatedAt   time.Time
    CompletedAt time.Time
}

type List []item

// Add creates a new todo item and adds it to the list
func (l *List) Add(task string) {
    t := item{
        Task:           task,
        Done:           false,
        CreatedAt:      time.Now(),
        CompletedAt:    time.Time{},
    }
    *l = append(*l, t)
}

// Complete marks a ToDo item as completed by setting Done = true
// and CompletedAt to the current time
func (l *List) Complete(i int) error {
    ls := *l
    if i <= 0 || i > len(ls) {
        return fmt.Errorf("Item %d does not exist", i)
    }
    ls[i-1].Done = true
    ls[i-1].CompletedAt = time.Now()
    return nil
}

// Delete removes a ToDo item from the list
func (l *List) Delete(i int) error {
    ls := *l
    if i <= 0 || i > len(ls) {
        return fmt.Errorf("Item %d does not exist", i)
    }
    *l = append(ls[:i-1], ls[i:]...)
    return nil
}

// Save encodes the List as JSON, and saves it to the given filename
func (l *List) Save(filename string) error {
    js, err := json.Marshal(l)
    if err != nil {
        return err
    }
    return os.WriteFile(filename, js, 0644)
}

// Get opens the given filename, decodes the JSON, and returns a list
func (l *List) Get(filename string) error {
    file, err :=  os.ReadFile(filename)
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            return nil
        }
        return err
    }
    if len(file) == 0 {
        return nil
    }
    return json.Unmarshal(file,l)
}





