package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-spec/pkg/command"
	"github.com/neurosnap/sentences/english"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var rootCmd = &cobra.Command{
	Use:   "carapace-spec-botocore",
	Short: "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		parse(args[0])
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
}

func CamelCaseToDash(s string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	s = re.ReplaceAllString(s, "${1}-${2}")

	re = regexp.MustCompile("([A-Z])([A-Z])([a-z0-9])")
	s = re.ReplaceAllString(s, "${1}-${2}${3}")

	s = strings.ToLower(s)

	quickfix := map[string]string{
		"application-i-ds":                         "application-ids",
		"asset-sh-a256":                            "asset-sha256",
		"autonomous-data-storage-size-in-t-bs":     "autonomous-data-storage-size-in-tbs",
		"callback-ur-ls":                           "callback-urls",
		"data-storage-size-in-t-bs":                "data-storage-size-in-tbs",
		"db-node-storage-size-in-g-bs":             "db-node-storage-size-in-gbs",
		"e-c2-family-filter":                       "ec2-family-filter",
		"e-c2-inbound-permissions":                 "ec2-inbound-permissions",
		"e-c2-instance-id":                         "ec2-instance-id",
		"e-c2-instance-type":                       "ec2-instance-type",
		"e-c2-security-group-id":                   "ec2-security-group-id",
		"e-c2-security-group-name":                 "ec2-security-group-name",
		"e-c2-security-group-owner-id":             "ec2-security-group-owner-id",
		"edn-s0-client-subnet-ip":                  "edn-s0-client-subnet-mask",
		"edns0-client-subnet-ip":                   "edns0-client-subnet-mask",
		"launch-configuration-template-i-ds":       "launch-configuration-template-ids",
		"logout-ur-ls":                             "logout-urls",
		"memory-per-oracle-compute-unit-in-g-bs":   "memory-per-oracle-compute-unit-in-gbs",
		"memory-size-in-g-bs":                      "memory-size-in-gbs",
		"no-retrieve-a-zs":                         "no-retrieve-azs",
		"notification-ar-ns":                       "notification-arns",
		"open-id-connect-provider-ar-ns":           "open-id-connect-provider-arns",
		"pool-ar-ns":                               "pool-arns",
		"provider-ar-ns":                           "provider-arns",
		"recovery-instance-i-ds":                   "recovery-instance-ids",
		"replication-servers-security-groups-i-ds": "replication-servers-security-groups-ids",
		"retrieve-a-zs":                            "retrieve-azs",
		"saml-provider-ar-ns":                      "saml-provider-arns",
		"source-server-i-d":                        "source-server-ids",
		"tape-ar-ns":                               "tape-arns",
		"target-group-ar-ns":                       "target-group-arns",
		"vtl-device-ar-ns":                         "vtl-device-arns",
	}
	if replacement, ok := quickfix[s]; ok {
		return replacement
	}
	return s
}

type Service struct {
	Version  string `json:"version,omitempty"`
	Metadata struct {
		APIVersion          string   `json:"apiVersion,omitempty"`
		EndpointPrefix      string   `json:"endpointPrefix,omitempty"`
		Protocol            string   `json:"protocol,omitempty"`
		Protocols           []string `json:"protocols,omitempty"`
		ServiceAbbreviation string   `json:"serviceAbbreviation,omitempty"`
		ServiceFullName     string   `json:"serviceFullName,omitempty"`
		ServiceID           string   `json:"serviceId,omitempty"`
		SignatureVersion    string   `json:"signatureVersion,omitempty"`
		UID                 string   `json:"uid,omitempty"`
		XMLNamespace        string   `json:"xmlNamespace,omitempty"`
		Auth                []string `json:"auth,omitempty"`
	} `json:"metadata"`
	Operations    map[string]Operation
	Shapes        map[string]Shape `json:"shapes"`
	Documentation string           `json:"documentation,omitempty"`
}

type Shape struct {
	Type     string            `json:"type,omitempty"`
	Enum     []string          `json:"enum,omitempty"`
	Member   Member            `json:"member"`
	Members  map[string]Member `json:"members"`
	Required []string          `json:"required,omitempty"`
}

type Member struct {
	Shape         string `json:"shape,omitempty"`
	Documentation string `json:"documentation,omitempty"`
	LocationName  string `json:"locationName,omitempty"`
}

type Operation struct {
	Name string `json:"name"`
	HTTP struct {
		Method     string `json:"method"`
		RequestURI string `json:"requestUri"`
	} `json:"http"`
	Input struct {
		Shape string `json:"shape"`
	} `json:"input"`
	Output struct {
		Shape string `json:"shape"`
	} `json:"output"`
	Documentation string `json:"documentation"`
}

func parse(path string) {
	services, err := os.ReadDir(path)
	if err != nil {
		panic(err.Error())
	}

	cmd := command.Command{
		Name:        "aws",
		Description: "The AWS Command Line Interface is a unified tool to manage your AWS services.",
	}

	for _, serviceDir := range services {
		if !serviceDir.IsDir() {
			continue
		}
		path := filepath.Join(os.Args[1], serviceDir.Name())
		versions, err := os.ReadDir(path)
		if err != nil {
			panic(err.Error())
		}
		cmd.Commands = append(cmd.Commands,
			parseService(serviceDir.Name(), filepath.Join(path, versions[len(versions)-1].Name())),
		)
	}

	m, err := yaml.Marshal(cmd)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(m))

}

type Paginator struct {
	LimitKey *string `json:"limit_key,omitempty"`
}

func parsePaginators(folder string) (map[string]Paginator, error) {
	content, err := os.ReadFile(filepath.Join(folder, "paginators-1.json"))
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]Paginator), nil
		}
		return nil, err
	}

	var paginators struct {
		Pagination map[string]Paginator `json:"pagination"`
	}
	if err := json.Unmarshal(content, &paginators); err != nil {
		return nil, err
	}
	return paginators.Pagination, nil
}

func parseService(name, folder string) command.Command {
	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err.Error())
	}

	content, err := os.ReadFile(filepath.Join(folder, "service-2.json"))
	if err != nil {
		panic(err.Error())
	}

	var service Service
	if err := json.Unmarshal(content, &service); err != nil {
		panic(err.Error())
	}

	paginators, err := parsePaginators(folder)
	if err != nil {
		panic(err.Error())
	}

	cmd := command.Command{Name: name}
	cmd.Description = service.Metadata.ServiceFullName
	cmd.Documentation.Command, _ = htmltomarkdown.ConvertString(service.Documentation)

	for name, operation := range service.Operations {
		opdoc, _ := htmltomarkdown.ConvertString(operation.Documentation)
		subCmd := command.Command{
			Name: CamelCaseToDash(name),
		}
		if tokens := tokenizer.Tokenize(opdoc); len(tokens) > 0 {
			subCmd.Description = strings.Split(tokens[0].Text, "\n")[0]
		}
		subCmd.Documentation.Command = opdoc
		subCmd.Documentation.Flag = make(map[string]string)

		// custom flags added by awscli
		subCmd.AddFlag(command.Flag{Longhand: "--cli-input-json", Usage: "Read arguments from the JSON string provided.", Value: true})
		subCmd.AddFlag(command.Flag{Longhand: "--cli-input-yaml", Usage: "Read arguments from the YAML string provided.", Value: true})
		subCmd.AddFlag(command.Flag{Longhand: "--generate-cli-skeleton", Usage: "Prints a JSON skeleton to standard output without sending an API request."})

		subCmd.Documentation.Flag["cli-input-json"] = `Reads arguments from the JSON string provided.
The JSON string follows the  format  provided  by --generate-cli-skeleton.
If other arguments are provided on the command line,  those  values  will override the JSON-provided values.
It is not possible to pass arbitrary binary values using a JSON-provided value as the string will be taken literally.
This may  not  be  specified  along with --cli-input-yaml.`
		subCmd.Documentation.Flag["cli-input-yaml"] = `Reads arguments from the JSON string provided.
The JSON string follows the  format  provided  by --generate-cli-skeleton.
If other arguments are provided on the command line,  those  values  will override the JSON-provided values.
It is not possible to pass arbitrary binary values using a JSON-provided value as the string will be taken literally.
This may  not  be  specified  along with --cli-input-yaml.`
		subCmd.Documentation.Flag["generate-cli-skeleton"] = `If other arguments are provided on the command line,  those  values  will override the JSON-provided values. It is not possible to pass arbitrary binary values using a JSON-provided value as the string will be taken literally.
This may  not  be  specified  along with --cli-input-yaml.`

		if paginator, ok := paginators[name]; ok {
			subCmd.AddFlag(command.Flag{Longhand: "--max-items", Usage: "The  total number of items to return in the command's output.", Value: true})
			subCmd.AddFlag(command.Flag{Longhand: "--starting-token", Usage: "A token to specify where to start paginating.", Value: true})

			subCmd.Documentation.Flag["max-items"] = `The total number of items to return in the command's output.
If the total number of items available is more than the value specified, a NextToken is provided in the command's output.
To resume pagination, provide the NextToken value in the starting-token argument of a subsequent  command.
Do not use the NextToken response element directly outside of the AWS CLI.

For usage examples, see Pagination in the AWS Command Line Interface User Guide.`
			subCmd.Documentation.Flag["starting-token"] = `A token to specify where to start paginating.
This is the NextToken from a previously truncated response.

For usage examples, see Pagination in the AWS Command Line Interface User Guide.`

			if paginator.LimitKey != nil {
				subCmd.AddFlag(command.Flag{Longhand: "--page-size", Usage: "The size of each page to get in the AWS service call.", Value: true})

				subCmd.Documentation.Flag["page-size"] = `The size of each page to get in the AWS service call.
This does not affect the number of items returned in the command's output.
Setting a smaller page size results in more calls to the AWS service, retrieving fewer items in each call.
This can help prevent the AWS service calls from timing out.

For usage examples, see Pagination in the AWS Command Line Interface User Guide.`
			}
		}

		if shape, ok := service.Shapes[operation.Input.Shape]; ok {
			switch shape.Type {
			case "structure":
				// TODO shape.Member
				for name, member := range shape.Members {
					required := slices.Contains(shape.Required, name)
					memberdoc, _ := htmltomarkdown.ConvertString(member.Documentation)

					flagName := renameFlag(cmd.Name, subCmd.Name, CamelCaseToDash(name))
					subCmd.Documentation.Flag[flagName] = memberdoc
					if tokens := tokenizer.Tokenize(memberdoc); len(tokens) > 0 {
						memberdoc = strings.Split(tokens[0].Text, "\n")[0]
					}

					boolFlag := isBoolean(member.Shape)
					if !boolFlag {
						if shape, ok := service.Shapes[member.Shape]; ok && isBoolean(shape.Type) {
							boolFlag = true
						}
					}

					subCmd.AddFlag(command.Flag{
						Longhand: "--" + flagName,
						Usage:    memberdoc,
						Value:    !boolFlag,
						Required: required,
					})
					if boolFlag {
						subCmd.AddFlag(command.Flag{
							Longhand: "--no-" + flagName,
							Usage:    memberdoc,
							Required: required,
						})
					}

					if memberShape, ok := service.Shapes[member.Shape]; ok &&
						strings.ToLower(memberShape.Type) == "string" && len(memberShape.Enum) != 0 {
						if subCmd.Completion.Flag == nil {
							subCmd.Completion.Flag = map[string][]string{}
						}
						subCmd.Completion.Flag[CamelCaseToDash(name)] = memberShape.Enum
					}
				}
			default:
				// TODO others
			}
		}
		cmd.Commands = append(cmd.Commands, subCmd)
	}
	return cmd
}

func isBoolean(s string) bool {
	return (strings.HasSuffix(strings.ToLower(s), "boolean") && !strings.HasPrefix(strings.ToLower(s), "map")) || strings.HasSuffix(s, "AttributeBooleanValue")
}
