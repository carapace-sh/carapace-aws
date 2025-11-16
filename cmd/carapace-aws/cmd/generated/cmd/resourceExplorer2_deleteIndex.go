package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_deleteIndexCmd = &cobra.Command{
	Use:   "delete-index",
	Short: "Deletes the specified index and turns off Amazon Web Services Resource Explorer in the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_deleteIndexCmd).Standalone()

	resourceExplorer2_deleteIndexCmd.Flags().String("arn", "", "The [Amazon resource name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the index that you want to delete.")
	resourceExplorer2_deleteIndexCmd.MarkFlagRequired("arn")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_deleteIndexCmd)
}
