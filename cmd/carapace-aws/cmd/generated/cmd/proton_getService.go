package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getServiceCmd = &cobra.Command{
	Use:   "get-service",
	Short: "Get detailed data for a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_getServiceCmd).Standalone()

		proton_getServiceCmd.Flags().String("name", "", "The name of the service that you want to get the detailed data for.")
		proton_getServiceCmd.MarkFlagRequired("name")
	})
	protonCmd.AddCommand(proton_getServiceCmd)
}
