package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_getMembershipCmd = &cobra.Command{
	Use:   "get-membership",
	Short: "Returns the attributes of a membership.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_getMembershipCmd).Standalone()

	securityIr_getMembershipCmd.Flags().String("membership-id", "", "Required element for GetMembership to identify the membership ID to query.")
	securityIr_getMembershipCmd.MarkFlagRequired("membership-id")
	securityIrCmd.AddCommand(securityIr_getMembershipCmd)
}
