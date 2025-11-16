package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyAquaConfigurationCmd = &cobra.Command{
	Use:   "modify-aqua-configuration",
	Short: "This operation is retired.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyAquaConfigurationCmd).Standalone()

	redshift_modifyAquaConfigurationCmd.Flags().String("aqua-configuration-status", "", "This parameter is retired.")
	redshift_modifyAquaConfigurationCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster to be modified.")
	redshift_modifyAquaConfigurationCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_modifyAquaConfigurationCmd)
}
