package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_createRecipeJobCmd = &cobra.Command{
	Use:   "create-recipe-job",
	Short: "Creates a new job to transform input data, using steps defined in an existing Glue DataBrew recipe",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_createRecipeJobCmd).Standalone()

	databrew_createRecipeJobCmd.Flags().String("data-catalog-outputs", "", "One or more artifacts that represent the Glue Data Catalog output from running the job.")
	databrew_createRecipeJobCmd.Flags().String("database-outputs", "", "Represents a list of JDBC database output objects which defines the output destination for a DataBrew recipe job to write to.")
	databrew_createRecipeJobCmd.Flags().String("dataset-name", "", "The name of the dataset that this job processes.")
	databrew_createRecipeJobCmd.Flags().String("encryption-key-arn", "", "The Amazon Resource Name (ARN) of an encryption key that is used to protect the job.")
	databrew_createRecipeJobCmd.Flags().String("encryption-mode", "", "The encryption mode for the job, which can be one of the following:")
	databrew_createRecipeJobCmd.Flags().String("log-subscription", "", "Enables or disables Amazon CloudWatch logging for the job.")
	databrew_createRecipeJobCmd.Flags().String("max-capacity", "", "The maximum number of nodes that DataBrew can consume when the job processes data.")
	databrew_createRecipeJobCmd.Flags().String("max-retries", "", "The maximum number of times to retry the job after a job run fails.")
	databrew_createRecipeJobCmd.Flags().String("name", "", "A unique name for the job.")
	databrew_createRecipeJobCmd.Flags().String("outputs", "", "One or more artifacts that represent the output from running the job.")
	databrew_createRecipeJobCmd.Flags().String("project-name", "", "Either the name of an existing project, or a combination of a recipe and a dataset to associate with the recipe.")
	databrew_createRecipeJobCmd.Flags().String("recipe-reference", "", "")
	databrew_createRecipeJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role to be assumed when DataBrew runs the job.")
	databrew_createRecipeJobCmd.Flags().String("tags", "", "Metadata tags to apply to this job.")
	databrew_createRecipeJobCmd.Flags().String("timeout", "", "The job's timeout in minutes.")
	databrew_createRecipeJobCmd.MarkFlagRequired("name")
	databrew_createRecipeJobCmd.MarkFlagRequired("role-arn")
	databrewCmd.AddCommand(databrew_createRecipeJobCmd)
}
