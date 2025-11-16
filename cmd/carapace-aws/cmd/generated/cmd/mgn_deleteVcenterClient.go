package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_deleteVcenterClientCmd = &cobra.Command{
	Use:   "delete-vcenter-client",
	Short: "Deletes a given vCenter client by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_deleteVcenterClientCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_deleteVcenterClientCmd).Standalone()

		mgn_deleteVcenterClientCmd.Flags().String("vcenter-client-id", "", "ID of resource to be deleted.")
		mgn_deleteVcenterClientCmd.MarkFlagRequired("vcenter-client-id")
	})
	mgnCmd.AddCommand(mgn_deleteVcenterClientCmd)
}
