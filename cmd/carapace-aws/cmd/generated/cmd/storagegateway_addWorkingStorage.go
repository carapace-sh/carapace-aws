package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_addWorkingStorageCmd = &cobra.Command{
	Use:   "add-working-storage",
	Short: "Configures one or more gateway local disks as working storage for a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_addWorkingStorageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_addWorkingStorageCmd).Standalone()

		storagegateway_addWorkingStorageCmd.Flags().String("disk-ids", "", "An array of strings that identify disks that are to be configured as working storage.")
		storagegateway_addWorkingStorageCmd.Flags().String("gateway-arn", "", "")
		storagegateway_addWorkingStorageCmd.MarkFlagRequired("disk-ids")
		storagegateway_addWorkingStorageCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_addWorkingStorageCmd)
}
