package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_deleteCustomLogSourceCmd = &cobra.Command{
	Use:   "delete-custom-log-source",
	Short: "Removes a custom log source from Amazon Security Lake, to stop sending data from the custom source to Security Lake.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_deleteCustomLogSourceCmd).Standalone()

	securitylake_deleteCustomLogSourceCmd.Flags().String("source-name", "", "The source name of custom log source that you want to delete.")
	securitylake_deleteCustomLogSourceCmd.Flags().String("source-version", "", "The source version for the third-party custom source.")
	securitylake_deleteCustomLogSourceCmd.MarkFlagRequired("source-name")
	securitylakeCmd.AddCommand(securitylake_deleteCustomLogSourceCmd)
}
