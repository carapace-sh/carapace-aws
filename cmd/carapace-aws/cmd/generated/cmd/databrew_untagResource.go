package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes metadata tags from a DataBrew resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_untagResourceCmd).Standalone()

	databrew_untagResourceCmd.Flags().String("resource-arn", "", "A DataBrew resource from which you want to remove a tag or tags.")
	databrew_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys (names) of one or more tags to be removed.")
	databrew_untagResourceCmd.MarkFlagRequired("resource-arn")
	databrew_untagResourceCmd.MarkFlagRequired("tag-keys")
	databrewCmd.AddCommand(databrew_untagResourceCmd)
}
