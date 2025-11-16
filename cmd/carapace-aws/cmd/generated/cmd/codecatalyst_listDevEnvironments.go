package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_listDevEnvironmentsCmd = &cobra.Command{
	Use:   "list-dev-environments",
	Short: "Retrieves a list of Dev Environments in a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_listDevEnvironmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_listDevEnvironmentsCmd).Standalone()

		codecatalyst_listDevEnvironmentsCmd.Flags().String("filters", "", "Information about filters to apply to narrow the results returned in the list.")
		codecatalyst_listDevEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of results to show in a single call to this API.")
		codecatalyst_listDevEnvironmentsCmd.Flags().String("next-token", "", "A token returned from a call to this API to indicate the next batch of results to return, if any.")
		codecatalyst_listDevEnvironmentsCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_listDevEnvironmentsCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_listDevEnvironmentsCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_listDevEnvironmentsCmd)
}
