# gobrainfuck

> Interpreter for `BrainFuck` language impl with go.



## BrainFuck Language

| Token | Explain                                                      |
| ----- | ------------------------------------------------------------ |
| >     | Pointer moves one square to the right                        |
| <     | Pointer moves one square to the left                         |
| +     | Add 1 to the byte value of the current cell of the pointer   |
| -     | Sub 1 to the byte value of the current cell of the pointer   |
| ,     | Read 1 byte from input and store it in the current cell of the pointer |
| .     | Display the byte value of the current cell of the pointer    |
| [     | When the current value of the pointer is 0, the program jumps to the corresponding ']'; otherwise, the program continues |
| ]     | The program jumps back to the corresponding '[' place        |



## Usage

```shell
gobrainfuck '${brainfuck_code}' '${input}'
```

**Note: this program does not support line breaks.**



For example:

1. Hello World!

   ```shell
   gobrainfuck '>+++++++++[<++++++++>-]<.>+++++++[<++++>-]<+.+++++++..+++.>>>++++++++[<++++>-]<.>>>++++++++++[<+++++++++>-]<---.<<<<.+++.------.--------.>>+.>++++++++++.'
   ```

2. Outputs square numbers from 0 to 10000

   ```shell
   gobrainfuck '++++[>+++++<-]>[<+++++>-]+<+[>[>+>+<<-]++>>[<<+>>-]>>>[-]++>[-]+>>>+[[-]++++++>>>]<<<[[<++++++++<++>>-]+<.<[>----<-]<]<<[>>>>>[>>>[-]+++++++++<[>-<-]+++++++++>[-[<->-]+[<<<]]<[>+<-]>]<<-]<<-]'
   ```

   

