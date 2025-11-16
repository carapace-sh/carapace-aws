package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_listKxEnvironmentsCmd = &cobra.Command{
	Use:   "list-kx-environments",
	Short: "Returns a list of kdb environments created in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_listKxEnvironmentsCmd).Standalone()

	finspace_listKxEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	finspace_listKxEnvironmentsCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspaceCmd.AddCommand(finspace_listKxEnvironmentsCmd)
}
