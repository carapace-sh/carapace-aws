package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_putPrincipalMappingCmd = &cobra.Command{
	Use:   "put-principal-mapping",
	Short: "Maps users to their groups so that you only need to provide the user ID when you issue the query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_putPrincipalMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_putPrincipalMappingCmd).Standalone()

		kendra_putPrincipalMappingCmd.Flags().String("data-source-id", "", "The identifier of the data source you want to map users to their groups.")
		kendra_putPrincipalMappingCmd.Flags().String("group-id", "", "The identifier of the group you want to map its users to.")
		kendra_putPrincipalMappingCmd.Flags().String("group-members", "", "The list that contains your users that belong the same group.")
		kendra_putPrincipalMappingCmd.Flags().String("index-id", "", "The identifier of the index you want to map users to their groups.")
		kendra_putPrincipalMappingCmd.Flags().String("ordering-id", "", "The timestamp identifier you specify to ensure Amazon Kendra doesn't override the latest `PUT` action with previous actions.")
		kendra_putPrincipalMappingCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that has access to the S3 file that contains your list of users that belong to a group.")
		kendra_putPrincipalMappingCmd.MarkFlagRequired("group-id")
		kendra_putPrincipalMappingCmd.MarkFlagRequired("group-members")
		kendra_putPrincipalMappingCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_putPrincipalMappingCmd)
}
