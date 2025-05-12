package list

// List is implemented by types that can store data-type
// in a list structure.
type List[T any] interface {
	// Insert adds 'value' into the 'pos' index in the list.
	Insert(pos int, value T)
	
	// Delete removes and return the item at 'pos' index.
	Delete(pos int) T
	
	// Get retrieves the item at 'pos' index.
	Get(pos int) T
	
	// Set replaces the item at 'pos' with 'value'.
	Set(pos, value T)
	
	// Length returns the number of items in the list.
	Length() int
	
	// IsEmpty checks if the list is empty.
	IsEmpty() bool
	
	// Find returns the position of the first occurence of 'value'.
	Find(value T) int
	
	// Clear removes all the items of the list.
	Clear()
}