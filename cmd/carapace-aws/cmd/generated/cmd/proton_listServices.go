package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listServicesCmd = &cobra.Command{
	Use:   "list-services",
	Short: "List services with summaries of detail data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listServicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listServicesCmd).Standalone()

		proton_listServicesCmd.Flags().String("max-results", "", "The maximum number of services to list.")
		proton_listServicesCmd.Flags().String("next-token", "", "A token that indicates the location of the next service in the array of services, after the list of services that was previously requested.")
	})
	protonCmd.AddCommand(proton_listServicesCmd)
}
