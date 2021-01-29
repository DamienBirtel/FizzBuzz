# FizzBuzz



This is a simple implementation of FizzBuzz in a REST API.

I chose REST to build it because it is widely known and well understood by the large majority of programmers.
So it can be easy and frictionless to integrate to a bigger system of services and APIs.

This project is production ready in the sense that it will do what's expected of it, is easy to read and understand,
and also easy to modify.


## Installation

First, make sure you have the [GO](https://golang.org/dl/) language installed on your computer.

Then just clone the repository.

```bash
git clone https://github.com/DamienBirtel/FizzBuzz
```

Build the executable (will be named FizzBuzz by default), and run it.

```bash
go build
./FizzBuzz
```


## Usage

You can either send a GET or a POST request to the default address "localhost:9090"

Example of a GET request using curl:

Type in your terminal
```bash
curl http://localhost:9090
```

The result expected is the list of numbers from 1 to 200 being printed in your terminal,
with multiples of 7 being replaced by Fizz, multiples of 9 being replaced by Buzz,
and multiples of both being replaced by FizzBuzz.


If you want to customize the output, you can send a POST request with a JSON object
containing the custom parameters.
Here's an example of how the body should be formatted using the default parameters:

```JSON
{
    "length":200,
    "fizznum":7,
    "buzznum":9,
    "fizzword":"Fizz",
    "buzzword":"Buzz"
}
```

You can omit customizing some parameters as long as you send a valid JSON. The missing parameters will be replaced by the default ones.
So finally, here is an example of a custom POST request using curl:

```bash
curl -d '{"length":15,"fizznum":3,"buzzword":"Toto"}' -X POST http://localhost:9090
```
