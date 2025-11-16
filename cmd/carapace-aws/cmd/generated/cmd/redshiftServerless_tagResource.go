package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_tagResourceCmd).Standalone()

	redshiftServerless_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to tag.")
	redshiftServerless_tagResourceCmd.Flags().String("tags", "", "The map of the key-value pairs used to tag the resource.")
	redshiftServerless_tagResourceCmd.MarkFlagRequired("resource-arn")
	redshiftServerless_tagResourceCmd.MarkFlagRequired("tags")
	redshiftServerlessCmd.AddCommand(redshiftServerless_tagResourceCmd)
}
