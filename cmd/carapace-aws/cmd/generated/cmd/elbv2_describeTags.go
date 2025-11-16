package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeTagsCmd = &cobra.Command{
	Use:   "describe-tags",
	Short: "Describes the tags for the specified Elastic Load Balancing resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_describeTagsCmd).Standalone()

		elbv2_describeTagsCmd.Flags().String("resource-arns", "", "The Amazon Resource Names (ARN) of the resources.")
		elbv2_describeTagsCmd.MarkFlagRequired("resource-arns")
	})
	elbv2Cmd.AddCommand(elbv2_describeTagsCmd)
}
