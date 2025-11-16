package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to an Athena resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_tagResourceCmd).Standalone()

		athena_tagResourceCmd.Flags().String("resource-arn", "", "Specifies the ARN of the Athena resource to which tags are to be added.")
		athena_tagResourceCmd.Flags().String("tags", "", "A collection of one or more tags, separated by commas, to be added to an Athena resource.")
		athena_tagResourceCmd.MarkFlagRequired("resource-arn")
		athena_tagResourceCmd.MarkFlagRequired("tags")
	})
	athenaCmd.AddCommand(athena_tagResourceCmd)
}
