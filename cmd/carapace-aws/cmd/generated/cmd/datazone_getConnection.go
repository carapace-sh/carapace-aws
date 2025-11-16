package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getConnectionCmd = &cobra.Command{
	Use:   "get-connection",
	Short: "Gets a connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getConnectionCmd).Standalone()

		datazone_getConnectionCmd.Flags().String("domain-identifier", "", "The ID of the domain where we get the connection.")
		datazone_getConnectionCmd.Flags().String("identifier", "", "The connection ID.")
		datazone_getConnectionCmd.Flags().Bool("no-with-secret", false, "Specifies whether a connection has a secret.")
		datazone_getConnectionCmd.Flags().Bool("with-secret", false, "Specifies whether a connection has a secret.")
		datazone_getConnectionCmd.MarkFlagRequired("domain-identifier")
		datazone_getConnectionCmd.MarkFlagRequired("identifier")
		datazone_getConnectionCmd.Flag("no-with-secret").Hidden = true
	})
	datazoneCmd.AddCommand(datazone_getConnectionCmd)
}
