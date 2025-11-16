package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_getNamespaceCmd = &cobra.Command{
	Use:   "get-namespace",
	Short: "Gets information about a namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_getNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_getNamespaceCmd).Standalone()

		servicediscovery_getNamespaceCmd.Flags().String("id", "", "The ID or Amazon Resource Name (ARN) of the namespace that you want to get information about.")
		servicediscovery_getNamespaceCmd.MarkFlagRequired("id")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_getNamespaceCmd)
}
