package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_getViewCmd = &cobra.Command{
	Use:   "get-view",
	Short: "Retrieves details of the specified view.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_getViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_getViewCmd).Standalone()

		resourceExplorer2_getViewCmd.Flags().String("view-arn", "", "The [Amazon resource name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the view that you want information about.")
		resourceExplorer2_getViewCmd.MarkFlagRequired("view-arn")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_getViewCmd)
}
