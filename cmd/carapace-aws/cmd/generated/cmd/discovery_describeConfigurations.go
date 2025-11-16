package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_describeConfigurationsCmd = &cobra.Command{
	Use:   "describe-configurations",
	Short: "Retrieves attributes for a list of configuration item IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_describeConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_describeConfigurationsCmd).Standalone()

		discovery_describeConfigurationsCmd.Flags().String("configuration-ids", "", "One or more configuration IDs.")
		discovery_describeConfigurationsCmd.MarkFlagRequired("configuration-ids")
	})
	discoveryCmd.AddCommand(discovery_describeConfigurationsCmd)
}
