package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeWorkingStorageCmd = &cobra.Command{
	Use:   "describe-working-storage",
	Short: "Returns information about the working storage of a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeWorkingStorageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeWorkingStorageCmd).Standalone()

		storagegateway_describeWorkingStorageCmd.Flags().String("gateway-arn", "", "")
		storagegateway_describeWorkingStorageCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeWorkingStorageCmd)
}
