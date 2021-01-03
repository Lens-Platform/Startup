package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"go.uber.org/zap"
)

// CaseInsensitiveContains checks if a string witholds a substring irrespective of case
func CaseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

// UpdateMakefile updates the sevices makefile
func UpdateMakefile(serviceName string) error {
	src := "./generator/templates/additive-make.txt"
	dest := fmt.Sprintf("./%s/Makefile", serviceName)

	// read date present in the current make file
	oldData, err := ioutil.ReadFile(dest)
	if err != nil {
		return LogAndReturnCmdErrorIfExist(err)
	}

	// read the file of interest
	newData, err := ioutil.ReadFile(src)
	if err != nil {
		return LogAndReturnCmdErrorIfExist(err)
	}

	oldData = append(oldData, newData...)

	err = ioutil.WriteFile(dest, oldData, 0644)
	if err != nil {
		return LogAndReturnCmdErrorIfExist(err)
	}

	return nil
}

// RemoveAllReferencesFromFile removes all references of a search string in a file located at a specific location
func RemoveAllReferencesFromFile(path string, search string, to string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if CaseInsensitiveContains(line, search) {
			re := regexp.MustCompile(`(?i)`+search)
			lines[i] = re.ReplaceAllString(line, to)
		}

		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(path, []byte(output), 0644)
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func WalkAndUpdate(serviceName string){
	searchDir := "./" + serviceName

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	var updateStr =  blackspaceRepo+"/src/services/"+serviceName
	glog.Info("removing all references of authentication_handler_service from files")
	for _, file := range fileList {

		if filepath.Ext(file) != ".go" && filepath.Ext(file) != ".png" && filepath.Ext(file) != ".json" {
				glog.Info("file", zap.String("ext", file))
				// remove all references of authentication_handler_service in file
				RemoveAllReferencesFromFile(file, "authentication_handler_service", serviceName)
			} else {
				RemoveAllReferencesFromFile(file, "github.com/stefanprodan/authentication_handler_service",updateStr)
				RemoveAllReferencesFromFile(file, "stefanprodan/authentication_handler_service", updateStr)
				RemoveAllReferencesFromFile(file, "authentication_handler_service", serviceName)
			}
			// remove all references of stefanprodan
			// var stefanSearchStr = []string{ "stefanprodan", "Stefan Prodan"}
			RemoveAllReferencesFromFile(file, "github.com/stefanprodan/"+serviceName,updateStr)
			RemoveAllReferencesFromFile(file, "stefanprodan/"+serviceName,blackspaceRepo+"/"+serviceName)
			RemoveAllReferencesFromFile(file, "Stefan Prodan", "Yoan Yomba, Vic Amupitan, Samira, Cameron Burford")
			RemoveAllReferencesFromFile(file, "stefanprodan.github.io", blackspaceRepo)
			RemoveAllReferencesFromFile(file, "maintainer=\"stefanprodan\"", "maintainer=" + "\"" + blackspaceRepo + "\"")
			RemoveAllReferencesFromFile(file, "DOCKER_REPOSITORY:=stefanprodan", "DOCKER_REPOSITORY:="+blackspaceRepo)
			RemoveAllReferencesFromFile(file, "name: stefanprodan", "name: Yoan Yomba")
			RemoveAllReferencesFromFile(file, "stefanprodan@users.noreply.github.com", serviceName+"@users.noreply.github.com")
	}
	glog.Info("successfully removed all references of authentication_handler_service from files")
}

// setupCommonTemplates sets up template files for generated service
func setupCommonTemplates(serviceName string) error {
	// iterate over file map and copy over templates from template folder to actual service folder
	for source, target := range fileMap {
		var targetPath = fmt.Sprintf("./%s/%s", serviceName, target)

		var res = exec.Command("cp", "-rf", source, targetPath).Run()
		if res != nil {
			glog.Info("source directory", zap.String("source", source))
			glog.Info("target directory", zap.String("target", targetPath))
			return LogAndReturnCmdErrorIfExist(res)
		}
	}
	return nil
}

// renameRepositoryAndSetupService renames a repository and skaffolds service specific directories and files
func renameRepositoryAndSetupService(serviceName string) error {
	glog.Info("renaming cloned repository", zap.String("repo", serviceName))
	var res = exec.Command("mv", "-f", repoName, serviceName).Run()
	if res != nil {
		return LogAndReturnCmdErrorIfExist(res)
	}

	if res := RemoveGitRepo(serviceName); res != nil {
		return LogAndReturnCmdErrorIfExist(res)
	}
	glog.Info("setting up service", zap.String("service", serviceName))

	// create pkg specific directories
	res = setupCommonTemplates(serviceName)
	if res != nil {
		return LogAndReturnCmdErrorIfExist(res)
	}

	glog.Info("finished up service", zap.String("service", serviceName))
	return nil
}

// Clone Template clones a template microservice
func CloneTemplate() error {
	glog.Info("cloning template repository", zap.String("repo", templateRepoUrl))
	var res = exec.Command("git", "clone", templateRepoUrl).Run()
	return LogAndReturnCmdErrorIfExist(res)
}

// RemoveGitRepo removes the .git repository
func RemoveGitRepo(serviceName string) error {
	glog.Info("removing .git repository", zap.String("folder", serviceName))
	var res = exec.Command("rm", "-rf", "./"+serviceName+"/.git").Run()
	return LogAndReturnCmdErrorIfExist(res)
}

// LogAndReturnCmdErrorIfExist logs and returns an error if it exists
func LogAndReturnCmdErrorIfExist(res error) error {
	if res != nil && res.Error() != EMPTY {
		glog.Error(res.Error())
	}
	return res
}
