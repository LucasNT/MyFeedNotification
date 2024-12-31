/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/LucasNT/MyFeed/internal/adapters/badger"
	gofeed "github.com/LucasNT/MyFeed/internal/adapters/go_feed"
	"github.com/LucasNT/MyFeed/internal/adapters/interfaces"
	notifysend "github.com/LucasNT/MyFeed/internal/adapters/notifySend"
	usecase "github.com/LucasNT/MyFeed/internal/useCase"
	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:        "MyFeed [flags] rss_link",
	Short:      "Get feed and show a notification",
	Run:        commandExecute,
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"rss_link"},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var link string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.MyFeed.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func commandExecute(cmd *cobra.Command, args []string) {
	dbPath := xdg.DataHome + "/myFeed/database"
	err := os.MkdirAll(dbPath, 0700)
	if err != nil {
		log.Fatal(err)
	}

	baseCtx, baseCtxCancel := context.WithCancel(context.Background())
	defer baseCtxCancel()
	link := args[0]
	var getFeed interfaces.FeedGetter
	var notificationSender interfaces.NotificationSender
	var validateNewFeed interfaces.ValidateNewFeed

	getFeed = gofeed.NewGoFeed()
	notificationSender = notifysend.New()
	aux, err := badger.New(dbPath)
	defer aux.Close()
	validateNewFeed = aux

	executor := usecase.NewGetFeedSaveAndNotify(getFeed, validateNewFeed, notificationSender)

	err = executor.Execute(baseCtx, link)

	if err != nil {
		fmt.Println(err)
	}
}
