package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untags a resource in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_untagResourceCmd).Standalone()

	datazone_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to be untagged in Amazon DataZone.")
	datazone_untagResourceCmd.Flags().String("tag-keys", "", "Specifies the tag keys for the `UntagResource` action.")
	datazone_untagResourceCmd.MarkFlagRequired("resource-arn")
	datazone_untagResourceCmd.MarkFlagRequired("tag-keys")
	datazoneCmd.AddCommand(datazone_untagResourceCmd)
}
