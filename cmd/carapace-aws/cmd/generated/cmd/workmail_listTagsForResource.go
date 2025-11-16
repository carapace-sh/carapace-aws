package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags applied to an WorkMail organization resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listTagsForResourceCmd).Standalone()

	workmail_listTagsForResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
	workmail_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	workmailCmd.AddCommand(workmail_listTagsForResourceCmd)
}
