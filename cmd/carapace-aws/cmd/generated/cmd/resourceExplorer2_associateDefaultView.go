package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_associateDefaultViewCmd = &cobra.Command{
	Use:   "associate-default-view",
	Short: "Sets the specified view as the default for the Amazon Web Services Region in which you call this operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_associateDefaultViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_associateDefaultViewCmd).Standalone()

		resourceExplorer2_associateDefaultViewCmd.Flags().String("view-arn", "", "The [Amazon resource name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the view to set as the default for the Amazon Web Services Region and Amazon Web Services account in which you call this operation.")
		resourceExplorer2_associateDefaultViewCmd.MarkFlagRequired("view-arn")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_associateDefaultViewCmd)
}
