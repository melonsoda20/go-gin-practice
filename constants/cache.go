package constants

type cacheConst struct {
	GetAllToDoKey string
}

var CacheKeys = cacheConst{
	GetAllToDoKey: "get_all_todo",
}
