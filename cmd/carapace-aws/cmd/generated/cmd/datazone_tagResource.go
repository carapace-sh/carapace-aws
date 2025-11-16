package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a resource in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_tagResourceCmd).Standalone()

	datazone_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to be tagged in Amazon DataZone.")
	datazone_tagResourceCmd.Flags().String("tags", "", "Specifies the tags for the `TagResource` action.")
	datazone_tagResourceCmd.MarkFlagRequired("resource-arn")
	datazone_tagResourceCmd.MarkFlagRequired("tags")
	datazoneCmd.AddCommand(datazone_tagResourceCmd)
}
