package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteRegistryCmd = &cobra.Command{
	Use:   "delete-registry",
	Short: "Delete the entire registry including schema and all of its versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteRegistryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_deleteRegistryCmd).Standalone()

		glue_deleteRegistryCmd.Flags().String("registry-id", "", "This is a wrapper structure that may contain the registry name and Amazon Resource Name (ARN).")
		glue_deleteRegistryCmd.MarkFlagRequired("registry-id")
	})
	glueCmd.AddCommand(glue_deleteRegistryCmd)
}
