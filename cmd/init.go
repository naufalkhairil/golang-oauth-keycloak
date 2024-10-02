/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/naufalkhairil/golang-oauth-keycloak/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("Input email: ")
		var email string
		_, err := fmt.Scanln(&email)
		if err != nil {
			return errors.Wrap(err, "Error reading input 'email'")
		}

		fmt.Println("")
		client.InitClient(email)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
