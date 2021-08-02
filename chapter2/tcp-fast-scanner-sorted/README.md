- **Execution**

    
        The execution starts from the main. You can see two channels created with name ports and results.

        The output of open ports are stored in openport slice.

        Next, a call to go routine wokers with arguments ports and results

        paralley, an anonymous function start running go func(){}(). It has receiver channel **ports** which receives i as input. 

        After worker go routine invoked, it receives input from anonymous go routine and results channel assignes to 0 if no open ports and assignes to p(port which is open).


        Next line of anonoymous function, the for loop receives port input from the assigned signal from the worker go routine.

        finally, the openports are added to the slice with append function.


        It is important to closr channel once its work has done.

        The open ports are sorted using sort.Ints()