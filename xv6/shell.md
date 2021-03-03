xv6 Shell
=========

Overview
--------

The basic logic of shell is 

1. read input
2. fork a process and run command

Not all the input commands are external binary to be executed. For example, the `cd` command is built-in shell 
which update the work directory.

Work directory is a concept in OS rather than the relative path in shell. Change the dir with `chdir` system call.

Input EOF (Ctrl-D) to quit the shell.


Code analysis
-------------

1. lines 1-258 are the core implement of run commands
2. lines 262-493 are the logic parse the input string into command


New Process
-----------

`fork()` return different value for parent process and child process

1. pid of child process to parent process
2. 0 to child process

`wait()` wait until any of the children process terminates


Command
-------

`cmd` is the abstract type of command. It has multiple implementation

1. exec: argv, eargv (env)
2. redir: cmd, file (output), efile (end of the file name), mode, fd
3. pipe: left cmd, right cmd
4. list: left cmd, right cmd
5. back: cmd

`*cmd` is converted to concrete type of cmd via type convert

Note: the sub-command can be any type of commands

```
struct execcmd* ecmd = (struct execcmd*)cmd;
```

Run Command
-----------

### exec

call `exec(argv[0], argv)`. The eargv is not used yet.

### redir

open `output` file in specified `mode` as stdout and run `cmd`

### pipe

generate input/output pipe with system call `pipe` firstly. Create two child processes

* For process A: bind the output to the writer of pipe, run the left command
* For process B: bind the input to the reader of pipe, run the right command
* For parent process: wait until both children processes quit

### list

There are two command (or more). Run the left command in subprocess. Then run the right command regardless the status of left command

```
date; echo "end"
```

### back

run `cmd` in child process. Parent process will not wait for the finish of the child process.

So we can not find teh child process status again. How to manage the job?


Parse Command
-------------

### peek

```
int peek(char **ps, char *es, char *toks)
```

skip the whitespace and return the string starts with any characters in `toks`.
The `ps` pointer is updated to the begin of the string (might be different with the return `int` value

### parseexec

```
struct cmd* parseexec(char **ps, char *es)
```

1. build the redirect command if needed 
2. construct the subcommand for execcmd

Question:

1. why try to parse redir command in each round of parsing

#### sample 1

```
grep 192.168.0 /etc/hosts > hosts.txt
```

parsed to

```
redirect(execcmd("grep", "192.168.0", "/etc/hosts"), "hosts.txt", O_WRONLY | O_CREATE | O_TRUNC, 1)
```


#### sample 2


```
grep 192.168.0 /etc/hosts > hosts.txt > result.txt
```

parsed to

```
redirect(redirect(execcmd("grep", "192.168.0", "/etc/hosts"), "hosts.txt", O_WRONLY | O_CREATE | O_TRUNC, 1),
  "result.txt, O_WRONLY | O_CREATE | O_TRUNC", 1)
```

### parseblock

```
struct cmd* parseblock(char **ps, char *es)
```

```
(cat /etc/hosts | grep -v localhost) > abc.txt
```


### parseline

```
struct cmd* parseline(char **ps, char *es)
```

Call `gettoken(ps, es, 0, 0)` firstly so the string in `[ps .. es]` is a valid token (symbol, or normal word).

`parseline`

### parsepipe

```
struct cmd* parsepipe(char **ps, char *es)
```

Can have multiple pipelines

### gettoken

```
int gettoken(char **ps, char *es, char **q, char **eq)
```

* return the first valid character value in int
  * 'a': normal token
  * symbol: `|();&<+`, `+` means `>>`
* `ps` is update to point to the first char of next token ()
* `q`, if not null, is assigned with the pointer to first non-whitelist character of `ps`
* `eq`, if not null, will be assigned with the pointer to the end of the token


### parseredirs

```
struct cmd* parseredirs(struct cmd *cmd, char **ps, char *es)
```

1. search for `<>` 
2. tok is either `<` or `>` or `+` (short for `>>`)
3. `[q .. eq]` is the next symbol

For example, 

```
cat "hello world" >> abc.txt
```

will be parsed as

```
redircmd(cmd, "abc.txt", O_WRONLY|O_CREATE, 1)
```

```
cat < /etc/hosts
```

will be parsed as

```
redircmd(cmd, "/etc/hosts", O_RDONLY, 0)
```
