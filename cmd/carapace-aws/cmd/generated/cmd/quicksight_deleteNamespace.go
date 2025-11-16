package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteNamespaceCmd = &cobra.Command{
	Use:   "delete-namespace",
	Short: "Deletes a namespace and the users and groups that are associated with the namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteNamespaceCmd).Standalone()

		quicksight_deleteNamespaceCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that you want to delete the Quick Sight namespace from.")
		quicksight_deleteNamespaceCmd.Flags().String("namespace", "", "The namespace that you want to delete.")
		quicksight_deleteNamespaceCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteNamespaceCmd.MarkFlagRequired("namespace")
	})
	quicksightCmd.AddCommand(quicksight_deleteNamespaceCmd)
}
