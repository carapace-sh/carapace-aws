package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_listApplicationVersionsCmd = &cobra.Command{
	Use:   "list-application-versions",
	Short: "Returns a list of the application versions for a specific application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_listApplicationVersionsCmd).Standalone()

	m2_listApplicationVersionsCmd.Flags().String("application-id", "", "The unique identifier of the application.")
	m2_listApplicationVersionsCmd.Flags().String("max-results", "", "The maximum number of application versions to return.")
	m2_listApplicationVersionsCmd.Flags().String("next-token", "", "A pagination token returned from a previous call to this operation.")
	m2_listApplicationVersionsCmd.MarkFlagRequired("application-id")
	m2Cmd.AddCommand(m2_listApplicationVersionsCmd)
}
