package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_createOrUpdateTagsCmd = &cobra.Command{
	Use:   "create-or-update-tags",
	Short: "Creates or updates tags for the specified Auto Scaling group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_createOrUpdateTagsCmd).Standalone()

	autoscaling_createOrUpdateTagsCmd.Flags().String("tags", "", "One or more tags.")
	autoscaling_createOrUpdateTagsCmd.MarkFlagRequired("tags")
	autoscalingCmd.AddCommand(autoscaling_createOrUpdateTagsCmd)
}
