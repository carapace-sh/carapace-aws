package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var frauddetector_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(frauddetector_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(frauddetector_tagResourceCmd).Standalone()

		frauddetector_tagResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
		frauddetector_tagResourceCmd.Flags().String("tags", "", "The tags to assign to the resource.")
		frauddetector_tagResourceCmd.MarkFlagRequired("resource-arn")
		frauddetector_tagResourceCmd.MarkFlagRequired("tags")
	})
	frauddetectorCmd.AddCommand(frauddetector_tagResourceCmd)
}
