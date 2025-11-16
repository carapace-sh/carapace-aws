package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_listListenersCmd = &cobra.Command{
	Use:   "list-listeners",
	Short: "List the listeners for an accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_listListenersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_listListenersCmd).Standalone()

		globalaccelerator_listListenersCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the accelerator for which you want to list listener objects.")
		globalaccelerator_listListenersCmd.Flags().String("max-results", "", "The number of listener objects that you want to return with this call.")
		globalaccelerator_listListenersCmd.Flags().String("next-token", "", "The token for the next set of results.")
		globalaccelerator_listListenersCmd.MarkFlagRequired("accelerator-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_listListenersCmd)
}
