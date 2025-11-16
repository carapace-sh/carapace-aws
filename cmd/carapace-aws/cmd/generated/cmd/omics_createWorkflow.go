package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_createWorkflowCmd = &cobra.Command{
	Use:   "create-workflow",
	Short: "Creates a private workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_createWorkflowCmd).Standalone()

	omics_createWorkflowCmd.Flags().String("accelerators", "", "The computational accelerator specified to run the workflow.")
	omics_createWorkflowCmd.Flags().String("container-registry-map", "", "(Optional) Use a container registry map to specify mappings between the ECR private repository and one or more upstream registries.")
	omics_createWorkflowCmd.Flags().String("container-registry-map-uri", "", "(Optional) URI of the S3 location for the registry mapping file.")
	omics_createWorkflowCmd.Flags().String("definition-repository", "", "The repository information for the workflow definition.")
	omics_createWorkflowCmd.Flags().String("definition-uri", "", "The S3 URI of a definition for the workflow.")
	omics_createWorkflowCmd.Flags().String("definition-zip", "", "A ZIP archive containing the main workflow definition file and dependencies that it imports for the workflow.")
	omics_createWorkflowCmd.Flags().String("description", "", "A description for the workflow.")
	omics_createWorkflowCmd.Flags().String("engine", "", "The workflow engine for the workflow.")
	omics_createWorkflowCmd.Flags().String("main", "", "The path of the main definition file for the workflow.")
	omics_createWorkflowCmd.Flags().String("name", "", "Name (optional but highly recommended) for the workflow to locate relevant information in the CloudWatch logs and Amazon Web Services HealthOmics console.")
	omics_createWorkflowCmd.Flags().String("parameter-template", "", "A parameter template for the workflow.")
	omics_createWorkflowCmd.Flags().String("parameter-template-path", "", "The path to the workflow parameter template JSON file within the repository.")
	omics_createWorkflowCmd.Flags().String("readme-markdown", "", "The markdown content for the workflow's README file.")
	omics_createWorkflowCmd.Flags().String("readme-path", "", "The path to the workflow README markdown file within the repository.")
	omics_createWorkflowCmd.Flags().String("readme-uri", "", "The S3 URI of the README file for the workflow.")
	omics_createWorkflowCmd.Flags().String("request-id", "", "An idempotency token to ensure that duplicate workflows are not created when Amazon Web Services HealthOmics submits retry requests.")
	omics_createWorkflowCmd.Flags().String("storage-capacity", "", "The default static storage capacity (in gibibytes) for runs that use this workflow or workflow version.")
	omics_createWorkflowCmd.Flags().String("storage-type", "", "The default storage type for runs that use this workflow.")
	omics_createWorkflowCmd.Flags().String("tags", "", "Tags for the workflow.")
	omics_createWorkflowCmd.Flags().String("workflow-bucket-owner-id", "", "The Amazon Web Services account ID of the expected owner of the S3 bucket that contains the workflow definition.")
	omics_createWorkflowCmd.MarkFlagRequired("request-id")
	omicsCmd.AddCommand(omics_createWorkflowCmd)
}
