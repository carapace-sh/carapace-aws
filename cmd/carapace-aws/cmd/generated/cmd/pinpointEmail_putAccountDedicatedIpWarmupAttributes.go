package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_putAccountDedicatedIpWarmupAttributesCmd = &cobra.Command{
	Use:   "put-account-dedicated-ip-warmup-attributes",
	Short: "Enable or disable the automatic warm-up feature for dedicated IP addresses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_putAccountDedicatedIpWarmupAttributesCmd).Standalone()

	pinpointEmail_putAccountDedicatedIpWarmupAttributesCmd.Flags().String("auto-warmup-enabled", "", "Enables or disables the automatic warm-up feature for dedicated IP addresses that are associated with your Amazon Pinpoint account in the current AWS Region.")
	pinpointEmailCmd.AddCommand(pinpointEmail_putAccountDedicatedIpWarmupAttributesCmd)
}
