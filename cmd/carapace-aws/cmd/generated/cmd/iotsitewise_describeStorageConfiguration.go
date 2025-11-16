package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeStorageConfigurationCmd = &cobra.Command{
	Use:   "describe-storage-configuration",
	Short: "Retrieves information about the storage configuration for IoT SiteWise.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeStorageConfigurationCmd).Standalone()

	iotsitewiseCmd.AddCommand(iotsitewise_describeStorageConfigurationCmd)
}
