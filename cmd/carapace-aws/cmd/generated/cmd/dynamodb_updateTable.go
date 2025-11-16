package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_updateTableCmd = &cobra.Command{
	Use:   "update-table",
	Short: "Modifies the provisioned throughput settings, global secondary indexes, or DynamoDB Streams settings for a given table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_updateTableCmd).Standalone()

	dynamodb_updateTableCmd.Flags().String("attribute-definitions", "", "An array of attributes that describe the key schema for the table and indexes.")
	dynamodb_updateTableCmd.Flags().String("billing-mode", "", "Controls how you are charged for read and write throughput and how you manage capacity.")
	dynamodb_updateTableCmd.Flags().String("deletion-protection-enabled", "", "Indicates whether deletion protection is to be enabled (true) or disabled (false) on the table.")
	dynamodb_updateTableCmd.Flags().String("global-secondary-index-updates", "", "An array of one or more global secondary indexes for the table.")
	dynamodb_updateTableCmd.Flags().String("global-table-witness-updates", "", "A list of witness updates for a MRSC global table.")
	dynamodb_updateTableCmd.Flags().String("multi-region-consistency", "", "Specifies the consistency mode for a new global table.")
	dynamodb_updateTableCmd.Flags().String("on-demand-throughput", "", "Updates the maximum number of read and write units for the specified table in on-demand capacity mode.")
	dynamodb_updateTableCmd.Flags().String("provisioned-throughput", "", "The new provisioned throughput settings for the specified table or index.")
	dynamodb_updateTableCmd.Flags().String("replica-updates", "", "A list of replica update actions (create, delete, or update) for the table.")
	dynamodb_updateTableCmd.Flags().String("ssespecification", "", "The new server-side encryption settings for the specified table.")
	dynamodb_updateTableCmd.Flags().String("stream-specification", "", "Represents the DynamoDB Streams configuration for the table.")
	dynamodb_updateTableCmd.Flags().String("table-class", "", "The table class of the table to be updated.")
	dynamodb_updateTableCmd.Flags().String("table-name", "", "The name of the table to be updated.")
	dynamodb_updateTableCmd.Flags().String("warm-throughput", "", "Represents the warm throughput (in read units per second and write units per second) for updating a table.")
	dynamodb_updateTableCmd.MarkFlagRequired("table-name")
	dynamodbCmd.AddCommand(dynamodb_updateTableCmd)
}
