#todo

### General 
#### Layers (from low to high level)
==Note: In most cases, adjacent levels are combined into larger ones (e.g [[TcpIP|TCP/IP]])==
- ***Physical***:
	Doesn't analyse/modify data, simply moves bits to destination as radio waves, electricity, etc. ^6a70f7
- ***Data link***:
	Manages low-level messages from higher level and the link, used by Physical layer ^d5d7c9
- ***Network***:
	Addressing, routing  ^15d3de
- ***Transport***:
	Network equipment abstraction. Linking and data transfer between processes on different hosts with some extra logic or guarantees  ^fabd59
- ***Session***:
	Link sessions management  ^13fa21
- ***Presentation***:
	Data syntax and semantics management, encryption.  ^224a8c
- ***Application***: 
	Application, that uses network as functionality (email, web-pages, ftp servers, etc.) ^562a05


### Protocol suites
- [[TcpIP|TCP/IP]]
- [[HTTP]]
- [[HTTP 2.0]]