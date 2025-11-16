package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_addTagsCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Adds the specified tags to the specified Elastic Load Balancing resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_addTagsCmd).Standalone()

	elbv2_addTagsCmd.Flags().String("resource-arns", "", "The Amazon Resource Name (ARN) of the resource.")
	elbv2_addTagsCmd.Flags().String("tags", "", "The tags.")
	elbv2_addTagsCmd.MarkFlagRequired("resource-arns")
	elbv2_addTagsCmd.MarkFlagRequired("tags")
	elbv2Cmd.AddCommand(elbv2_addTagsCmd)
}
