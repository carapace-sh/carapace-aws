package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_putRegistryScanningConfigurationCmd = &cobra.Command{
	Use:   "put-registry-scanning-configuration",
	Short: "Creates or updates the scanning configuration for your private registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_putRegistryScanningConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_putRegistryScanningConfigurationCmd).Standalone()

		ecr_putRegistryScanningConfigurationCmd.Flags().String("rules", "", "The scanning rules to use for the registry.")
		ecr_putRegistryScanningConfigurationCmd.Flags().String("scan-type", "", "The scanning type to set for the registry.")
	})
	ecrCmd.AddCommand(ecr_putRegistryScanningConfigurationCmd)
}
