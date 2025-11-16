package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "Gets a list of tags for the specified CloudHSM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_listTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsmv2_listTagsCmd).Standalone()

		cloudhsmv2_listTagsCmd.Flags().String("max-results", "", "The maximum number of tags to return in the response.")
		cloudhsmv2_listTagsCmd.Flags().String("next-token", "", "The `NextToken` value that you received in the previous response.")
		cloudhsmv2_listTagsCmd.Flags().String("resource-id", "", "The cluster identifier (ID) for the cluster whose tags you are getting.")
		cloudhsmv2_listTagsCmd.MarkFlagRequired("resource-id")
	})
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_listTagsCmd)
}
