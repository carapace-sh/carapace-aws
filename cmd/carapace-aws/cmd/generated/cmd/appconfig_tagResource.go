package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns metadata to an AppConfig resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_tagResourceCmd).Standalone()

	appconfig_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource for which to retrieve tags.")
	appconfig_tagResourceCmd.Flags().String("tags", "", "The key-value string map.")
	appconfig_tagResourceCmd.MarkFlagRequired("resource-arn")
	appconfig_tagResourceCmd.MarkFlagRequired("tags")
	appconfigCmd.AddCommand(appconfig_tagResourceCmd)
}
