--------Near Term Scope----------

var a = 5

var b = 6

var c = 'string one'

var d = 'string two'

var e = true

func addition(numOne numTwo) [
  var sum = add(numOne numTwo)
  print(sum)
]

var f = func trueOrFalse(booleanValue) [
  return true
]

print(add(5 10))

print(stringOne stringTwo)

add(5 6)

trueOrFalse(e)



-----Future------

// This is a comment. No multiline comments


class Greeting [

  var greeting = 'Hello!'

  func new(greet) [
    this.greeting = greet
  ]

  func hello(name) [

    if (name == null) [

      print('Hello world!')

    ] else [

      print('Hello' name '!')

    ]

  ]

]