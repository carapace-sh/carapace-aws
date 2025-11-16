package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_batchGetViewCmd = &cobra.Command{
	Use:   "batch-get-view",
	Short: "Retrieves details about a list of views.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_batchGetViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_batchGetViewCmd).Standalone()

		resourceExplorer2_batchGetViewCmd.Flags().String("view-arns", "", "A list of [Amazon resource names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) that identify the views you want details for.")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_batchGetViewCmd)
}
