package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_updateSmbsecurityStrategyCmd = &cobra.Command{
	Use:   "update-smbsecurity-strategy",
	Short: "Updates the SMB security strategy level for an Amazon S3 file gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_updateSmbsecurityStrategyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_updateSmbsecurityStrategyCmd).Standalone()

		storagegateway_updateSmbsecurityStrategyCmd.Flags().String("gateway-arn", "", "")
		storagegateway_updateSmbsecurityStrategyCmd.Flags().String("smbsecurity-strategy", "", "Specifies the type of security strategy.")
		storagegateway_updateSmbsecurityStrategyCmd.MarkFlagRequired("gateway-arn")
		storagegateway_updateSmbsecurityStrategyCmd.MarkFlagRequired("smbsecurity-strategy")
	})
	storagegatewayCmd.AddCommand(storagegateway_updateSmbsecurityStrategyCmd)
}
