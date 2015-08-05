# gouss

Gouss is a little script which triangulates matrices, doing row operations **concurrently**. I made this in order to **learn** the basics of **Go**. The code itself is pretty clumsy, but I guess that's part of the learning process. 

#Running gouss

Just copy and build **gouss** in the src folder of your GOPATH.

```bash
cd $GOPATH/src
git clone https://github.com/guidoarnone/gouss
cd gouss
go build
```

After that, you can execute gouss by doing ```./gouss```. To display the options and input format guide, you can do ```./gouss -h```.
