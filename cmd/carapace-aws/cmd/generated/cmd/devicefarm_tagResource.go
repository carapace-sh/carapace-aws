package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the specified tags to a resource with the specified `resourceArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_tagResourceCmd).Standalone()

	devicefarm_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource or resources to which to add tags.")
	devicefarm_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
	devicefarm_tagResourceCmd.MarkFlagRequired("resource-arn")
	devicefarm_tagResourceCmd.MarkFlagRequired("tags")
	devicefarmCmd.AddCommand(devicefarm_tagResourceCmd)
}
