package main

import (
	"context"

	"github.com/ehsanhossini/go/go_tutorial/context/utils"
)

func main() {
	ctx := context.Background()

	ctxUser := context.WithValue(ctx, "user_id", 10)

	utils.GetUserId(ctxUser)

}
