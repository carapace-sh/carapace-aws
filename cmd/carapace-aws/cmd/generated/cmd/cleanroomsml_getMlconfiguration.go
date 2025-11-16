package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getMlconfigurationCmd = &cobra.Command{
	Use:   "get-mlconfiguration",
	Short: "Returns information about a specific ML configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getMlconfigurationCmd).Standalone()

	cleanroomsml_getMlconfigurationCmd.Flags().String("membership-identifier", "", "The membership ID of the member that owns the ML configuration you want to return information about.")
	cleanroomsml_getMlconfigurationCmd.MarkFlagRequired("membership-identifier")
	cleanroomsmlCmd.AddCommand(cleanroomsml_getMlconfigurationCmd)
}
