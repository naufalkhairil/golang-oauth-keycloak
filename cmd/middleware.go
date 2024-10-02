/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/naufalkhairil/golang-oauth-keycloak/middleware"
)

// middlewareCmd represents the middleware command
var middlewareCmd = &cobra.Command{
	Use:   "middleware",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {

		middleware.InitMiddleware()

		r := gin.Default()

		r.GET("/", middleware.Home)
		r.GET("/auth/v1", middleware.Auth)
		r.GET("/auth/v1/callback", middleware.AuthCallback)
		r.GET("/auth/v1/token/:email", middleware.GetToken)
		r.GET("/success", middleware.Success)

		r.Run(":8181")
	},
}

func init() {
	rootCmd.AddCommand(middlewareCmd)
}
