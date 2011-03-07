Wordfinder let's you find words in Scrabble and other word games. It comes packaged with a word list I found somewhere on the internet (twl.txt). 

To start it, run `make run` on a shell. It starts a server on port 8080.

If you navigate to `localhost:8080`, you're all set. 

Just start typing your available letters, and it returns all possible valid words. For instance, entering `hello` will yield:

    hello
    lo
    oe
    oh
    he
    ho
    hole
    hoe
    ole
    el
    helo
    hell
    ell

Optionally, you can pass a constraint as a regular expression. For instance, if you have the letters 'hello' but want the letters 'h' and 'l' separated by two spaces, just type: `hello[h..l]`:

    hello
    hell

No, I did not create this program to run on my mobile phone while I play Scrabble on the weekends :)

Warning: depending on the word list this program might consume a lot of memory. Also, searches may get slow after about 14 letters, depending your CPU. The word finder is a fairly basic implementation of BFS. 