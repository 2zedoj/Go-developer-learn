# Challenge: Ping Pong

Write a code in Go using all the knowledge acquired so far! And in this code you will need, based on our concurrency classes, to use channels and goroutines so that your program displays, in alternation, the words ping and pong.



## Learnings

To solve this project, I used the resources taught during previous classes, such as concepts of concurrency, with goroutines and channels, functions and packages.
Below I will list the concepts learned and how they were incorporated into the project to solve this problem.

  **1° Step** :  We have created two channels, which will be used to transmit the "ping" and "pong" 
  

    c1  :=  make(chan  string)
    c2  :=  make(chan  string)

  **2° Step** :  We created the ping and pong functions, both of which receive channels, in pointer forms. In ping, it inserts the message into the channel, waits and only continues until there is a message in the pong channel. In the pong function, we wait for some message in the ping channel, as soon as we have it, we show the result and then sleep for 2 seconds, after that it prints the "pong" message and sends this message to the pong channel. 
  

    func  ping(pingCh  *chan  string, pongCh  *chan  string){
	    for {
		    *pingCh  <-  "ping"
		    time.Sleep(1  *  time.Second)
		    <-  *pongCh
		}
	}
	func  pong(pingCh  *chan  string, pongCh  *chan  string){
		for {
			msg  :=  <-  *pingCh
			fmt.Println(msg)
			time.Sleep(1  *  time.Second)
			fmt.Println("pong")
			*pongCh  <-  "pong"
		}
	}

  **3° Step** : We call the ping function and then pong it with the goroutines
  
    go  ping(&c1, &c2)
    go  pong(&c1, &c2)

 **4° Step** : Finally, I implemented a system that as soon as the user presses enter, the program will end, with the help of Scanln.
 

    var  entrada  string
    fmt.Scanln(&entrada)