package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_getManagedViewCmd = &cobra.Command{
	Use:   "get-managed-view",
	Short: "Retrieves details of the specified [Amazon Web Services-managed view](https://docs.aws.amazon.com/resource-explorer/latest/userguide/aws-managed-views.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_getManagedViewCmd).Standalone()

	resourceExplorer2_getManagedViewCmd.Flags().String("managed-view-arn", "", "The Amazon resource name (ARN) of the managed view.")
	resourceExplorer2_getManagedViewCmd.MarkFlagRequired("managed-view-arn")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_getManagedViewCmd)
}
