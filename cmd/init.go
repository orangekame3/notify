/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize notify",
	RunE: func(cmd *cobra.Command, args []string) error {
		currentUser, err := user.Current()
		if err != nil {
			return err
		}
		notifyPath := filepath.Join(currentUser.HomeDir, ".config", "notify")
		if _, err := os.Stat(notifyPath); os.IsNotExist(err) {
			if err = os.MkdirAll(notifyPath, 0755); err != nil {
				return err
			}
			fmt.Println(".config/notify directory created")
		} else {
			fmt.Println("Already initialized")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
