package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putDedicatedIpPoolScalingAttributesCmd = &cobra.Command{
	Use:   "put-dedicated-ip-pool-scaling-attributes",
	Short: "Used to convert a dedicated IP pool to a different scaling mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putDedicatedIpPoolScalingAttributesCmd).Standalone()

	sesv2_putDedicatedIpPoolScalingAttributesCmd.Flags().String("pool-name", "", "The name of the dedicated IP pool.")
	sesv2_putDedicatedIpPoolScalingAttributesCmd.Flags().String("scaling-mode", "", "The scaling mode to apply to the dedicated IP pool.")
	sesv2_putDedicatedIpPoolScalingAttributesCmd.MarkFlagRequired("pool-name")
	sesv2_putDedicatedIpPoolScalingAttributesCmd.MarkFlagRequired("scaling-mode")
	sesv2Cmd.AddCommand(sesv2_putDedicatedIpPoolScalingAttributesCmd)
}
