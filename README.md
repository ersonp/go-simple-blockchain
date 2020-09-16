# go-simple-blockchain

## v0.1.0 core

This tutorial is adapted from this excellent [post](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc) about writing a basic blockchain using Go.<br>
This is the very basic and barebones code where the core fundamentals will be built upon and other functionality will be added later.

## v0.2 p2p

In this version the libp2p library is added to the core. Then previous version was basically a single node server.<br>
This is addapted from this [post](https://medium.com/@mycoralhealth/code-a-simple-p2p-blockchain-in-go-46662601f417) where there is no consensus (It will be added later) and has a simple p2p blockchain structure.<br>
Note: The go.mod file is taken from [here](https://github.com/libp2p/go-libp2p-examples/blob/master/go.mod) as this p2p code is based on the examples provided by the go-libp2p library and they are old. So in order for our code to work for now we will be using the old versions (This will also be updated later).

### v0.2.0

Here the blocks are added via the command prompt and the gorilla/mux code is not in use.


### v0.2.1

Here the blocks are now added via the API made in gorilla/mux removing the input capability from the terminal hereafter. There is still no consensus algorithm. The node simply adds a block to it's blockchain and shares it's copy to other nodes where the other nodes simply compare the lenth of their blockchain and the recieved one and replace it with the longest.