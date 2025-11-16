package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateServiceCmd = &cobra.Command{
	Use:   "update-service",
	Short: "Edit a service description or use a spec to add and delete service instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_updateServiceCmd).Standalone()

		proton_updateServiceCmd.Flags().String("description", "", "The edited service description.")
		proton_updateServiceCmd.Flags().String("name", "", "The name of the service to edit.")
		proton_updateServiceCmd.Flags().String("spec", "", "Lists the service instances to add and the existing service instances to remain.")
		proton_updateServiceCmd.MarkFlagRequired("name")
	})
	protonCmd.AddCommand(proton_updateServiceCmd)
}
