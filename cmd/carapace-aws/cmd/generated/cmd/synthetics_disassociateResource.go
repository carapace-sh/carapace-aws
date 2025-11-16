package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_disassociateResourceCmd = &cobra.Command{
	Use:   "disassociate-resource",
	Short: "Removes a canary from a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_disassociateResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(synthetics_disassociateResourceCmd).Standalone()

		synthetics_disassociateResourceCmd.Flags().String("group-identifier", "", "Specifies the group.")
		synthetics_disassociateResourceCmd.Flags().String("resource-arn", "", "The ARN of the canary that you want to remove from the specified group.")
		synthetics_disassociateResourceCmd.MarkFlagRequired("group-identifier")
		synthetics_disassociateResourceCmd.MarkFlagRequired("resource-arn")
	})
	syntheticsCmd.AddCommand(synthetics_disassociateResourceCmd)
}
