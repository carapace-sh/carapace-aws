package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeConnectionLoaCmd = &cobra.Command{
	Use:   "describe-connection-loa",
	Short: "Deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeConnectionLoaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_describeConnectionLoaCmd).Standalone()

		directconnect_describeConnectionLoaCmd.Flags().String("connection-id", "", "The ID of the connection.")
		directconnect_describeConnectionLoaCmd.Flags().String("loa-content-type", "", "The standard media type for the LOA-CFA document.")
		directconnect_describeConnectionLoaCmd.Flags().String("provider-name", "", "The name of the APN partner or service provider who establishes connectivity on your behalf.")
		directconnect_describeConnectionLoaCmd.MarkFlagRequired("connection-id")
	})
	directconnectCmd.AddCommand(directconnect_describeConnectionLoaCmd)
}
