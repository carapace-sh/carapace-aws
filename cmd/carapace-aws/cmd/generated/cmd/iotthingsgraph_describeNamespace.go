package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_describeNamespaceCmd = &cobra.Command{
	Use:   "describe-namespace",
	Short: "Gets the latest version of the user's namespace and the public version that it is tracking.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_describeNamespaceCmd).Standalone()

	iotthingsgraph_describeNamespaceCmd.Flags().String("namespace-name", "", "The name of the user's namespace.")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_describeNamespaceCmd)
}
