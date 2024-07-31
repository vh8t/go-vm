section .data
  filename db "text.txt", 0
  prompt db "Enter a message: ", 0

section .bss
  input resb 128

section .text
  global _start

_start:
  mov %0, 1       ; sys_write
  mov %1, 1       ; stdout
  mov %2, prompt
  mov %3, 17
  syscall

  mov %0, 0       ; sys_read
  mov %1, 0       ; stdin
  mov %3, input
  mov %4, 128
  syscall

  mov %0, 2         ; sys_open
  mov %1, filename
  mov %2, 577       ; flags
  mov %3, 420       ; mode
  syscall

  mov %15, %0   ; move fd

  mov %0, input
  mov %1, 0

counter:
  mov %2, [%0]
  cmp %2, 0
  je continue
  inc %0
  inc %1
  jmp counter

continue:
  mov %14, %1   ; move input size

  mov %0, 1     ; sys_write
  mov %1, %15   ; move fd
  mov %2, input
  mov %3, %14   ; move input size
  syscall

  mov %0, 3   ; sys_close
  mov %1, %15 ; move fd
  syscall

  hlt
