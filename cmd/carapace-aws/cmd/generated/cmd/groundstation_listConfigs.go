package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_listConfigsCmd = &cobra.Command{
	Use:   "list-configs",
	Short: "Returns a list of `Config` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_listConfigsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_listConfigsCmd).Standalone()

		groundstation_listConfigsCmd.Flags().String("max-results", "", "Maximum number of `Configs` returned.")
		groundstation_listConfigsCmd.Flags().String("next-token", "", "Next token returned in the request of a previous `ListConfigs` call.")
	})
	groundstationCmd.AddCommand(groundstation_listConfigsCmd)
}
