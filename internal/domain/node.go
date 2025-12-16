package domain

type Node string

const (
	Client Node = "client"
	Router Node = "router"
	Server Node = "server"
)
