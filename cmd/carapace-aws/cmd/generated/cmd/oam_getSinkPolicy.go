package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oam_getSinkPolicyCmd = &cobra.Command{
	Use:   "get-sink-policy",
	Short: "Returns the current sink policy attached to this sink.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oam_getSinkPolicyCmd).Standalone()

	oam_getSinkPolicyCmd.Flags().String("sink-identifier", "", "The ARN of the sink to retrieve the policy of.")
	oam_getSinkPolicyCmd.MarkFlagRequired("sink-identifier")
	oamCmd.AddCommand(oam_getSinkPolicyCmd)
}
