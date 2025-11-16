package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscoveryCmd = &cobra.Command{
	Use:   "servicediscovery",
	Short: "Cloud Map\n\nWith Cloud Map, you can configure public DNS, private DNS, or HTTP namespaces that your microservice applications run in.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscoveryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscoveryCmd).Standalone()

	})
	rootCmd.AddCommand(servicediscoveryCmd)
}
