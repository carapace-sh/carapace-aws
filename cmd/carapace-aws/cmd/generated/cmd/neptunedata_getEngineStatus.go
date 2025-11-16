package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getEngineStatusCmd = &cobra.Command{
	Use:   "get-engine-status",
	Short: "Retrieves the status of the graph database on the host.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getEngineStatusCmd).Standalone()

	neptunedataCmd.AddCommand(neptunedata_getEngineStatusCmd)
}
