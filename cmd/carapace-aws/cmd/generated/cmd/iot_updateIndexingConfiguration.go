package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateIndexingConfigurationCmd = &cobra.Command{
	Use:   "update-indexing-configuration",
	Short: "Updates the search configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateIndexingConfigurationCmd).Standalone()

	iot_updateIndexingConfigurationCmd.Flags().String("thing-group-indexing-configuration", "", "Thing group indexing configuration.")
	iot_updateIndexingConfigurationCmd.Flags().String("thing-indexing-configuration", "", "Thing indexing configuration.")
	iotCmd.AddCommand(iot_updateIndexingConfigurationCmd)
}
