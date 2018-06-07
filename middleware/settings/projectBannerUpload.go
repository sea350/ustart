package settings

import (
	"fmt"
	"net/http"

	get "github.com/sea350/ustart_go/get/project"
	post "github.com/sea350/ustart_go/post/project"
)

//ProjectBannerUpload ... pushes a new banner image into ES
func ProjectBannerUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	clientFile, header, err := r.FormFile("raw-banner")
	if err != nil {
		fmt.Println("err: middleware/settings/projectBannerUpload line 14\n", err)
	}
	blob := r.FormValue("banner-data")

	//Get projectID
	proj, err := get.ProjectByID(eclient, r.FormValue("projectID"))
	if err != nil {
		fmt.Println("err: middleware/settings/projectbannerupload line 21\n", err)
	}

	buffer := make([]byte, 512)
	_, _ = clientFile.Read(buffer)
	defer clientFile.Close()
	if http.DetectContentType(buffer)[0:5] == "image" || header.Size == 0 {
		//Update the project banner
		err = post.UpdateProject(eclient, r.FormValue("projectID"), "Banner", blob)
		if err != nil {
			fmt.Println("err: middleware/settings/projectbannerupload line 32\n", err)
		}
	} else {
		fmt.Println("err: middleware/settings/projectBannerUpload invalid file upload")
	}

	http.Redirect(w, r, "/Projects/"+proj.URLName, http.StatusFound)
	return
}
