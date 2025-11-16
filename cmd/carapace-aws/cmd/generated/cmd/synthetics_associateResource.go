package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_associateResourceCmd = &cobra.Command{
	Use:   "associate-resource",
	Short: "Associates a canary with a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_associateResourceCmd).Standalone()

	synthetics_associateResourceCmd.Flags().String("group-identifier", "", "Specifies the group.")
	synthetics_associateResourceCmd.Flags().String("resource-arn", "", "The ARN of the canary that you want to associate with the specified group.")
	synthetics_associateResourceCmd.MarkFlagRequired("group-identifier")
	synthetics_associateResourceCmd.MarkFlagRequired("resource-arn")
	syntheticsCmd.AddCommand(synthetics_associateResourceCmd)
}
