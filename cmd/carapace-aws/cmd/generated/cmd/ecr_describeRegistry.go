package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_describeRegistryCmd = &cobra.Command{
	Use:   "describe-registry",
	Short: "Describes the settings for a registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_describeRegistryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_describeRegistryCmd).Standalone()

	})
	ecrCmd.AddCommand(ecr_describeRegistryCmd)
}
