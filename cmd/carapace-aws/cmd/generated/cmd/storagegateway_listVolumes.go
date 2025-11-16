package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listVolumesCmd = &cobra.Command{
	Use:   "list-volumes",
	Short: "Lists the iSCSI stored volumes of a gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listVolumesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_listVolumesCmd).Standalone()

		storagegateway_listVolumesCmd.Flags().String("gateway-arn", "", "")
		storagegateway_listVolumesCmd.Flags().String("limit", "", "Specifies that the list of volumes returned be limited to the specified number of items.")
		storagegateway_listVolumesCmd.Flags().String("marker", "", "A string that indicates the position at which to begin the returned list of volumes.")
	})
	storagegatewayCmd.AddCommand(storagegateway_listVolumesCmd)
}
