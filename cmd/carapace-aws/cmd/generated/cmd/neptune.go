package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneCmd = &cobra.Command{
	Use:   "neptune",
	Short: "Amazon Neptune\n\nAmazon Neptune is a fast, reliable, fully-managed graph database service that makes it easy to build and run applications that work with highly connected datasets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneCmd).Standalone()

	rootCmd.AddCommand(neptuneCmd)
}
