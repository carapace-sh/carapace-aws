package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_updateGlobalTableCmd = &cobra.Command{
	Use:   "update-global-table",
	Short: "Adds or removes replicas in the specified global table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_updateGlobalTableCmd).Standalone()

	dynamodb_updateGlobalTableCmd.Flags().String("global-table-name", "", "The global table name.")
	dynamodb_updateGlobalTableCmd.Flags().String("replica-updates", "", "A list of Regions that should be added or removed from the global table.")
	dynamodb_updateGlobalTableCmd.MarkFlagRequired("global-table-name")
	dynamodb_updateGlobalTableCmd.MarkFlagRequired("replica-updates")
	dynamodbCmd.AddCommand(dynamodb_updateGlobalTableCmd)
}
