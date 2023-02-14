This is the fastest implementation of merge sort among all languages, which is, of course, implemented in Go.

Obviously, it is subjective since I don't know all programming languages. However, I would be very surprised if the above statement is not valid.

The average time for sorting an array with randomized values that I was able to achieve on MacBook Pro (M1, 2020) while employing all 8 performance cores of the M1 chip:

    2.332401 seconds

I encourage all the programmers that are reading this to try beating the record using a language that they think is faster.



DISCLAIMERS!!!

1. Computers with more and stronger CPU cores will definitely demonstrate faster sorting times. Therefore to keep the hardware parameters the same, run the experiment on a MacBook Pro with an M1 chip.

2. There is a sorting algorithm called bucket sort, whose time complexity turned out to be O(n+k) compared to merge sorts (n*log(n)). In the current experiment, the focus is more to check the capabilities of the programming language and how much concurrency improves the performance of a program rather than the performance of an algorithm. The reason why merge sort was chosen for this experiment is that merge sort is a divide-and-conquer sorting algorithm that perfectly fits the purpose of the experiment.

3. The whole experiment took about an hour and pretty heated up the computer, hence I strongly recommend tracking the temperature of your computer. I choose parameters (size of the array = 100,000,000 and the number of iterations = 100) that are optimal for my machine. Maybe, you will have to experiment a bit to find out what are optimal ones in your case.



In order to run the experiment on your own machine, clone the project, enter the directory of the project and run the command:

    $ go run main.go

To run the command, Go programming language should have been installed, which can be downloaded from https://go.dev/dl

After the program finishes, you should be able to see three new files:

    multi.txt     (results when utilizing all the available CPU cores of your machine)

    single.txt    (results when using only one core of the CPU)

    builtin.txt   (results of sort.Ints() of the "sort" package)

Give it a try, compare with the results that I got. I assure you you will get tons of fun ;)
