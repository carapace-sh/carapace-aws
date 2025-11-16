package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listEnvironmentTemplatesCmd = &cobra.Command{
	Use:   "list-environment-templates",
	Short: "List environment templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listEnvironmentTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listEnvironmentTemplatesCmd).Standalone()

		proton_listEnvironmentTemplatesCmd.Flags().String("max-results", "", "The maximum number of environment templates to list.")
		proton_listEnvironmentTemplatesCmd.Flags().String("next-token", "", "A token that indicates the location of the next environment template in the array of environment templates, after the list of environment templates that was previously requested.")
	})
	protonCmd.AddCommand(proton_listEnvironmentTemplatesCmd)
}
