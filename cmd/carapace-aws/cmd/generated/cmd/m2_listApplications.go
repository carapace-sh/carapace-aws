package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists the applications associated with a specific Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_listApplicationsCmd).Standalone()

	m2_listApplicationsCmd.Flags().String("environment-id", "", "The unique identifier of the runtime environment where the applications are deployed.")
	m2_listApplicationsCmd.Flags().String("max-results", "", "The maximum number of applications to return.")
	m2_listApplicationsCmd.Flags().String("names", "", "The names of the applications.")
	m2_listApplicationsCmd.Flags().String("next-token", "", "A pagination token to control the number of applications displayed in the list.")
	m2Cmd.AddCommand(m2_listApplicationsCmd)
}
