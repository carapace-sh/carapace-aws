package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appmeshCmd = &cobra.Command{
	Use:   "appmesh",
	Short: "App Mesh is a service mesh based on the Envoy proxy that makes it easy to monitor and control microservices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appmeshCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appmeshCmd).Standalone()

	})
	rootCmd.AddCommand(appmeshCmd)
}
