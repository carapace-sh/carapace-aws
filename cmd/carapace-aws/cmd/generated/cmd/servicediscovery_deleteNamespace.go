package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_deleteNamespaceCmd = &cobra.Command{
	Use:   "delete-namespace",
	Short: "Deletes a namespace from the current account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_deleteNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_deleteNamespaceCmd).Standalone()

		servicediscovery_deleteNamespaceCmd.Flags().String("id", "", "The ID or Amazon Resource Name (ARN) of the namespace that you want to delete.")
		servicediscovery_deleteNamespaceCmd.MarkFlagRequired("id")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_deleteNamespaceCmd)
}
