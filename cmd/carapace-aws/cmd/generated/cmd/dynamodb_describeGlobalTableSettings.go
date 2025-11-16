package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeGlobalTableSettingsCmd = &cobra.Command{
	Use:   "describe-global-table-settings",
	Short: "Describes Region-specific settings for a global table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeGlobalTableSettingsCmd).Standalone()

	dynamodb_describeGlobalTableSettingsCmd.Flags().String("global-table-name", "", "The name of the global table to describe.")
	dynamodb_describeGlobalTableSettingsCmd.MarkFlagRequired("global-table-name")
	dynamodbCmd.AddCommand(dynamodb_describeGlobalTableSettingsCmd)
}
