package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsyncCmd = &cobra.Command{
	Use:   "appsync",
	Short: "AppSync provides API actions for creating and interacting with data sources using GraphQL from your application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsyncCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsyncCmd).Standalone()

	})
	rootCmd.AddCommand(appsyncCmd)
}
