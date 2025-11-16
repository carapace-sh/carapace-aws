package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_listFhirdatastoresCmd = &cobra.Command{
	Use:   "list-fhirdatastores",
	Short: "List all FHIR-enabled data stores in a user’s account, regardless of data store status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_listFhirdatastoresCmd).Standalone()

	healthlake_listFhirdatastoresCmd.Flags().String("filter", "", "List all filters associated with a FHIR data store request.")
	healthlake_listFhirdatastoresCmd.Flags().String("max-results", "", "The maximum number of data stores returned on a page.")
	healthlake_listFhirdatastoresCmd.Flags().String("next-token", "", "The token used to retrieve the next page of data stores when results are paginated.")
	healthlakeCmd.AddCommand(healthlake_listFhirdatastoresCmd)
}
