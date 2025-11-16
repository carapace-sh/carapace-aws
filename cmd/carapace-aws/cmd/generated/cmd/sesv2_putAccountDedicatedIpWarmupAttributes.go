package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putAccountDedicatedIpWarmupAttributesCmd = &cobra.Command{
	Use:   "put-account-dedicated-ip-warmup-attributes",
	Short: "Enable or disable the automatic warm-up feature for dedicated IP addresses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putAccountDedicatedIpWarmupAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putAccountDedicatedIpWarmupAttributesCmd).Standalone()

		sesv2_putAccountDedicatedIpWarmupAttributesCmd.Flags().String("auto-warmup-enabled", "", "Enables or disables the automatic warm-up feature for dedicated IP addresses that are associated with your Amazon SES account in the current Amazon Web Services Region.")
	})
	sesv2Cmd.AddCommand(sesv2_putAccountDedicatedIpWarmupAttributesCmd)
}
