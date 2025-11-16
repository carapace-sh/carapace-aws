package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getPackageConfigurationCmd = &cobra.Command{
	Use:   "get-package-configuration",
	Short: "Gets information about the specified software package's configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getPackageConfigurationCmd).Standalone()

	iotCmd.AddCommand(iot_getPackageConfigurationCmd)
}
