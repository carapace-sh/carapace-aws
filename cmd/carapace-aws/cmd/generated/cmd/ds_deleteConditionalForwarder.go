package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_deleteConditionalForwarderCmd = &cobra.Command{
	Use:   "delete-conditional-forwarder",
	Short: "Deletes a conditional forwarder that has been set up for your Amazon Web Services directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_deleteConditionalForwarderCmd).Standalone()

	ds_deleteConditionalForwarderCmd.Flags().String("directory-id", "", "The directory ID for which you are deleting the conditional forwarder.")
	ds_deleteConditionalForwarderCmd.Flags().String("remote-domain-name", "", "The fully qualified domain name (FQDN) of the remote domain with which you are deleting the conditional forwarder.")
	ds_deleteConditionalForwarderCmd.MarkFlagRequired("directory-id")
	ds_deleteConditionalForwarderCmd.MarkFlagRequired("remote-domain-name")
	dsCmd.AddCommand(ds_deleteConditionalForwarderCmd)
}
