package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_updateRecipeJobCmd = &cobra.Command{
	Use:   "update-recipe-job",
	Short: "Modifies the definition of an existing DataBrew recipe job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_updateRecipeJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_updateRecipeJobCmd).Standalone()

		databrew_updateRecipeJobCmd.Flags().String("data-catalog-outputs", "", "One or more artifacts that represent the Glue Data Catalog output from running the job.")
		databrew_updateRecipeJobCmd.Flags().String("database-outputs", "", "Represents a list of JDBC database output objects which defines the output destination for a DataBrew recipe job to write into.")
		databrew_updateRecipeJobCmd.Flags().String("encryption-key-arn", "", "The Amazon Resource Name (ARN) of an encryption key that is used to protect the job.")
		databrew_updateRecipeJobCmd.Flags().String("encryption-mode", "", "The encryption mode for the job, which can be one of the following:")
		databrew_updateRecipeJobCmd.Flags().String("log-subscription", "", "Enables or disables Amazon CloudWatch logging for the job.")
		databrew_updateRecipeJobCmd.Flags().String("max-capacity", "", "The maximum number of nodes that DataBrew can consume when the job processes data.")
		databrew_updateRecipeJobCmd.Flags().String("max-retries", "", "The maximum number of times to retry the job after a job run fails.")
		databrew_updateRecipeJobCmd.Flags().String("name", "", "The name of the job to update.")
		databrew_updateRecipeJobCmd.Flags().String("outputs", "", "One or more artifacts that represent the output from running the job.")
		databrew_updateRecipeJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role to be assumed when DataBrew runs the job.")
		databrew_updateRecipeJobCmd.Flags().String("timeout", "", "The job's timeout in minutes.")
		databrew_updateRecipeJobCmd.MarkFlagRequired("name")
		databrew_updateRecipeJobCmd.MarkFlagRequired("role-arn")
	})
	databrewCmd.AddCommand(databrew_updateRecipeJobCmd)
}
