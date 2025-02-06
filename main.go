/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Dindonpingpong/todo-cli/cmd"
	storage "github.com/Dindonpingpong/todo-cli/storage/filestorage"
)

func main() {
	store := storage.NewStorage()
	cmd.Execute(store)
}
