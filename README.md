### MathOp client
MathOp is CLI client for the [mathOp service](https://github.com/bigrewal/mathOp_service). It is used to execute simple math operations like additon and square-root. Check the usage instructions below.

##### How to run it?
1. Clone the repo and place it in the `{GOPATH}/src directory`
2. `cd mathOp`
3. Go to `/config/config.yaml` and update the host and port number of the backend mathOp service, if required.
4. Execute `go build`. (This step not required if you haven't made any changes to the code, just go to step 5).
5. Copy the `mathOp` executable inside `/usr/local/bin`. 
6. go in terminal and start using it. For example execute `mathOp add 1 2`. Checkout the usage section below.

##### Problems faced.
1. Couldn't get `go install` to work.
2. Couldn't get goMock to work in order to unit test the mathOp service client.

#### Usage
`mathOp` only supports:
    a. Addition of 2 numbers.
    b. Square-root of a number

1. `mathOp add {num1} {num2}` will add num1 and num2.
2. `mathOp sqrt {num}` will perform the square root of the number. (Decimal point numbers will work aswell)
