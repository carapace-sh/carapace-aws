package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_getRegistryScanningConfigurationCmd = &cobra.Command{
	Use:   "get-registry-scanning-configuration",
	Short: "Retrieves the scanning configuration for a registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_getRegistryScanningConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_getRegistryScanningConfigurationCmd).Standalone()

	})
	ecrCmd.AddCommand(ecr_getRegistryScanningConfigurationCmd)
}
