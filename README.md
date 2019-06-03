#### Math Op client.

##### How to run it?

1. `cd mathOp`
2. Go to `/config/config.yaml` and update the host and port number of the backend mathOp service, if required.
3. Execute `go build` and copy the `mathOp` executable inside `/usr/local/bin`
4. go in terminal and start using it. For example execute `mathOp add 1 2`. Checkout the usage section below.

##### Problems faced.

1. Couldn't get `go install` to work.
2. Couldn't get goMock to work in order to unit test the mathOp service client.

#### Usage

`mathOp` only supports Addition and square-root.

1. `mathOp add {num1} {num2}` will add num1 and num2.
2. `mathOp sqrt {num}` will perform the square root of the number.
