package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_associateProfileCmd = &cobra.Command{
	Use:   "associate-profile",
	Short: "Associates a Route 53 Profiles profile with a VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_associateProfileCmd).Standalone()

	route53profiles_associateProfileCmd.Flags().String("name", "", "A name for the association.")
	route53profiles_associateProfileCmd.Flags().String("profile-id", "", "ID of the Profile.")
	route53profiles_associateProfileCmd.Flags().String("resource-id", "", "The ID of the VPC.")
	route53profiles_associateProfileCmd.Flags().String("tags", "", "A list of the tag keys and values that you want to identify the Profile association.")
	route53profiles_associateProfileCmd.MarkFlagRequired("name")
	route53profiles_associateProfileCmd.MarkFlagRequired("profile-id")
	route53profiles_associateProfileCmd.MarkFlagRequired("resource-id")
	route53profilesCmd.AddCommand(route53profiles_associateProfileCmd)
}
