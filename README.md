# Firecrest 🔥

***As Powerful as Node.js, as Slick as Wren***

## What ❔

⚙️ A modern programming language with a C-like syntax, inspired mainly by Wren but also JavaScript and Go.

⚙️ Reduce visual noise, and make programming feel cleaner to read and easier to write.

⚙️ Meant to be a clean little scripting language, powered by a rich ecosystem.

⚙️ Have a standard library small enough to learn in a day, but large enough to be used in production.

⚙️ Be extensible, with the ability to implement your own Node.js functions as Fircrest functions.

## Simple Examples 💻

Hello World (esque):
```go

func hello(name) [

  print('Hello' name '!')

]

var name = 'Nick'

hello(name)

```

Equal To:
```go

func isEqual(valueOne valueTwo) [

  if ( valueOne == valueTwo) [

    print('They are equal 🎉')

  ] else [

    print('They are not equal 😔')

  ]

]

isEqual(true false)

```

## Implementation 🔨

⚙️ Lexing - [moo.js](https://github.com/no-context/moo)

⚙️ Parsing - [nearley.js](https://github.com/kach/nearley)

⚙️ Interpreting - [interpreter.js](https://github.com/firecrest-lang/firecrest/blob/main/interpreter.js)
