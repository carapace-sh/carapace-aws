package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createPackageCmd = &cobra.Command{
	Use:   "create-package",
	Short: "Creates an IoT software package that can be deployed to your fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createPackageCmd).Standalone()

	iot_createPackageCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iot_createPackageCmd.Flags().String("description", "", "A summary of the package being created.")
	iot_createPackageCmd.Flags().String("package-name", "", "The name of the new software package.")
	iot_createPackageCmd.Flags().String("tags", "", "Metadata that can be used to manage the package.")
	iot_createPackageCmd.MarkFlagRequired("package-name")
	iotCmd.AddCommand(iot_createPackageCmd)
}
