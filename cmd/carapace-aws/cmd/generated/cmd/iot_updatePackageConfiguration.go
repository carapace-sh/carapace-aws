package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updatePackageConfigurationCmd = &cobra.Command{
	Use:   "update-package-configuration",
	Short: "Updates the software package configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updatePackageConfigurationCmd).Standalone()

	iot_updatePackageConfigurationCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iot_updatePackageConfigurationCmd.Flags().String("version-update-by-jobs-config", "", "Configuration to manage job's package version reporting.")
	iotCmd.AddCommand(iot_updatePackageConfigurationCmd)
}
