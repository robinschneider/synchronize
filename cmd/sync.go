/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/google/go-containerregistry/pkg/crane"
	"log"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: CopyImages,
}

func CopyImages(cmd *cobra.Command, args []string) {
	fmt.Println("Copy image ...")
	fmt.Printf("%s", cmd.Flag("tag").Value)
	err := crane.Copy("source", "destination")
	if err != nil {
		log.Fatal(err.Error())
	}
}

var Source string
var Tags []string

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().StringSliceVarP(&Tags, "tag", "t", nil, "Tag to sync, you can use this flag multiple times")
	syncCmd.Flags().StringVarP(&Source, "source", "s", "", "Description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
