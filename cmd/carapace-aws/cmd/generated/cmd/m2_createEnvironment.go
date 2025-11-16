package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_createEnvironmentCmd = &cobra.Command{
	Use:   "create-environment",
	Short: "Creates a runtime environment for a given runtime engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_createEnvironmentCmd).Standalone()

	m2_createEnvironmentCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request to create an environment.")
	m2_createEnvironmentCmd.Flags().String("description", "", "The description of the runtime environment.")
	m2_createEnvironmentCmd.Flags().String("engine-type", "", "The engine type for the runtime environment.")
	m2_createEnvironmentCmd.Flags().String("engine-version", "", "The version of the engine type for the runtime environment.")
	m2_createEnvironmentCmd.Flags().String("high-availability-config", "", "The details of a high availability configuration for this runtime environment.")
	m2_createEnvironmentCmd.Flags().String("instance-type", "", "The type of instance for the runtime environment.")
	m2_createEnvironmentCmd.Flags().String("kms-key-id", "", "The identifier of a customer managed key.")
	m2_createEnvironmentCmd.Flags().String("name", "", "The name of the runtime environment.")
	m2_createEnvironmentCmd.Flags().String("network-type", "", "The network type required for the runtime environment.")
	m2_createEnvironmentCmd.Flags().Bool("no-publicly-accessible", false, "Specifies whether the runtime environment is publicly accessible.")
	m2_createEnvironmentCmd.Flags().String("preferred-maintenance-window", "", "Configures the maintenance window that you want for the runtime environment.")
	m2_createEnvironmentCmd.Flags().Bool("publicly-accessible", false, "Specifies whether the runtime environment is publicly accessible.")
	m2_createEnvironmentCmd.Flags().String("security-group-ids", "", "The list of security groups for the VPC associated with this runtime environment.")
	m2_createEnvironmentCmd.Flags().String("storage-configurations", "", "Optional.")
	m2_createEnvironmentCmd.Flags().String("subnet-ids", "", "The list of subnets associated with the VPC for this runtime environment.")
	m2_createEnvironmentCmd.Flags().String("tags", "", "The tags for the runtime environment.")
	m2_createEnvironmentCmd.MarkFlagRequired("engine-type")
	m2_createEnvironmentCmd.MarkFlagRequired("instance-type")
	m2_createEnvironmentCmd.MarkFlagRequired("name")
	m2_createEnvironmentCmd.Flag("no-publicly-accessible").Hidden = true
	m2Cmd.AddCommand(m2_createEnvironmentCmd)
}
