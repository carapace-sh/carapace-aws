package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getApiCmd = &cobra.Command{
	Use:   "get-api",
	Short: "Retrieves an `Api` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_getApiCmd).Standalone()

		appsync_getApiCmd.Flags().String("api-id", "", "The `Api` ID.")
		appsync_getApiCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_getApiCmd)
}
