package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_putProtocolsListCmd = &cobra.Command{
	Use:   "put-protocols-list",
	Short: "Creates an Firewall Manager protocols list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_putProtocolsListCmd).Standalone()

	fms_putProtocolsListCmd.Flags().String("protocols-list", "", "The details of the Firewall Manager protocols list to be created.")
	fms_putProtocolsListCmd.Flags().String("tag-list", "", "The tags associated with the resource.")
	fms_putProtocolsListCmd.MarkFlagRequired("protocols-list")
	fmsCmd.AddCommand(fms_putProtocolsListCmd)
}
