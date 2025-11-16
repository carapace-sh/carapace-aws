package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspaces_modifyAccountCmd = &cobra.Command{
	Use:   "modify-account",
	Short: "Modifies the configuration of Bring Your Own License (BYOL) for the specified account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaces_modifyAccountCmd).Standalone()

	workspaces_modifyAccountCmd.Flags().String("dedicated-tenancy-management-cidr-range", "", "The IP address range, specified as an IPv4 CIDR block, for the management network interface.")
	workspaces_modifyAccountCmd.Flags().String("dedicated-tenancy-support", "", "The status of BYOL.")
	workspacesCmd.AddCommand(workspaces_modifyAccountCmd)
}
