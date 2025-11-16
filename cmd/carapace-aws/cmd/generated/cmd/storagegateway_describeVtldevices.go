package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeVtldevicesCmd = &cobra.Command{
	Use:   "describe-vtldevices",
	Short: "Returns a description of virtual tape library (VTL) devices for the specified tape gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeVtldevicesCmd).Standalone()

	storagegateway_describeVtldevicesCmd.Flags().String("gateway-arn", "", "")
	storagegateway_describeVtldevicesCmd.Flags().String("limit", "", "Specifies that the number of VTL devices described be limited to the specified number.")
	storagegateway_describeVtldevicesCmd.Flags().String("marker", "", "An opaque string that indicates the position at which to begin describing the VTL devices.")
	storagegateway_describeVtldevicesCmd.Flags().String("vtldevice-arns", "", "An array of strings, where each string represents the Amazon Resource Name (ARN) of a VTL device.")
	storagegateway_describeVtldevicesCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_describeVtldevicesCmd)
}
