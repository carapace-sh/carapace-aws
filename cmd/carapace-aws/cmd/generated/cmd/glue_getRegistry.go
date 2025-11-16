package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getRegistryCmd = &cobra.Command{
	Use:   "get-registry",
	Short: "Describes the specified registry in detail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getRegistryCmd).Standalone()

	glue_getRegistryCmd.Flags().String("registry-id", "", "This is a wrapper structure that may contain the registry name and Amazon Resource Name (ARN).")
	glue_getRegistryCmd.MarkFlagRequired("registry-id")
	glueCmd.AddCommand(glue_getRegistryCmd)
}
