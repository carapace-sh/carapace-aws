package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates one or more tags that are associated with an Amazon Security Lake resource: a subscriber, or the data lake configuration for your Amazon Web Services account in a particular Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securitylake_tagResourceCmd).Standalone()

		securitylake_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon Security Lake resource to add or update the tags for.")
		securitylake_tagResourceCmd.Flags().String("tags", "", "An array of objects, one for each tag (key and value) to associate with the Amazon Security Lake resource.")
		securitylake_tagResourceCmd.MarkFlagRequired("resource-arn")
		securitylake_tagResourceCmd.MarkFlagRequired("tags")
	})
	securitylakeCmd.AddCommand(securitylake_tagResourceCmd)
}
