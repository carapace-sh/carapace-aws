package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_getDataLakeSourcesCmd = &cobra.Command{
	Use:   "get-data-lake-sources",
	Short: "Retrieves a snapshot of the current Region, including whether Amazon Security Lake is enabled for those accounts and which sources Security Lake is collecting data from.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_getDataLakeSourcesCmd).Standalone()

	securitylake_getDataLakeSourcesCmd.Flags().String("accounts", "", "The Amazon Web Services account ID for which a static snapshot of the current Amazon Web Services Region, including enabled accounts and log sources, is retrieved.")
	securitylake_getDataLakeSourcesCmd.Flags().String("max-results", "", "The maximum limit of accounts for which the static snapshot of the current Region, including enabled accounts and log sources, is retrieved.")
	securitylake_getDataLakeSourcesCmd.Flags().String("next-token", "", "Lists if there are more results available.")
	securitylakeCmd.AddCommand(securitylake_getDataLakeSourcesCmd)
}
