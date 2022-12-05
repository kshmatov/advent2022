package stock

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type stack struct {
	stack []string
	top   int
}

func (s *stack) pop() (string, error) {
	if s.top == -1 {
		return "", errors.New("empty stack")
	}
	res := s.stack[s.top]
	if res == "" {
		return "", errors.New("unconditioned stack")
	}
	s.stack = s.stack[:s.top]
	s.top--
	return res, nil
}

func (s *stack) push(i string) {
	s.stack = append(s.stack, i)
	s.top++
}

func (s *stack) head() string {
	if s.top == -1 {
		return ""
	}
	return s.stack[s.top]
}

func (s *stack) bottomPop(v string) {
	st := make([]string, 1, s.top+2)
	st[0] = v
	st = append(st, s.stack...)
	s.stack = st
	s.top = len(s.stack) - 1
}

func newStack() *stack {
	return &stack{nil, -1}
}

type stock struct {
	columns []*stack
}

func (s *stock) move(count, from, to int) error {
	from--
	to--
	for i := 0; i < count; i++ {
		v, err := s.columns[from].pop()
		if err != nil {
			return err
		}
		s.columns[to].push(v)
	}
	return nil
}

func (s *stock) moveOrdered(count, from, to int) error {
	from--
	to--
	pack := make([]string, count)
	for i := count - 1; i >= 0; i-- {
		v, err := s.columns[from].pop()
		if err != nil {
			return err
		}
		pack[i] = v
	}
	for _, v := range pack {
		s.columns[to].push(v)
	}
	return nil
}

func (s *stock) insert(v string, to int) {
	s.columns[to].bottomPop(v)
}

func (s *stock) getHead() string {
	res := ""
	for _, st := range s.columns {
		res += st.head()
	}
	return res
}

func newStock(cols int) *stock {
	s := stock{columns: make([]*stack, cols)}
	for i := 0; i < cols; i++ {
		s.columns[i] = newStack()
	}
	return &s
}

func buildStock(s []string, cols int) *stock {
	st := newStock(cols)
	for _, row := range s {
		if !strings.ContainsRune(row, '[') {
			continue
		}
		for i := 0; i < cols; i++ {
			pos := 1 + 4*i
			if pos > len(row) {
				break
			}
			symbol := row[pos]
			if symbol == ' ' {
				continue
			}
			st.insert(string(symbol), i)
		}
	}
	return st
}

func parseMoveCmd(s string) (int, int, int, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 6 {
		return 0, 0, 0, errors.New("bad format")
	}
	cnt, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, 0, 0, errors.New("cant parse count")
	}
	from, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		return 0, 0, 0, errors.New("cant parse from")
	}
	to, err := strconv.ParseInt(parts[5], 10, 64)
	if err != nil {
		return 0, 0, 0, errors.New("cant parse to")
	}
	return int(cnt), int(from), int(to), nil
}

func moveStock(s []string, f func(int, int, int) error) error {
	for _, line := range s {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		cnt, from, to, err := parseMoveCmd(line)
		if err != nil {
			return errors.Wrapf(err, "line: <%v>", line)
		}
		err = f(cnt, from, to)
		if err != nil {
			return errors.Wrapf(err, "command: <%v>", line)
		}
	}
	return nil
}

func StockOps(s []string) (string, error) {
	edge := 0
	for i, line := range s {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			edge = i
			break
		}
	}
	st := buildStock(s[:edge], 9)
	err := moveStock(s[edge:], st.move)
	if err != nil {
		return "", errors.Wrap(err, "move stock")
	}
	return st.getHead(), nil
}

func Stock9001(s []string) (string, error) {
	edge := 0
	for i, line := range s {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			edge = i
			break
		}
	}
	st := buildStock(s[:edge], 9)
	err := moveStock(s[edge:], st.moveOrdered)
	if err != nil {
		return "", errors.Wrap(err, "move stock")
	}
	return st.getHead(), nil
}
