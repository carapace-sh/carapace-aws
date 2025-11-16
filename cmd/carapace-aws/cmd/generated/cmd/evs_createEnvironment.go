package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_createEnvironmentCmd = &cobra.Command{
	Use:   "create-environment",
	Short: "Creates an Amazon EVS environment that runs VCF software, such as SDDC Manager, NSX Manager, and vCenter Server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_createEnvironmentCmd).Standalone()

	evs_createEnvironmentCmd.Flags().String("client-token", "", "This parameter is not used in Amazon EVS currently.")
	evs_createEnvironmentCmd.Flags().String("connectivity-info", "", "The connectivity configuration for the environment.")
	evs_createEnvironmentCmd.Flags().String("environment-name", "", "The name to give to your environment.")
	evs_createEnvironmentCmd.Flags().String("hosts", "", "The ESXi hosts to add to the environment.")
	evs_createEnvironmentCmd.Flags().String("initial-vlans", "", "The initial VLAN subnets for the Amazon EVS environment.")
	evs_createEnvironmentCmd.Flags().String("kms-key-id", "", "A unique ID for the customer-managed KMS key that is used to encrypt the VCF credential pairs for SDDC Manager, NSX Manager, and vCenter appliances.")
	evs_createEnvironmentCmd.Flags().String("license-info", "", "The license information that Amazon EVS requires to create an environment.")
	evs_createEnvironmentCmd.Flags().Bool("no-terms-accepted", false, "Customer confirmation that the customer has purchased and will continue to maintain the required number of VCF software licenses to cover all physical processor cores in the Amazon EVS environment.")
	evs_createEnvironmentCmd.Flags().String("service-access-security-groups", "", "The security group that controls communication between the Amazon EVS control plane and VPC.")
	evs_createEnvironmentCmd.Flags().String("service-access-subnet-id", "", "The subnet that is used to establish connectivity between the Amazon EVS control plane and VPC.")
	evs_createEnvironmentCmd.Flags().String("site-id", "", "The Broadcom Site ID that is allocated to you as part of your electronic software delivery.")
	evs_createEnvironmentCmd.Flags().String("tags", "", "Metadata that assists with categorization and organization.")
	evs_createEnvironmentCmd.Flags().Bool("terms-accepted", false, "Customer confirmation that the customer has purchased and will continue to maintain the required number of VCF software licenses to cover all physical processor cores in the Amazon EVS environment.")
	evs_createEnvironmentCmd.Flags().String("vcf-hostnames", "", "The DNS hostnames for the virtual machines that host the VCF management appliances.")
	evs_createEnvironmentCmd.Flags().String("vcf-version", "", "The VCF version to use for the environment.")
	evs_createEnvironmentCmd.Flags().String("vpc-id", "", "A unique ID for the VPC that the environment is deployed inside.")
	evs_createEnvironmentCmd.MarkFlagRequired("connectivity-info")
	evs_createEnvironmentCmd.MarkFlagRequired("hosts")
	evs_createEnvironmentCmd.MarkFlagRequired("initial-vlans")
	evs_createEnvironmentCmd.MarkFlagRequired("license-info")
	evs_createEnvironmentCmd.Flag("no-terms-accepted").Hidden = true
	evs_createEnvironmentCmd.MarkFlagRequired("no-terms-accepted")
	evs_createEnvironmentCmd.MarkFlagRequired("service-access-subnet-id")
	evs_createEnvironmentCmd.MarkFlagRequired("site-id")
	evs_createEnvironmentCmd.MarkFlagRequired("terms-accepted")
	evs_createEnvironmentCmd.MarkFlagRequired("vcf-hostnames")
	evs_createEnvironmentCmd.MarkFlagRequired("vcf-version")
	evs_createEnvironmentCmd.MarkFlagRequired("vpc-id")
	evsCmd.AddCommand(evs_createEnvironmentCmd)
}
