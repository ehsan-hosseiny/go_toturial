package utils

import (
	"context"
	"fmt"
)

func GetUserId(ctx context.Context) {
	fmt.Println(ctx.Value("user_id"))

}
