package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_listDataLakeExceptionsCmd = &cobra.Command{
	Use:   "list-data-lake-exceptions",
	Short: "Lists the Amazon Security Lake exceptions that you can use to find the source of problems and fix them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_listDataLakeExceptionsCmd).Standalone()

	securitylake_listDataLakeExceptionsCmd.Flags().String("max-results", "", "Lists the maximum number of failures in Security Lake.")
	securitylake_listDataLakeExceptionsCmd.Flags().String("next-token", "", "Lists if there are more results available.")
	securitylake_listDataLakeExceptionsCmd.Flags().String("regions", "", "The Amazon Web Services Regions from which exceptions are retrieved.")
	securitylakeCmd.AddCommand(securitylake_listDataLakeExceptionsCmd)
}
