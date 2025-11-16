package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_createApiCmd = &cobra.Command{
	Use:   "create-api",
	Short: "Creates an `Api` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_createApiCmd).Standalone()

	appsync_createApiCmd.Flags().String("event-config", "", "The Event API configuration.")
	appsync_createApiCmd.Flags().String("name", "", "The name for the `Api`.")
	appsync_createApiCmd.Flags().String("owner-contact", "", "The owner contact information for the `Api`.")
	appsync_createApiCmd.Flags().String("tags", "", "")
	appsync_createApiCmd.MarkFlagRequired("name")
	appsyncCmd.AddCommand(appsync_createApiCmd)
}
