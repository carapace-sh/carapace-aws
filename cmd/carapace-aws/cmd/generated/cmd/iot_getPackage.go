package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_getPackageCmd = &cobra.Command{
	Use:   "get-package",
	Short: "Gets information about the specified software package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_getPackageCmd).Standalone()

	iot_getPackageCmd.Flags().String("package-name", "", "The name of the target software package.")
	iot_getPackageCmd.MarkFlagRequired("package-name")
	iotCmd.AddCommand(iot_getPackageCmd)
}
