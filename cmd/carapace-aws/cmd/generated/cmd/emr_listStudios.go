package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listStudiosCmd = &cobra.Command{
	Use:   "list-studios",
	Short: "Returns a list of all Amazon EMR Studios associated with the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listStudiosCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_listStudiosCmd).Standalone()

		emr_listStudiosCmd.Flags().String("marker", "", "The pagination token that indicates the set of results to retrieve.")
	})
	emrCmd.AddCommand(emr_listStudiosCmd)
}
