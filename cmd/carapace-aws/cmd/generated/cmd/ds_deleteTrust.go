package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_deleteTrustCmd = &cobra.Command{
	Use:   "delete-trust",
	Short: "Deletes an existing trust relationship between your Managed Microsoft AD directory and an external domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_deleteTrustCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_deleteTrustCmd).Standalone()

		ds_deleteTrustCmd.Flags().String("delete-associated-conditional-forwarder", "", "Delete a conditional forwarder as part of a DeleteTrustRequest.")
		ds_deleteTrustCmd.Flags().String("trust-id", "", "The Trust ID of the trust relationship to be deleted.")
		ds_deleteTrustCmd.MarkFlagRequired("trust-id")
	})
	dsCmd.AddCommand(ds_deleteTrustCmd)
}
