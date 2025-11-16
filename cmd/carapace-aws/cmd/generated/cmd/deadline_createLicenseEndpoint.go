package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createLicenseEndpointCmd = &cobra.Command{
	Use:   "create-license-endpoint",
	Short: "Creates a license endpoint to integrate your various licensed software used for rendering on Deadline Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createLicenseEndpointCmd).Standalone()

	deadline_createLicenseEndpointCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_createLicenseEndpointCmd.Flags().String("security-group-ids", "", "The security group IDs.")
	deadline_createLicenseEndpointCmd.Flags().String("subnet-ids", "", "The subnet IDs.")
	deadline_createLicenseEndpointCmd.Flags().String("tags", "", "Each tag consists of a tag key and a tag value.")
	deadline_createLicenseEndpointCmd.Flags().String("vpc-id", "", "The VPC (virtual private cloud) ID to use with the license endpoint.")
	deadline_createLicenseEndpointCmd.MarkFlagRequired("security-group-ids")
	deadline_createLicenseEndpointCmd.MarkFlagRequired("subnet-ids")
	deadline_createLicenseEndpointCmd.MarkFlagRequired("vpc-id")
	deadlineCmd.AddCommand(deadline_createLicenseEndpointCmd)
}
