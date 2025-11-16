package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listEnvironmentsCmd = &cobra.Command{
	Use:   "list-environments",
	Short: "List environments with detail data summaries.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listEnvironmentsCmd).Standalone()

	proton_listEnvironmentsCmd.Flags().String("environment-templates", "", "An array of the versions of the environment template.")
	proton_listEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of environments to list.")
	proton_listEnvironmentsCmd.Flags().String("next-token", "", "A token that indicates the location of the next environment in the array of environments, after the list of environments that was previously requested.")
	protonCmd.AddCommand(proton_listEnvironmentsCmd)
}
