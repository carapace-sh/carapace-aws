package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createSoftwareUpdateJobCmd = &cobra.Command{
	Use:   "create-software-update-job",
	Short: "Creates a software update for a core or group of cores (specified as an IoT thing group.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createSoftwareUpdateJobCmd).Standalone()

	greengrass_createSoftwareUpdateJobCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
	greengrass_createSoftwareUpdateJobCmd.Flags().String("s3-url-signer-role", "", "")
	greengrass_createSoftwareUpdateJobCmd.Flags().String("software-to-update", "", "")
	greengrass_createSoftwareUpdateJobCmd.Flags().String("update-agent-log-level", "", "")
	greengrass_createSoftwareUpdateJobCmd.Flags().String("update-targets", "", "")
	greengrass_createSoftwareUpdateJobCmd.Flags().String("update-targets-architecture", "", "")
	greengrass_createSoftwareUpdateJobCmd.Flags().String("update-targets-operating-system", "", "")
	greengrass_createSoftwareUpdateJobCmd.MarkFlagRequired("s3-url-signer-role")
	greengrass_createSoftwareUpdateJobCmd.MarkFlagRequired("software-to-update")
	greengrass_createSoftwareUpdateJobCmd.MarkFlagRequired("update-targets")
	greengrass_createSoftwareUpdateJobCmd.MarkFlagRequired("update-targets-architecture")
	greengrass_createSoftwareUpdateJobCmd.MarkFlagRequired("update-targets-operating-system")
	greengrassCmd.AddCommand(greengrass_createSoftwareUpdateJobCmd)
}
