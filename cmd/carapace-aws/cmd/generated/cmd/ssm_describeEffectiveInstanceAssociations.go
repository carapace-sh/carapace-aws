package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeEffectiveInstanceAssociationsCmd = &cobra.Command{
	Use:   "describe-effective-instance-associations",
	Short: "All associations for the managed nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeEffectiveInstanceAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeEffectiveInstanceAssociationsCmd).Standalone()

		ssm_describeEffectiveInstanceAssociationsCmd.Flags().String("instance-id", "", "The managed node ID for which you want to view all associations.")
		ssm_describeEffectiveInstanceAssociationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_describeEffectiveInstanceAssociationsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_describeEffectiveInstanceAssociationsCmd.MarkFlagRequired("instance-id")
	})
	ssmCmd.AddCommand(ssm_describeEffectiveInstanceAssociationsCmd)
}
