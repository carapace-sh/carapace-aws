package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putResourceConfigCmd = &cobra.Command{
	Use:   "put-resource-config",
	Short: "Records the configuration state for the resource provided in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putResourceConfigCmd).Standalone()

	config_putResourceConfigCmd.Flags().String("configuration", "", "The configuration object of the resource in valid JSON format.")
	config_putResourceConfigCmd.Flags().String("resource-id", "", "Unique identifier of the resource.")
	config_putResourceConfigCmd.Flags().String("resource-name", "", "Name of the resource.")
	config_putResourceConfigCmd.Flags().String("resource-type", "", "The type of the resource.")
	config_putResourceConfigCmd.Flags().String("schema-version-id", "", "Version of the schema registered for the ResourceType in CloudFormation.")
	config_putResourceConfigCmd.Flags().String("tags", "", "Tags associated with the resource.")
	config_putResourceConfigCmd.MarkFlagRequired("configuration")
	config_putResourceConfigCmd.MarkFlagRequired("resource-id")
	config_putResourceConfigCmd.MarkFlagRequired("resource-type")
	config_putResourceConfigCmd.MarkFlagRequired("schema-version-id")
	configCmd.AddCommand(config_putResourceConfigCmd)
}
