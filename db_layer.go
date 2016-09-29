package main

//Step3: Implement of interaction with database
type dbDriver interface {
    Create(t Task) error
    ReadById(id *int64) (TaskList, error)
    ReadByAlias(alias *string) (TaskList, error)
    ReadAll() (TaskList, error)
    Update(t Task) error
    Delete(t Task) error
}
