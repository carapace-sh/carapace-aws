package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags attached to the specified telemetry rule resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_listTagsForResourceCmd).Standalone()

	observabilityadmin_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the telemetry rule resource whose tags you want to list.")
	observabilityadmin_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	observabilityadminCmd.AddCommand(observabilityadmin_listTagsForResourceCmd)
}
