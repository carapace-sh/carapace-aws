package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_getSchemaCreationStatusCmd = &cobra.Command{
	Use:   "get-schema-creation-status",
	Short: "Retrieves the current status of a schema creation operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_getSchemaCreationStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appsync_getSchemaCreationStatusCmd).Standalone()

		appsync_getSchemaCreationStatusCmd.Flags().String("api-id", "", "The API ID.")
		appsync_getSchemaCreationStatusCmd.MarkFlagRequired("api-id")
	})
	appsyncCmd.AddCommand(appsync_getSchemaCreationStatusCmd)
}
