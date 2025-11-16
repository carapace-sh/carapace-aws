package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeGroupCmd = &cobra.Command{
	Use:   "describe-group",
	Short: "Returns an Amazon Quick Sight group's description and Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeGroupCmd).Standalone()

		quicksight_describeGroupCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the group is in.")
		quicksight_describeGroupCmd.Flags().String("group-name", "", "The name of the group that you want to describe.")
		quicksight_describeGroupCmd.Flags().String("namespace", "", "The namespace of the group that you want described.")
		quicksight_describeGroupCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeGroupCmd.MarkFlagRequired("group-name")
		quicksight_describeGroupCmd.MarkFlagRequired("namespace")
	})
	quicksightCmd.AddCommand(quicksight_describeGroupCmd)
}
