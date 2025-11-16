package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putDedicatedIpWarmupAttributesCmd = &cobra.Command{
	Use:   "put-dedicated-ip-warmup-attributes",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putDedicatedIpWarmupAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putDedicatedIpWarmupAttributesCmd).Standalone()

		sesv2_putDedicatedIpWarmupAttributesCmd.Flags().String("ip", "", "The dedicated IP address that you want to update the warm-up attributes for.")
		sesv2_putDedicatedIpWarmupAttributesCmd.Flags().String("warmup-percentage", "", "The warm-up percentage that you want to associate with the dedicated IP address.")
		sesv2_putDedicatedIpWarmupAttributesCmd.MarkFlagRequired("ip")
		sesv2_putDedicatedIpWarmupAttributesCmd.MarkFlagRequired("warmup-percentage")
	})
	sesv2Cmd.AddCommand(sesv2_putDedicatedIpWarmupAttributesCmd)
}
