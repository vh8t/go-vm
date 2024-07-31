package cpu

import (
	"encoding/binary"
	"os"
	"syscall"
)

/*
Syscalls

0 - sys_read
1 - sys_write
2 - sys_open
3 - sys_close
4 - sys_stat
*/

func (c *CPU) SysRead() {
	fd, _ := c.ReadRegister(1)
	buf, _ := c.ReadRegister(2)
	count, _ := c.ReadRegister(3)

	buffer := make([]byte, count)

	bytesRead, err := syscall.Read(int(fd), buffer)
	Error(err)

	for i := 0; i < bytesRead && i < int(count); i++ {
		c.memory.Write(uint64(buf)+uint64(i), buffer[i])
	}
}

func (c *CPU) SysWrite() {
	fd, _ := c.ReadRegister(1)
	buf, _ := c.ReadRegister(2)
	count, _ := c.ReadRegister(3)

	buffer, err := c.memory.ReadBytes(uint64(buf), uint64(count))
	Error(err)

	switch fd {
	case 0:
		_, err = os.Stdin.Write(buffer)
	case 1:
		_, err = os.Stdout.Write(buffer)
	case 2:
		_, err = os.Stderr.Write(buffer)
	default:
		_, err = syscall.Write(int(fd), buffer)
	}

	Error(err)
}

func (c *CPU) SysOpen() {
	filename, _ := c.ReadRegister(1)
	flags, _ := c.ReadRegister(2)
	mode, _ := c.ReadRegister(3)

	var stringified string
	for {
		b, _ := c.memory.Read(uint64(filename))
		if b == 0 {
			break
		}
		stringified += string(b)
		filename++
	}

	fd, err := syscall.Open(stringified, int(flags), uint32(mode))
	Error(err)

	err = c.LoadRegister(0, int64(fd))
	Error(err)
}

func (c *CPU) SysClose() {
	fd, _ := c.ReadRegister(1)
	err := syscall.Close(int(fd))
	Error(err)
}

func (c *CPU) SysStat() {
	filename, _ := c.ReadRegister(1)
	statbuf, _ := c.ReadRegister(2)

	var stringified string
	for {
		b, _ := c.memory.Read(uint64(filename))
		if b == 0 {
			break
		}
		stringified += string(b)
		filename++
	}

	var buf []byte
	var stat syscall.Stat_t

	err := syscall.Stat(stringified, &stat)
	Error(err)

	appendUint64 := func(v uint64) {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, v)
		buf = append(buf, b...)
	}

	appendUint32 := func(v uint32) {
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, v)
		buf = append(buf, b...)
	}

	appendInt64 := func(v int64) {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(v))
		buf = append(buf, b...)
	}

	appendInt32 := func(v int32) {
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, uint32(v))
		buf = append(buf, b...)
	}

	appendTimespec := func(t syscall.Timespec) {
		appendInt64(t.Sec)
		appendInt64(t.Nsec)
	}

	appendUint64(stat.Dev)
	appendUint64(stat.Ino)
	appendUint64(stat.Nlink)
	appendUint32(stat.Mode)
	appendUint32(stat.Uid)
	appendUint32(stat.Gid)
	appendInt32(stat.X__pad0)
	appendUint64(stat.Rdev)
	appendInt64(stat.Size)
	appendInt64(stat.Blksize)
	appendInt64(stat.Blocks)
	appendTimespec(stat.Atim)
	appendTimespec(stat.Mtim)
	appendTimespec(stat.Ctim)
	for _, v := range stat.X__unused {
		appendInt64(v)
	}

	for i, b := range buf {
		err = c.memory.Write(uint64(statbuf)+uint64(i), b)
		Error(err)
	}
}
