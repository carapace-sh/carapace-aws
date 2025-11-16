package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_putMlconfigurationCmd = &cobra.Command{
	Use:   "put-mlconfiguration",
	Short: "Assigns information about an ML configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_putMlconfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_putMlconfigurationCmd).Standalone()

		cleanroomsml_putMlconfigurationCmd.Flags().String("default-output-location", "", "The default Amazon S3 location where ML output is stored for the specified member.")
		cleanroomsml_putMlconfigurationCmd.Flags().String("membership-identifier", "", "The membership ID of the member that is being configured.")
		cleanroomsml_putMlconfigurationCmd.MarkFlagRequired("default-output-location")
		cleanroomsml_putMlconfigurationCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_putMlconfigurationCmd)
}
