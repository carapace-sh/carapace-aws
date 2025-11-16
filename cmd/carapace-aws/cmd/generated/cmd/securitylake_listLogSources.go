package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_listLogSourcesCmd = &cobra.Command{
	Use:   "list-log-sources",
	Short: "Retrieves the log sources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_listLogSourcesCmd).Standalone()

	securitylake_listLogSourcesCmd.Flags().String("accounts", "", "The list of Amazon Web Services accounts for which log sources are displayed.")
	securitylake_listLogSourcesCmd.Flags().String("max-results", "", "The maximum number of accounts for which the log sources are displayed.")
	securitylake_listLogSourcesCmd.Flags().String("next-token", "", "If nextToken is returned, there are more results available.")
	securitylake_listLogSourcesCmd.Flags().String("regions", "", "The list of Regions for which log sources are displayed.")
	securitylake_listLogSourcesCmd.Flags().String("sources", "", "The list of sources for which log sources are displayed.")
	securitylakeCmd.AddCommand(securitylake_listLogSourcesCmd)
}
