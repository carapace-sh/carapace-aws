package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getGroupVersionCmd = &cobra.Command{
	Use:   "get-group-version",
	Short: "Retrieves information about a group version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getGroupVersionCmd).Standalone()

	greengrass_getGroupVersionCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
	greengrass_getGroupVersionCmd.Flags().String("group-version-id", "", "The ID of the group version.")
	greengrass_getGroupVersionCmd.MarkFlagRequired("group-id")
	greengrass_getGroupVersionCmd.MarkFlagRequired("group-version-id")
	greengrassCmd.AddCommand(greengrass_getGroupVersionCmd)
}
