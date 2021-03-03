package main

import "git.fuyu.moe/Fuyu/router"

func authMiddleware(f router.Handle) router.Handle {
	return f
}
