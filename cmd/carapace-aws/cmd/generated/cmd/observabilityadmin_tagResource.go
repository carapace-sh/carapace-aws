package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadmin_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates tags for a telemetry rule resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadmin_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(observabilityadmin_tagResourceCmd).Standalone()

		observabilityadmin_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the telemetry rule resource to tag.")
		observabilityadmin_tagResourceCmd.Flags().String("tags", "", "The key-value pairs to add or update for the telemetry rule resource.")
		observabilityadmin_tagResourceCmd.MarkFlagRequired("resource-arn")
		observabilityadmin_tagResourceCmd.MarkFlagRequired("tags")
	})
	observabilityadminCmd.AddCommand(observabilityadmin_tagResourceCmd)
}
