package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies tags to an existing Amazon Web Services X-Ray group or sampling rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_tagResourceCmd).Standalone()

	xray_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Number (ARN) of an X-Ray group or sampling rule.")
	xray_tagResourceCmd.Flags().String("tags", "", "A map that contains one or more tag keys and tag values to attach to an X-Ray group or sampling rule.")
	xray_tagResourceCmd.MarkFlagRequired("resource-arn")
	xray_tagResourceCmd.MarkFlagRequired("tags")
	xrayCmd.AddCommand(xray_tagResourceCmd)
}
