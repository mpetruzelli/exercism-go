package react

const (
	testVersion = 5
)

var callbackID int

// SpreadsheetCell defines a cell in a spreadsheet.
type SpreadsheetCell struct {
	value      int
	dependents []spreadsheetCellDependency
	callbacks  map[int]func(int)
	dirty      bool
	lastValue  int
}

type spreadsheetCellDependency struct {
	cell   *SpreadsheetCell
	update func(int)
}

// NewCell creates a new SpreadsheetCell.
func NewCell(value int) *SpreadsheetCell {
	return &SpreadsheetCell{value: value, callbacks: map[int]func(int){}}
}

// Value returns the current value of the cell.
func (cell *SpreadsheetCell) Value() int {
	return cell.value
}

// SetValue sets the value of the cell.
func (cell *SpreadsheetCell) SetValue(value int) {
	cell.setValue(value)
	cell.performCallbacks()
}

func (cell *SpreadsheetCell) setValue(value int) {
	if value != cell.value {
		if !cell.dirty {
			cell.lastValue = cell.value
			cell.dirty = true
		} else if value == cell.lastValue {
			cell.dirty = false
		}

		cell.value = value

		for _, dependent := range cell.dependents {
			dependent.update(value)
		}
	}
}

func (cell *SpreadsheetCell) performCallbacks() {
	if cell.dirty {
		for _, callback := range cell.callbacks {
			callback(cell.value)
		}

		cell.dirty = false

		for _, dependent := range cell.dependents {
			dependent.cell.performCallbacks()
		}
	}
}

type spreadsheetCallbackCanceler struct {
	cell  *SpreadsheetCell
	index int
}

// AddCallback adds a callback which will be called when the value changes.
// It returns a Canceler which can be used to remove the callback.
func (cell *SpreadsheetCell) AddCallback(callback func(int)) Canceler {
	callbackID++
	cell.callbacks[callbackID] = callback
	return spreadsheetCallbackCanceler{cell: cell, index: callbackID}
}

// Cancel removes the callback.
func (canceler spreadsheetCallbackCanceler) Cancel() {
	delete(canceler.cell.callbacks, canceler.index)
}

func (cell *SpreadsheetCell) addDependent(dependent *SpreadsheetCell, update func(int)) {
	cell.dependents = append(cell.dependents, spreadsheetCellDependency{cell: dependent, update: update})
}

// Spreadsheet defines a spreadsheet.
type Spreadsheet struct{}

// New creates a new Spreadsheet.
func New() Spreadsheet {
	return Spreadsheet{}
}

// CreateInput creates an input cell linked into the reactor
// with the given initial value.
func (spreadsheet Spreadsheet) CreateInput(value int) InputCell {
	return NewCell(value)
}

// CreateCompute1 creates a compute cell which computes its value
// based on one other cell. The compute function will only be called
// if the value of the passed cell changes.
func (spreadsheet Spreadsheet) CreateCompute1(cell Cell, compute func(int) int) ComputeCell {
	c := NewCell(compute(cell.Value()))

	cell.(*SpreadsheetCell).addDependent(
		c,
		func(value int) {
			c.setValue(compute(value))
		},
	)

	return c
}

// CreateCompute2 is like CreateCompute1, but depending on two cells.
// The compute function will only be called if the value of any of the
// passed cells changes.
func (spreadsheet Spreadsheet) CreateCompute2(cell1 Cell, cell2 Cell, compute func(int, int) int) ComputeCell {
	c := NewCell(compute(cell1.Value(), cell2.Value()))

	cell1.(*SpreadsheetCell).addDependent(
		c,
		func(value1 int) {
			c.setValue(compute(value1, cell2.Value()))
		},
	)
	cell2.(*SpreadsheetCell).addDependent(
		c,
		func(value2 int) {
			c.setValue(compute(cell1.Value(), value2))
		},
	)

	return c
}
