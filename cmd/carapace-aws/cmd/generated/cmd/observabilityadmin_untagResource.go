package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from a telemetry rule resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_untagResourceCmd).Standalone()

	observabilityadmin_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the telemetry rule resource to remove tags from.")
	observabilityadmin_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the telemetry rule resource.")
	observabilityadmin_untagResourceCmd.MarkFlagRequired("resource-arn")
	observabilityadmin_untagResourceCmd.MarkFlagRequired("tag-keys")
	observabilityadminCmd.AddCommand(observabilityadmin_untagResourceCmd)
}
