package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_listLifecycleExecutionsCmd = &cobra.Command{
	Use:   "list-lifecycle-executions",
	Short: "Get the lifecycle runtime history for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_listLifecycleExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_listLifecycleExecutionsCmd).Standalone()

		imagebuilder_listLifecycleExecutionsCmd.Flags().String("max-results", "", "Specify the maximum number of items to return in a request.")
		imagebuilder_listLifecycleExecutionsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
		imagebuilder_listLifecycleExecutionsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which to get a list of lifecycle runtime instances.")
		imagebuilder_listLifecycleExecutionsCmd.MarkFlagRequired("resource-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_listLifecycleExecutionsCmd)
}
