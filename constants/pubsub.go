package constants

type pubSubConst struct {
	CreateToDoTopicID string
}

var PubSubKeys = pubSubConst{
	CreateToDoTopicID: "create-todo-topic",
}
