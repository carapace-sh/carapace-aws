package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_createWorkgroupCmd = &cobra.Command{
	Use:   "create-workgroup",
	Short: "Creates an workgroup in Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_createWorkgroupCmd).Standalone()

	redshiftServerless_createWorkgroupCmd.Flags().String("base-capacity", "", "The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).")
	redshiftServerless_createWorkgroupCmd.Flags().String("config-parameters", "", "An array of parameters to set for advanced control over a database.")
	redshiftServerless_createWorkgroupCmd.Flags().Bool("enhanced-vpc-routing", false, "The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.")
	redshiftServerless_createWorkgroupCmd.Flags().String("ip-address-type", "", "The IP address type that the workgroup supports.")
	redshiftServerless_createWorkgroupCmd.Flags().String("max-capacity", "", "The maximum data-warehouse capacity Amazon Redshift Serverless uses to serve queries.")
	redshiftServerless_createWorkgroupCmd.Flags().String("namespace-name", "", "The name of the namespace to associate with the workgroup.")
	redshiftServerless_createWorkgroupCmd.Flags().Bool("no-enhanced-vpc-routing", false, "The value that specifies whether to turn on enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC instead of over the internet.")
	redshiftServerless_createWorkgroupCmd.Flags().Bool("no-publicly-accessible", false, "A value that specifies whether the workgroup can be accessed from a public network.")
	redshiftServerless_createWorkgroupCmd.Flags().String("port", "", "The custom port to use when connecting to a workgroup.")
	redshiftServerless_createWorkgroupCmd.Flags().String("price-performance-target", "", "An object that represents the price performance target settings for the workgroup.")
	redshiftServerless_createWorkgroupCmd.Flags().Bool("publicly-accessible", false, "A value that specifies whether the workgroup can be accessed from a public network.")
	redshiftServerless_createWorkgroupCmd.Flags().String("security-group-ids", "", "An array of security group IDs to associate with the workgroup.")
	redshiftServerless_createWorkgroupCmd.Flags().String("subnet-ids", "", "An array of VPC subnet IDs to associate with the workgroup.")
	redshiftServerless_createWorkgroupCmd.Flags().String("tags", "", "A array of tag instances.")
	redshiftServerless_createWorkgroupCmd.Flags().String("track-name", "", "An optional parameter for the name of the track for the workgroup.")
	redshiftServerless_createWorkgroupCmd.Flags().String("workgroup-name", "", "The name of the created workgroup.")
	redshiftServerless_createWorkgroupCmd.MarkFlagRequired("namespace-name")
	redshiftServerless_createWorkgroupCmd.Flag("no-enhanced-vpc-routing").Hidden = true
	redshiftServerless_createWorkgroupCmd.Flag("no-publicly-accessible").Hidden = true
	redshiftServerless_createWorkgroupCmd.MarkFlagRequired("workgroup-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_createWorkgroupCmd)
}
