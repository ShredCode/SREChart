package ast

type Flow struct {
  Start string
  Options  []Option
  End string
}

type Option stuct {
  Name string
  Decision string
}
