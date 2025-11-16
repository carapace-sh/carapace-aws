package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_createBgppeerCmd = &cobra.Command{
	Use:   "create-bgppeer",
	Short: "Creates a BGP peer on the specified virtual interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_createBgppeerCmd).Standalone()

	directconnect_createBgppeerCmd.Flags().String("new-bgppeer", "", "Information about the BGP peer.")
	directconnect_createBgppeerCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface.")
	directconnectCmd.AddCommand(directconnect_createBgppeerCmd)
}
