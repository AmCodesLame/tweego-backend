package db

import (
	"context"
	"fmt"
)

func Connect() (clnt *PrismaClient, cntxt context.Context, err error) {
	client := NewClient()
	if err := client.Prisma.Connect(); err != nil {
		fmt.Println("[Prisma] Error: Cannot Connect")
		return nil, nil, err
	}
	ctx := context.Background()

	return client, ctx, err
}
