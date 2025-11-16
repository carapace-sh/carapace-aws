package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmIncidents_untagResourceCmd).Standalone()

		ssmIncidents_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the response plan you're removing a tag from.")
		ssmIncidents_untagResourceCmd.Flags().String("tag-keys", "", "The name of the tag to remove from the response plan.")
		ssmIncidents_untagResourceCmd.MarkFlagRequired("resource-arn")
		ssmIncidents_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	ssmIncidentsCmd.AddCommand(ssmIncidents_untagResourceCmd)
}
