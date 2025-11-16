package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_deleteDataLakeCmd = &cobra.Command{
	Use:   "delete-data-lake",
	Short: "When you disable Amazon Security Lake from your account, Security Lake is disabled in all Amazon Web Services Regions and it stops collecting data from your sources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_deleteDataLakeCmd).Standalone()

	securitylake_deleteDataLakeCmd.Flags().String("regions", "", "The list of Regions where Security Lake is enabled.")
	securitylake_deleteDataLakeCmd.MarkFlagRequired("regions")
	securitylakeCmd.AddCommand(securitylake_deleteDataLakeCmd)
}
