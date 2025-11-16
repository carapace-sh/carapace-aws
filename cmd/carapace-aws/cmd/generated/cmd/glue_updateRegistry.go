package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateRegistryCmd = &cobra.Command{
	Use:   "update-registry",
	Short: "Updates an existing registry which is used to hold a collection of schemas.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateRegistryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_updateRegistryCmd).Standalone()

		glue_updateRegistryCmd.Flags().String("description", "", "A description of the registry.")
		glue_updateRegistryCmd.Flags().String("registry-id", "", "This is a wrapper structure that may contain the registry name and Amazon Resource Name (ARN).")
		glue_updateRegistryCmd.MarkFlagRequired("description")
		glue_updateRegistryCmd.MarkFlagRequired("registry-id")
	})
	glueCmd.AddCommand(glue_updateRegistryCmd)
}
