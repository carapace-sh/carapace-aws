package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listInstanceProfileTagsCmd = &cobra.Command{
	Use:   "list-instance-profile-tags",
	Short: "Lists the tags that are attached to the specified IAM instance profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listInstanceProfileTagsCmd).Standalone()

	iam_listInstanceProfileTagsCmd.Flags().String("instance-profile-name", "", "The name of the IAM instance profile whose tags you want to see.")
	iam_listInstanceProfileTagsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_listInstanceProfileTagsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_listInstanceProfileTagsCmd.MarkFlagRequired("instance-profile-name")
	iamCmd.AddCommand(iam_listInstanceProfileTagsCmd)
}
