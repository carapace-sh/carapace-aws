package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createRegistryCmd = &cobra.Command{
	Use:   "create-registry",
	Short: "Creates a new registry which may be used to hold a collection of schemas.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createRegistryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createRegistryCmd).Standalone()

		glue_createRegistryCmd.Flags().String("description", "", "A description of the registry.")
		glue_createRegistryCmd.Flags().String("registry-name", "", "Name of the registry to be created of max length of 255, and may only contain letters, numbers, hyphen, underscore, dollar sign, or hash mark.")
		glue_createRegistryCmd.Flags().String("tags", "", "Amazon Web Services tags that contain a key value pair and may be searched by console, command line, or API.")
		glue_createRegistryCmd.MarkFlagRequired("registry-name")
	})
	glueCmd.AddCommand(glue_createRegistryCmd)
}
