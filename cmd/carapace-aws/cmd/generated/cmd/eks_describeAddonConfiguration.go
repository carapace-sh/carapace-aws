package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeAddonConfigurationCmd = &cobra.Command{
	Use:   "describe-addon-configuration",
	Short: "Returns configuration options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeAddonConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_describeAddonConfigurationCmd).Standalone()

		eks_describeAddonConfigurationCmd.Flags().String("addon-name", "", "The name of the add-on.")
		eks_describeAddonConfigurationCmd.Flags().String("addon-version", "", "The version of the add-on.")
		eks_describeAddonConfigurationCmd.MarkFlagRequired("addon-name")
		eks_describeAddonConfigurationCmd.MarkFlagRequired("addon-version")
	})
	eksCmd.AddCommand(eks_describeAddonConfigurationCmd)
}
