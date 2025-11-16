package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeTapeRecoveryPointsCmd = &cobra.Command{
	Use:   "describe-tape-recovery-points",
	Short: "Returns a list of virtual tape recovery points that are available for the specified tape gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeTapeRecoveryPointsCmd).Standalone()

	storagegateway_describeTapeRecoveryPointsCmd.Flags().String("gateway-arn", "", "")
	storagegateway_describeTapeRecoveryPointsCmd.Flags().String("limit", "", "Specifies that the number of virtual tape recovery points that are described be limited to the specified number.")
	storagegateway_describeTapeRecoveryPointsCmd.Flags().String("marker", "", "An opaque string that indicates the position at which to begin describing the virtual tape recovery points.")
	storagegateway_describeTapeRecoveryPointsCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_describeTapeRecoveryPointsCmd)
}
