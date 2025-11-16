package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds to or modifies the tags of the given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_tagResourceCmd).Standalone()

	iot_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource.")
	iot_tagResourceCmd.Flags().String("tags", "", "The new or modified tags for the resource.")
	iot_tagResourceCmd.MarkFlagRequired("resource-arn")
	iot_tagResourceCmd.MarkFlagRequired("tags")
	iotCmd.AddCommand(iot_tagResourceCmd)
}
