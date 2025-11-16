package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getIndexingConfigurationCmd = &cobra.Command{
	Use:   "get-indexing-configuration",
	Short: "Gets the indexing configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getIndexingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_getIndexingConfigurationCmd).Standalone()

	})
	iotCmd.AddCommand(iot_getIndexingConfigurationCmd)
}
