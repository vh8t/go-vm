package memory

import "fmt"

type Memory struct {
	data []byte
	size uint64
}

func NewMemory(size uint64) *Memory {
	return &Memory{
		data: make([]byte, size),
		size: size,
	}
}

func (m *Memory) LoadProgram(program []byte) {
	copy(m.data, program)
}

func (m *Memory) Read(address uint64) (byte, error) {
	if address < m.size {
		return m.data[address], nil
	}
	return 0, fmt.Errorf("Address out of bounds: read @ 0x%016X (%d)", address, address)
}

func (m *Memory) ReadBytes(address uint64, size uint64) ([]byte, error) {
	if address+size <= m.size {
		return m.data[address : address+size], nil
	}
	return []byte{}, fmt.Errorf("Address out of bounds: readBytes @ 0x%016X (%d)", address+size, address+size)
}

func (m *Memory) Write(address uint64, value byte) error {
	if address < m.size {
		m.data[address] = value
		return nil
	}
	return fmt.Errorf("Address out of bounds: write @ 0x%016X (%d)", address, address)
}
