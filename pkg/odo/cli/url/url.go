package url

import (
	"fmt"

	ktemplates "k8s.io/kubernetes/pkg/kubectl/util/templates"

	odoutil "github.com/openshift/odo/pkg/odo/util"
	"github.com/spf13/cobra"
)

// RecommendedCommandName is the recommended url command name
const RecommendedCommandName = "url"

var (
	urlShortDesc = `Expose component to the outside world`
	urlLongDesc  = ktemplates.LongDesc(`Expose component to the outside world.
		
		The URLs that are generated using this command, can be used to access the deployed components from outside the cluster.`)
)

// NewCmdURL returns the top-level url command
func NewCmdURL(name, fullName string) *cobra.Command {
	urlCreateCmd := NewCmdURLCreate(createRecommendedCommandName, odoutil.GetFullName(fullName, createRecommendedCommandName))
	urlDeleteCmd := NewCmdURLDelete(deleteRecommendedCommandName, odoutil.GetFullName(fullName, deleteRecommendedCommandName))
	urlListCmd := NewCmdURLList(listRecommendedCommandName, odoutil.GetFullName(fullName, listRecommendedCommandName))
	urlDescribeCmd := NewCmdURLDescribe(describeRecommendedCommandName, odoutil.GetFullName(fullName, describeRecommendedCommandName))
	urlCmd := &cobra.Command{
		Use:   name,
		Short: urlShortDesc,
		Long:  urlLongDesc,
		Example: fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s",
			urlCreateCmd.Example,
			urlDeleteCmd.Example,
			urlListCmd.Example,
			urlDescribeCmd.Example),
	}

	// Add a defined annotation in order to appear in the help menu
	urlCmd.Annotations = map[string]string{"command": "main"}
	urlCmd.SetUsageTemplate(odoutil.CmdUsageTemplate)
	urlCmd.AddCommand(urlCreateCmd, urlDeleteCmd, urlListCmd, urlDescribeCmd)

	return urlCmd
}
