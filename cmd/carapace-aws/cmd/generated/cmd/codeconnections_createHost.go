package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_createHostCmd = &cobra.Command{
	Use:   "create-host",
	Short: "Creates a resource that represents the infrastructure where a third-party provider is installed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_createHostCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeconnections_createHostCmd).Standalone()

		codeconnections_createHostCmd.Flags().String("name", "", "The name of the host to be created.")
		codeconnections_createHostCmd.Flags().String("provider-endpoint", "", "The endpoint of the infrastructure to be represented by the host after it is created.")
		codeconnections_createHostCmd.Flags().String("provider-type", "", "The name of the installed provider to be associated with your connection.")
		codeconnections_createHostCmd.Flags().String("tags", "", "Tags for the host to be created.")
		codeconnections_createHostCmd.Flags().String("vpc-configuration", "", "The VPC configuration to be provisioned for the host.")
		codeconnections_createHostCmd.MarkFlagRequired("name")
		codeconnections_createHostCmd.MarkFlagRequired("provider-endpoint")
		codeconnections_createHostCmd.MarkFlagRequired("provider-type")
	})
	codeconnectionsCmd.AddCommand(codeconnections_createHostCmd)
}
