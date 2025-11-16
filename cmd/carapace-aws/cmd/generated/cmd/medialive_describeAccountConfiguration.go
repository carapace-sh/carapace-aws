package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeAccountConfigurationCmd = &cobra.Command{
	Use:   "describe-account-configuration",
	Short: "Describe account configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeAccountConfigurationCmd).Standalone()

	medialiveCmd.AddCommand(medialive_describeAccountConfigurationCmd)
}
