package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Applies the specified tags to the specified WorkMailorganization resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_tagResourceCmd).Standalone()

		workmail_tagResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
		workmail_tagResourceCmd.Flags().String("tags", "", "The tag key-value pairs.")
		workmail_tagResourceCmd.MarkFlagRequired("resource-arn")
		workmail_tagResourceCmd.MarkFlagRequired("tags")
	})
	workmailCmd.AddCommand(workmail_tagResourceCmd)
}
