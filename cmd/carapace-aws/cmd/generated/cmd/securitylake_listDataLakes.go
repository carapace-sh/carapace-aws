package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_listDataLakesCmd = &cobra.Command{
	Use:   "list-data-lakes",
	Short: "Retrieves the Amazon Security Lake configuration object for the specified Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_listDataLakesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securitylake_listDataLakesCmd).Standalone()

		securitylake_listDataLakesCmd.Flags().String("regions", "", "The list of Regions where Security Lake is enabled.")
	})
	securitylakeCmd.AddCommand(securitylake_listDataLakesCmd)
}
