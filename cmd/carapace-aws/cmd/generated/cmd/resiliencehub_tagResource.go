package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies one or more tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_tagResourceCmd).Standalone()

	resiliencehub_tagResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) of the resource.")
	resiliencehub_tagResourceCmd.Flags().String("tags", "", "The tags to assign to the resource.")
	resiliencehub_tagResourceCmd.MarkFlagRequired("resource-arn")
	resiliencehub_tagResourceCmd.MarkFlagRequired("tags")
	resiliencehubCmd.AddCommand(resiliencehub_tagResourceCmd)
}
