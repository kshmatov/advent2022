package stock

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = strings.Split(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`, "\n")

func TestStack(t *testing.T) {
	s := newStack()
	assert.Equal(t, -1, s.top)
	assert.Nil(t, s.stack)

	s.bottomPop("A")
	assert.Equal(t, 0, s.top)
	assert.Equal(t, "A", s.stack[0])
	assert.Equal(t, 1, len(s.stack))

	s.push("B")
	assert.Equal(t, 1, s.top)
	assert.Equal(t, "B", s.stack[1])
	assert.Equal(t, 2, len(s.stack))

	s.bottomPop("A")
	assert.Equal(t, 2, s.top)
	assert.Equal(t, "A", s.stack[0])
	assert.Equal(t, 3, len(s.stack))

	b := s.head()
	assert.Equal(t, "B", b)
	assert.Equal(t, 2, s.top)
	assert.Equal(t, 3, len(s.stack))

	b, err := s.pop()
	assert.Nil(t, err)
	assert.Equal(t, "B", b)
	assert.Equal(t, 1, s.top)
	assert.Equal(t, 2, len(s.stack))
}

func TestStock(t *testing.T) {
	st := newStock(3)
	assert.Equal(t, 3, len(st.columns))
	for _, col := range st.columns {
		assert.NotNil(t, col)
	}
	st.insert("A", 0)
	assert.Equal(t, 1, len(st.columns[0].stack))
	assert.Equal(t, 0, st.columns[0].top)
	assert.Equal(t, "A", st.columns[0].head())
}

func TestBuildStock(t *testing.T) {
	data := testData[0:5]
	st := buildStock(data, 3)
	assert.Equal(t, 3, len(st.columns))
	assert.Equal(t, "NDP", st.getHead())
}

func TestStockOperations(t *testing.T) {
	data := testData[0:5]
	st := buildStock(data, 3)
	err := st.move(1, 2, 1)
	assert.Nil(t, err)
	assert.Equal(t, "DCP", st.getHead())

	err = st.move(3, 1, 3)
	assert.Nil(t, err)
	assert.Equal(t, "CZ", st.getHead())
	assert.Equal(t, "PDNZ", strings.Join(st.columns[2].stack, ""))

	err = st.move(2, 2, 1)
	assert.Nil(t, err)
	assert.Equal(t, "MZ", st.getHead())
	assert.Equal(t, "CM", strings.Join(st.columns[0].stack, ""))
}
func TestParseCommand(t *testing.T) {
	c, f, to, err := parseMoveCmd(testData[5])
	assert.Nil(t, err)
	assert.Equal(t, 1, c)
	assert.Equal(t, 2, f)
	assert.Equal(t, 1, to)
}

func TestMoveStock(t *testing.T) {
	data := testData[:5]
	move := testData[5:]
	st := buildStock(data, 3)
	err := moveStock(move, st.move)
	assert.Nil(t, err)
	assert.Equal(t, "CMZ", st.getHead())
}

func TestMoveOrdered(t *testing.T) {
	data := testData[:5]
	move := testData[5:]
	st := buildStock(data, 3)
	err := moveStock(move, st.moveOrdered)
	assert.Nil(t, err)
	assert.Equal(t, "MCD", st.getHead())
}
