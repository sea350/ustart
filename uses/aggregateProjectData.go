package uses

import (
	"fmt"

	getProject "github.com/sea350/ustart_go/get/project"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AggregateProjectData ...
//Adds a new widget to the UserWidgets array
func AggregateProjectData(eclient *elastic.Client, url string) (types.ProjectAggregate, error) {
	var projectData types.ProjectAggregate

	data, err := getProject.ProjectByURL(eclient, url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error aggregateProjectData.go 17")
	}
	projectData.ProjectData = data

	id, err := getProject.ProjectIDByURL(eclient, url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error aggregateProjectData.go 24")
	}
	projectData.DocID = id

	//Remember to load widgets seperately
	//Remember to load wall posts seperately

	for _, member := range data.Members {
		id := member.MemberID
		mem, err := ConvertUserToFloatingHead(eclient, id)
		if err != nil {
			panic(err)
		}
		projectData.MemberData = append(projectData.MemberData, mem)
	}

	return projectData, err
}
