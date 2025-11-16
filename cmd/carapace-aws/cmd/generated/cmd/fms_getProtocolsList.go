package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_getProtocolsListCmd = &cobra.Command{
	Use:   "get-protocols-list",
	Short: "Returns information about the specified Firewall Manager protocols list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_getProtocolsListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_getProtocolsListCmd).Standalone()

		fms_getProtocolsListCmd.Flags().Bool("default-list", false, "Specifies whether the list to retrieve is a default list owned by Firewall Manager.")
		fms_getProtocolsListCmd.Flags().String("list-id", "", "The ID of the Firewall Manager protocols list that you want the details for.")
		fms_getProtocolsListCmd.Flags().Bool("no-default-list", false, "Specifies whether the list to retrieve is a default list owned by Firewall Manager.")
		fms_getProtocolsListCmd.MarkFlagRequired("list-id")
		fms_getProtocolsListCmd.Flag("no-default-list").Hidden = true
	})
	fmsCmd.AddCommand(fms_getProtocolsListCmd)
}
