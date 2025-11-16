package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_listEnvironmentsCmd = &cobra.Command{
	Use:   "list-environments",
	Short: "Lists the Amazon EVS environments in your Amazon Web Services account in the specified Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_listEnvironmentsCmd).Standalone()

	evs_listEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	evs_listEnvironmentsCmd.Flags().String("next-token", "", "A unique pagination token for each page.")
	evs_listEnvironmentsCmd.Flags().String("state", "", "The state of an environment.")
	evsCmd.AddCommand(evs_listEnvironmentsCmd)
}
