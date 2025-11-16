package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_exportConfigurationsCmd = &cobra.Command{
	Use:   "export-configurations",
	Short: "Deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_exportConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_exportConfigurationsCmd).Standalone()

	})
	discoveryCmd.AddCommand(discovery_exportConfigurationsCmd)
}
