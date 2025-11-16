package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_listMembershipsCmd = &cobra.Command{
	Use:   "list-memberships",
	Short: "Returns the memberships that the calling principal can access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_listMembershipsCmd).Standalone()

	securityIr_listMembershipsCmd.Flags().String("max-results", "", "Request element for ListMemberships to limit the number of responses.")
	securityIr_listMembershipsCmd.Flags().String("next-token", "", "An optional string that, if supplied, must be copied from the output of a previous call to ListMemberships.")
	securityIrCmd.AddCommand(securityIr_listMembershipsCmd)
}
