package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_putBlockPublicAccessConfigurationCmd = &cobra.Command{
	Use:   "put-block-public-access-configuration",
	Short: "Creates or updates an Amazon EMR block public access configuration for your Amazon Web Services account in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_putBlockPublicAccessConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_putBlockPublicAccessConfigurationCmd).Standalone()

		emr_putBlockPublicAccessConfigurationCmd.Flags().String("block-public-access-configuration", "", "A configuration for Amazon EMR block public access.")
		emr_putBlockPublicAccessConfigurationCmd.MarkFlagRequired("block-public-access-configuration")
	})
	emrCmd.AddCommand(emr_putBlockPublicAccessConfigurationCmd)
}
