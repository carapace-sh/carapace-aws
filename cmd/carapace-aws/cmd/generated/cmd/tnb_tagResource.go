package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags an AWS TNB resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_tagResourceCmd).Standalone()

		tnb_tagResourceCmd.Flags().String("resource-arn", "", "Resource ARN.")
		tnb_tagResourceCmd.Flags().String("tags", "", "A tag is a label that you assign to an Amazon Web Services resource.")
		tnb_tagResourceCmd.MarkFlagRequired("resource-arn")
		tnb_tagResourceCmd.MarkFlagRequired("tags")
	})
	tnbCmd.AddCommand(tnb_tagResourceCmd)
}
