package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_getBlockPublicAccessConfigurationCmd = &cobra.Command{
	Use:   "get-block-public-access-configuration",
	Short: "Returns the Amazon EMR block public access configuration for your Amazon Web Services account in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_getBlockPublicAccessConfigurationCmd).Standalone()

	emrCmd.AddCommand(emr_getBlockPublicAccessConfigurationCmd)
}
