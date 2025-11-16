package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_deleteViewCmd = &cobra.Command{
	Use:   "delete-view",
	Short: "Deletes the specified view.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_deleteViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_deleteViewCmd).Standalone()

		resourceExplorer2_deleteViewCmd.Flags().String("view-arn", "", "The [Amazon resource name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the view that you want to delete.")
		resourceExplorer2_deleteViewCmd.MarkFlagRequired("view-arn")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_deleteViewCmd)
}
