package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_updateIndexTypeCmd = &cobra.Command{
	Use:   "update-index-type",
	Short: "Changes the type of the index from one of the following types to the other.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_updateIndexTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_updateIndexTypeCmd).Standalone()

		resourceExplorer2_updateIndexTypeCmd.Flags().String("arn", "", "The [Amazon resource name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the index that you want to update.")
		resourceExplorer2_updateIndexTypeCmd.Flags().String("type", "", "The type of the index.")
		resourceExplorer2_updateIndexTypeCmd.MarkFlagRequired("arn")
		resourceExplorer2_updateIndexTypeCmd.MarkFlagRequired("type")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_updateIndexTypeCmd)
}
