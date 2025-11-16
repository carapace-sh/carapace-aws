package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listLocalDisksCmd = &cobra.Command{
	Use:   "list-local-disks",
	Short: "Returns a list of the gateway's local disks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listLocalDisksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_listLocalDisksCmd).Standalone()

		storagegateway_listLocalDisksCmd.Flags().String("gateway-arn", "", "")
		storagegateway_listLocalDisksCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_listLocalDisksCmd)
}
