package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_updateWorkgroupCmd = &cobra.Command{
	Use:   "update-workgroup",
	Short: "Updates a workgroup with the specified configuration settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_updateWorkgroupCmd).Standalone()

	redshiftServerless_updateWorkgroupCmd.Flags().String("base-capacity", "", "The new base data warehouse capacity in Redshift Processing Units (RPUs).")
	redshiftServerless_updateWorkgroupCmd.Flags().String("config-parameters", "", "An array of parameters to set for advanced control over a database.")
	redshiftServerless_updateWorkgroupCmd.Flags().Bool("enhanced-vpc-routing", false, "The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.")
	redshiftServerless_updateWorkgroupCmd.Flags().String("ip-address-type", "", "The IP address type that the workgroup supports.")
	redshiftServerless_updateWorkgroupCmd.Flags().String("max-capacity", "", "The maximum data-warehouse capacity Amazon Redshift Serverless uses to serve queries.")
	redshiftServerless_updateWorkgroupCmd.Flags().Bool("no-enhanced-vpc-routing", false, "The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.")
	redshiftServerless_updateWorkgroupCmd.Flags().Bool("no-publicly-accessible", false, "A value that specifies whether the workgroup can be accessible from a public network.")
	redshiftServerless_updateWorkgroupCmd.Flags().String("port", "", "The custom port to use when connecting to a workgroup.")
	redshiftServerless_updateWorkgroupCmd.Flags().String("price-performance-target", "", "An object that represents the price performance target settings for the workgroup.")
	redshiftServerless_updateWorkgroupCmd.Flags().Bool("publicly-accessible", false, "A value that specifies whether the workgroup can be accessible from a public network.")
	redshiftServerless_updateWorkgroupCmd.Flags().String("security-group-ids", "", "An array of security group IDs to associate with the workgroup.")
	redshiftServerless_updateWorkgroupCmd.Flags().String("subnet-ids", "", "An array of VPC subnet IDs to associate with the workgroup.")
	redshiftServerless_updateWorkgroupCmd.Flags().String("track-name", "", "An optional parameter for the name of the track for the workgroup.")
	redshiftServerless_updateWorkgroupCmd.Flags().String("workgroup-name", "", "The name of the workgroup to update.")
	redshiftServerless_updateWorkgroupCmd.Flag("no-enhanced-vpc-routing").Hidden = true
	redshiftServerless_updateWorkgroupCmd.Flag("no-publicly-accessible").Hidden = true
	redshiftServerless_updateWorkgroupCmd.MarkFlagRequired("workgroup-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_updateWorkgroupCmd)
}
