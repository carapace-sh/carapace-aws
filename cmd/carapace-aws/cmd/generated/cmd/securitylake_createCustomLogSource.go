package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_createCustomLogSourceCmd = &cobra.Command{
	Use:   "create-custom-log-source",
	Short: "Adds a third-party custom source in Amazon Security Lake, from the Amazon Web Services Region where you want to create a custom source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_createCustomLogSourceCmd).Standalone()

	securitylake_createCustomLogSourceCmd.Flags().String("configuration", "", "The configuration used for the third-party custom source.")
	securitylake_createCustomLogSourceCmd.Flags().String("event-classes", "", "The Open Cybersecurity Schema Framework (OCSF) event classes which describes the type of data that the custom source will send to Security Lake.")
	securitylake_createCustomLogSourceCmd.Flags().String("source-name", "", "Specify the name for a third-party custom source.")
	securitylake_createCustomLogSourceCmd.Flags().String("source-version", "", "Specify the source version for the third-party custom source, to limit log collection to a specific version of custom data source.")
	securitylake_createCustomLogSourceCmd.MarkFlagRequired("configuration")
	securitylake_createCustomLogSourceCmd.MarkFlagRequired("source-name")
	securitylakeCmd.AddCommand(securitylake_createCustomLogSourceCmd)
}
