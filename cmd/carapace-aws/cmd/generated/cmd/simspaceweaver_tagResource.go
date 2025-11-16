package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to a SimSpace Weaver resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_tagResourceCmd).Standalone()

	simspaceweaver_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to add tags to.")
	simspaceweaver_tagResourceCmd.Flags().String("tags", "", "A list of tags to apply to the resource.")
	simspaceweaver_tagResourceCmd.MarkFlagRequired("resource-arn")
	simspaceweaver_tagResourceCmd.MarkFlagRequired("tags")
	simspaceweaverCmd.AddCommand(simspaceweaver_tagResourceCmd)
}
