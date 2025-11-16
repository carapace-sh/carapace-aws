package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getConfigurationCmd = &cobra.Command{
	Use:   "get-configuration",
	Short: "Retrieves setting configurations for Inspector scans.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_getConfigurationCmd).Standalone()

	})
	inspector2Cmd.AddCommand(inspector2_getConfigurationCmd)
}
