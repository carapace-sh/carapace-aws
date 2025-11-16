package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listTestGridProjectsCmd = &cobra.Command{
	Use:   "list-test-grid-projects",
	Short: "Gets a list of all Selenium testing projects in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listTestGridProjectsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_listTestGridProjectsCmd).Standalone()

		devicefarm_listTestGridProjectsCmd.Flags().String("max-result", "", "Return no more than this number of results.")
		devicefarm_listTestGridProjectsCmd.Flags().String("next-token", "", "From a response, used to continue a paginated listing.")
	})
	devicefarmCmd.AddCommand(devicefarm_listTestGridProjectsCmd)
}
