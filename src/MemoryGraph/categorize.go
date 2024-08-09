// src/MemoryGraph/categorize.go
package memorygraph

import (
	"strings"
	"GoGo/src/Direct"
	"net/http"
	"github.com/gin-gonic/gin"
)
	 //step 1: is still subject
	 //	if no add new subject node to the graph
	 // 	- then add a new Tree to the subject node
	 // 	- then add new Chat node to the tree
	 //	if yes then check if the prompt is still on the same topic in the subject
	 // 	- if no then add a new Tree node to the subject
	 // 	- then add a new Chat node to the tree
	 // 	- if yes then add a new Chat node to the tree
	 //step 2: get relivent data for the response
	 // 	- then get relivent data for the response
	 // 	- Return the response

func NewPrompt(c *gin.Context){
	 //get the prompt from the request
	 var request struct {
		Prompt string `json:"prompt"`
	 }
	 if err := c.BindJSON(&request); err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	 }

}

// handleSubject processes the prompt and assigns or creates a subject using AI.
func HandleSubject(prompt string) (string, error) {
	subjects := graph.GetSubjects()
	aiPrompt := "Subjects: " + formatSubjects(subjects) +
		" Please assign the following prompt to a subject, if no subject fits then please name a new one." +
		" use the formate for new subject use 'New: {}' and for add to current use 'AddTO: {}'" +
		"\nPrompt: " + prompt

	// Call the oneshot function from the direct package
	response, err := Direct.Oneshot(aiPrompt)
	if err != nil {
		return "", err
	}
	//if response is a new subject then add it to the graph
	if strings.Contains(response, "New: ") {
		addNewSubject(response)
		// get the subject from the response
		subject := strings.Split(response, "New: ")[1]
		// get the subject node from the graph
		SubjectNode := graph.FindNodeBySubject(subject)
		// add a new tree node to the subject
		SubjectNode.AddTree(makeTreeTitle(subject, prompt))
	}
	//if response is to add to a current subject then add it to the graph
	if strings.Contains(response, "AddTo: ") {
		// get the subject from the response
		subject := strings.Split(response, "AddTo: ")[1]
		// get the subject node from the graph
		SubjectNode := graph.FindNodeBySubject(subject)
		//
	}
	//else error


	return response, nil
}

func makeTreeTitle(subject string, prompt string) string {
	//ai oneshot to make the title
	aiPrompt := "Please make a title (use the formate 'Title: {}') for the tree with the subject: " + subject + " and the prompt: " + prompt
	response, err := Direct.Oneshot(aiPrompt)
	if err != nil {
		return ""
	}
	return strings.Split(response, "Title: ")[1]
}

// formatSubjects formats the subjects into a readable string.
func formatSubjects(subjects []string) string {
	return "[" + strings.Join(subjects, ", ") + "]"
}

// subjectExists checks if a subject already exists in the current subjects list.
func SubjectExists(subject string, subjects []string) bool {
	for _, s := range subjects {
		if s == subject {
			return true
		}
	}
	return false
}


// AddSubject adds a new subject to the graph's subjects list.
func addNewSubject(subject string) {
	//parse the subject from the response
	subject = strings.Split(subject, "New: ")[1]
	graph.AddNodeWithUniqueKey(subject)
}