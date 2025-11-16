package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeTagsCmd = &cobra.Command{
	Use:   "describe-tags",
	Short: "Describes the specified tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeTagsCmd).Standalone()

		autoscaling_describeTagsCmd.Flags().String("filters", "", "One or more filters to scope the tags to return.")
		autoscaling_describeTagsCmd.Flags().String("max-records", "", "The maximum number of items to return with this call.")
		autoscaling_describeTagsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	autoscalingCmd.AddCommand(autoscaling_describeTagsCmd)
}
