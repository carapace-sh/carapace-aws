package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_batchGetResourceConfigCmd = &cobra.Command{
	Use:   "batch-get-resource-config",
	Short: "Returns the `BaseConfigurationItem` for one or more requested resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_batchGetResourceConfigCmd).Standalone()

	config_batchGetResourceConfigCmd.Flags().String("resource-keys", "", "A list of resource keys to be processed with the current request.")
	config_batchGetResourceConfigCmd.MarkFlagRequired("resource-keys")
	configCmd.AddCommand(config_batchGetResourceConfigCmd)
}
