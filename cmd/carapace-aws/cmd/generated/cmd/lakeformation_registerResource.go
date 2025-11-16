package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_registerResourceCmd = &cobra.Command{
	Use:   "register-resource",
	Short: "Registers the resource as managed by the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_registerResourceCmd).Standalone()

	lakeformation_registerResourceCmd.Flags().String("hybrid-access-enabled", "", "Specifies whether the data access of tables pointing to the location can be managed by both Lake Formation permissions as well as Amazon S3 bucket policies.")
	lakeformation_registerResourceCmd.Flags().Bool("no-with-privileged-access", false, "Grants the calling principal the permissions to perform all supported Lake Formation operations on the registered data location.")
	lakeformation_registerResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to register.")
	lakeformation_registerResourceCmd.Flags().String("role-arn", "", "The identifier for the role that registers the resource.")
	lakeformation_registerResourceCmd.Flags().String("use-service-linked-role", "", "Designates an Identity and Access Management (IAM) service-linked role by registering this role with the Data Catalog.")
	lakeformation_registerResourceCmd.Flags().String("with-federation", "", "Whether or not the resource is a federated resource.")
	lakeformation_registerResourceCmd.Flags().Bool("with-privileged-access", false, "Grants the calling principal the permissions to perform all supported Lake Formation operations on the registered data location.")
	lakeformation_registerResourceCmd.Flag("no-with-privileged-access").Hidden = true
	lakeformation_registerResourceCmd.MarkFlagRequired("resource-arn")
	lakeformationCmd.AddCommand(lakeformation_registerResourceCmd)
}
