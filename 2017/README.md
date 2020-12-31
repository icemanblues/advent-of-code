# 2017

This was my first advent of code. I didn't really know what to expect. 

## Java

In 2017, Java was my preferred programming language.
Java may not be the best language, but it has been so reliable for me. I invested so much time learning the language,
the core libraries, the 3rd party libraries, and the greater community and culture.

I started off the year using Java. And quickly it became quite tedious to `public static void main` all the time.
The earlier days in 2017 really just needed one function to solve it all. And java brought so much boiler plate and ceremony,
in order to execute one function. Within a week or two, I switched to python.

## Python 

I really enjoy python, the programming language. With so little boilerplate, the language is very expressive.
My favorite thing is that it is perfect for smaller tasks and scripting glue. The community is excellent.
And I **LOVE** the zen of python. These guidelines make it a lot easier to write good code, share good code,
and to do it quickly.

## Redo

Since this was my first AoC, I didn't think to create helper functions. In fact, some of the scripts
where targeted to solving my puzzle inputs that it wasn't very scalable and re-usable.

Also, and more importantly, I had split the puzzles across multiple laptops and I didn't commit all of solutions.
So this repo is quite bare, although I do have all 50 stars for 2017. I want to return and provide the solutions
for all of the puzzles. This will allow me to follow my best practices learned from completing additional AoC events.

The challenges are given their own folder, their own module. There is no `util` or `common` module for this year.
I am aiming to keep the individuals puzzles as independent as possible, no 3rd party packages either.

That said, there is one puzzle that relies on a previous solution. `day14` has a direct dependency on `day10`. 
The easiest way to reuse that code is to set your `PYTHON_PATH` to the `2017` directory.

```
export PYTHON_PATH="advent-of-code/2017"
```

## Experiments and New Learnings

Since this is a redo, I am also using this as an opportunity to learn new things. 

### Type Hinting

I am using python's type hinting in all major functions. And for collections (List, Set, Tuple, Dict, etc).
This is to get more experience with python's typing system.

For LSP, using [pyright](https://github.com/microsoft/pyright) as it is the current best one with support for
type hinting. Unfortunately, at this time, it does not support formatting. Formatting will be done with `autopep8`.

For static typing, [mypy](https://github.com/python/mypy) to  confirm that all python modules are passing their type hinting checks.
