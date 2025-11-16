package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates tags for a resource in CodeCommit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_tagResourceCmd).Standalone()

		codecommit_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which you want to add or update tags.")
		codecommit_tagResourceCmd.Flags().String("tags", "", "The key-value pair to use when tagging this repository.")
		codecommit_tagResourceCmd.MarkFlagRequired("resource-arn")
		codecommit_tagResourceCmd.MarkFlagRequired("tags")
	})
	codecommitCmd.AddCommand(codecommit_tagResourceCmd)
}
