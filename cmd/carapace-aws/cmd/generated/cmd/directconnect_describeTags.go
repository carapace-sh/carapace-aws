package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeTagsCmd = &cobra.Command{
	Use:   "describe-tags",
	Short: "Describes the tags associated with the specified Direct Connect resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeTagsCmd).Standalone()

	directconnect_describeTagsCmd.Flags().String("resource-arns", "", "The Amazon Resource Names (ARNs) of the resources.")
	directconnect_describeTagsCmd.MarkFlagRequired("resource-arns")
	directconnectCmd.AddCommand(directconnect_describeTagsCmd)
}
