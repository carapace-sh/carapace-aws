package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_createTableCmd = &cobra.Command{
	Use:   "create-table",
	Short: "The `CreateTable` operation adds a new table to your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_createTableCmd).Standalone()

	dynamodb_createTableCmd.Flags().String("attribute-definitions", "", "An array of attributes that describe the key schema for the table and indexes.")
	dynamodb_createTableCmd.Flags().String("billing-mode", "", "Controls how you are charged for read and write throughput and how you manage capacity.")
	dynamodb_createTableCmd.Flags().String("deletion-protection-enabled", "", "Indicates whether deletion protection is to be enabled (true) or disabled (false) on the table.")
	dynamodb_createTableCmd.Flags().String("global-secondary-indexes", "", "One or more global secondary indexes (the maximum is 20) to be created on the table.")
	dynamodb_createTableCmd.Flags().String("key-schema", "", "Specifies the attributes that make up the primary key for a table or an index.")
	dynamodb_createTableCmd.Flags().String("local-secondary-indexes", "", "One or more local secondary indexes (the maximum is 5) to be created on the table.")
	dynamodb_createTableCmd.Flags().String("on-demand-throughput", "", "Sets the maximum number of read and write units for the specified table in on-demand capacity mode.")
	dynamodb_createTableCmd.Flags().String("provisioned-throughput", "", "Represents the provisioned throughput settings for a specified table or index.")
	dynamodb_createTableCmd.Flags().String("resource-policy", "", "An Amazon Web Services resource-based policy document in JSON format that will be attached to the table.")
	dynamodb_createTableCmd.Flags().String("ssespecification", "", "Represents the settings used to enable server-side encryption.")
	dynamodb_createTableCmd.Flags().String("stream-specification", "", "The settings for DynamoDB Streams on the table.")
	dynamodb_createTableCmd.Flags().String("table-class", "", "The table class of the new table.")
	dynamodb_createTableCmd.Flags().String("table-name", "", "The name of the table to create.")
	dynamodb_createTableCmd.Flags().String("tags", "", "A list of key-value pairs to label the table.")
	dynamodb_createTableCmd.Flags().String("warm-throughput", "", "Represents the warm throughput (in read units per second and write units per second) for creating a table.")
	dynamodb_createTableCmd.MarkFlagRequired("attribute-definitions")
	dynamodb_createTableCmd.MarkFlagRequired("key-schema")
	dynamodb_createTableCmd.MarkFlagRequired("table-name")
	dynamodbCmd.AddCommand(dynamodb_createTableCmd)
}
