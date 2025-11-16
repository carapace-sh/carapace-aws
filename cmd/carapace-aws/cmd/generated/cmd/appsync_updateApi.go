package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_updateApiCmd = &cobra.Command{
	Use:   "update-api",
	Short: "Updates an `Api`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_updateApiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_updateApiCmd).Standalone()

		appsync_updateApiCmd.Flags().String("api-id", "", "The `Api` ID.")
		appsync_updateApiCmd.Flags().String("event-config", "", "The new event configuration.")
		appsync_updateApiCmd.Flags().String("name", "", "The name of the Api.")
		appsync_updateApiCmd.Flags().String("owner-contact", "", "The owner contact information for the `Api`.")
		appsync_updateApiCmd.MarkFlagRequired("api-id")
		appsync_updateApiCmd.MarkFlagRequired("name")
	})
	appsyncCmd.AddCommand(appsync_updateApiCmd)
}
