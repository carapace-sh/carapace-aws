package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "The `TagResource` operation associates tags with an Amazon Managed Service for Prometheus resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_tagResourceCmd).Standalone()

		amp_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to apply tags to.")
		amp_tagResourceCmd.Flags().String("tags", "", "The list of tag keys and values to associate with the resource.")
		amp_tagResourceCmd.MarkFlagRequired("resource-arn")
		amp_tagResourceCmd.MarkFlagRequired("tags")
	})
	ampCmd.AddCommand(amp_tagResourceCmd)
}
