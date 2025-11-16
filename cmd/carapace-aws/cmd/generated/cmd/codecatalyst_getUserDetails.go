package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_getUserDetailsCmd = &cobra.Command{
	Use:   "get-user-details",
	Short: "Returns information about a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_getUserDetailsCmd).Standalone()

	codecatalyst_getUserDetailsCmd.Flags().String("id", "", "The system-generated unique ID of the user.")
	codecatalyst_getUserDetailsCmd.Flags().String("user-name", "", "The name of the user as displayed in Amazon CodeCatalyst.")
	codecatalystCmd.AddCommand(codecatalyst_getUserDetailsCmd)
}
