package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_putSinkPolicyCmd = &cobra.Command{
	Use:   "put-sink-policy",
	Short: "Creates or updates the resource policy that grants permissions to source accounts to link to the monitoring account sink.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_putSinkPolicyCmd).Standalone()

	oam_putSinkPolicyCmd.Flags().String("policy", "", "The JSON policy to use.")
	oam_putSinkPolicyCmd.Flags().String("sink-identifier", "", "The ARN of the sink to attach this policy to.")
	oam_putSinkPolicyCmd.MarkFlagRequired("policy")
	oam_putSinkPolicyCmd.MarkFlagRequired("sink-identifier")
	oamCmd.AddCommand(oam_putSinkPolicyCmd)
}
