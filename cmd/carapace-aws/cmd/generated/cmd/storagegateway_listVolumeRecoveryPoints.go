package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_listVolumeRecoveryPointsCmd = &cobra.Command{
	Use:   "list-volume-recovery-points",
	Short: "Lists the recovery points for a specified gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_listVolumeRecoveryPointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_listVolumeRecoveryPointsCmd).Standalone()

		storagegateway_listVolumeRecoveryPointsCmd.Flags().String("gateway-arn", "", "")
		storagegateway_listVolumeRecoveryPointsCmd.MarkFlagRequired("gateway-arn")
	})
	storagegatewayCmd.AddCommand(storagegateway_listVolumeRecoveryPointsCmd)
}
