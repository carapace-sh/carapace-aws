package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_deleteServiceCmd = &cobra.Command{
	Use:   "delete-service",
	Short: "Delete a service, with its instances and pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_deleteServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_deleteServiceCmd).Standalone()

		proton_deleteServiceCmd.Flags().String("name", "", "The name of the service to delete.")
		proton_deleteServiceCmd.MarkFlagRequired("name")
	})
	protonCmd.AddCommand(proton_deleteServiceCmd)
}
