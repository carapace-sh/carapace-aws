package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to a resource with the specified `resourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_tagResourceCmd).Standalone()

	ecrPublic_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to add tags to.")
	ecrPublic_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
	ecrPublic_tagResourceCmd.MarkFlagRequired("resource-arn")
	ecrPublic_tagResourceCmd.MarkFlagRequired("tags")
	ecrPublicCmd.AddCommand(ecrPublic_tagResourceCmd)
}
