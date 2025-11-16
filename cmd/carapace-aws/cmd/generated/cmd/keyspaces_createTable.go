package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_createTableCmd = &cobra.Command{
	Use:   "create-table",
	Short: "The `CreateTable` operation adds a new table to the specified keyspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_createTableCmd).Standalone()

	keyspaces_createTableCmd.Flags().String("auto-scaling-specification", "", "The optional auto scaling settings for a table in provisioned capacity mode.")
	keyspaces_createTableCmd.Flags().String("capacity-specification", "", "Specifies the read/write throughput capacity mode for the table.")
	keyspaces_createTableCmd.Flags().String("cdc-specification", "", "The CDC stream settings of the table.")
	keyspaces_createTableCmd.Flags().String("client-side-timestamps", "", "Enables client-side timestamps for the table.")
	keyspaces_createTableCmd.Flags().String("comment", "", "This parameter allows to enter a description of the table.")
	keyspaces_createTableCmd.Flags().String("default-time-to-live", "", "The default Time to Live setting in seconds for the table.")
	keyspaces_createTableCmd.Flags().String("encryption-specification", "", "Specifies how the encryption key for encryption at rest is managed for the table.")
	keyspaces_createTableCmd.Flags().String("keyspace-name", "", "The name of the keyspace that the table is going to be created in.")
	keyspaces_createTableCmd.Flags().String("point-in-time-recovery", "", "Specifies if `pointInTimeRecovery` is enabled or disabled for the table.")
	keyspaces_createTableCmd.Flags().String("replica-specifications", "", "The optional Amazon Web Services Region specific settings of a multi-Region table.")
	keyspaces_createTableCmd.Flags().String("schema-definition", "", "The `schemaDefinition` consists of the following parameters.")
	keyspaces_createTableCmd.Flags().String("table-name", "", "The name of the table.")
	keyspaces_createTableCmd.Flags().String("tags", "", "A list of key-value pair tags to be attached to the resource.")
	keyspaces_createTableCmd.Flags().String("ttl", "", "Enables Time to Live custom settings for the table.")
	keyspaces_createTableCmd.MarkFlagRequired("keyspace-name")
	keyspaces_createTableCmd.MarkFlagRequired("schema-definition")
	keyspaces_createTableCmd.MarkFlagRequired("table-name")
	keyspacesCmd.AddCommand(keyspaces_createTableCmd)
}
