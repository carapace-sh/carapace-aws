package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_putStorageConfigurationCmd = &cobra.Command{
	Use:   "put-storage-configuration",
	Short: "Configures storage settings for IoT SiteWise.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_putStorageConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_putStorageConfigurationCmd).Standalone()

		iotsitewise_putStorageConfigurationCmd.Flags().String("disallow-ingest-null-na-n", "", "Describes the configuration for ingesting NULL and NaN data.")
		iotsitewise_putStorageConfigurationCmd.Flags().String("disassociated-data-storage", "", "Contains the storage configuration for time series (data streams) that aren't associated with asset properties.")
		iotsitewise_putStorageConfigurationCmd.Flags().String("multi-layer-storage", "", "Identifies a storage destination.")
		iotsitewise_putStorageConfigurationCmd.Flags().String("retention-period", "", "")
		iotsitewise_putStorageConfigurationCmd.Flags().String("storage-type", "", "The storage tier that you specified for your data.")
		iotsitewise_putStorageConfigurationCmd.Flags().String("warm-tier", "", "A service managed storage tier optimized for analytical queries.")
		iotsitewise_putStorageConfigurationCmd.Flags().String("warm-tier-retention-period", "", "Set this period to specify how long your data is stored in the warm tier before it is deleted.")
		iotsitewise_putStorageConfigurationCmd.MarkFlagRequired("storage-type")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_putStorageConfigurationCmd)
}
