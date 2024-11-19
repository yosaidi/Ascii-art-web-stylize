package ascii

import (
	"html/template"
	"net/http"
)

// TemplateData for rendering the result and error information
type TemplateData struct {
	Result       string
	ErrorCode    int
	ErrorMessage string
	ErrorDetails string
}

// Global variable to hold the template
var Temp *template.Template

// Render error page with dynamic error message
func RenderErrorPage(w http.ResponseWriter, errorCode int, errorMessage, errorDetails string) {
	// Prepare the error data to be passed to the template
	errorData := TemplateData{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
		ErrorDetails: errorDetails,
	}

	w.WriteHeader(errorCode)
	Temp.ExecuteTemplate(w,"error.html", errorData)
}

// Handle main page
func MainPage(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == "POST" {
		AsciiProcess(w, r)
		return
	} else {

		err = Temp.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			RenderErrorPage(w, http.StatusInternalServerError, "Internal Server Error", "There was an issue rendering the page.")
		}
	}
}

// Process ASCII form submission
func AsciiProcess(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			RenderErrorPage(w, http.StatusInternalServerError, "Internal Server Error", "Something went wrong. Please try again later.")
		}
	}()

	if r.Method == "POST" {
		input := r.FormValue("text")
		templ := r.FormValue("template")
		if input == "" || templ == "" {
			RenderErrorPage(w, http.StatusBadRequest, "Bad Request", "Please fill in all fields.")
			return
		}
		if templ != "standard" && templ != "shadow" && templ != "thinkertoy" {
			RenderErrorPage(w, http.StatusBadRequest, "Bad Request", "Invalid template selected.")
			return
		}
		helperInput := []rune(input)
		if !AreStringValid(helperInput) {
			RenderErrorPage(w, http.StatusBadRequest, "Bad Request", "Input contains invalid characters.")
			return
		}
		if len(helperInput) > 600 {
			RenderErrorPage(w, http.StatusBadRequest, "Bad Request", "Input exceeds maximum length of 600 characters.")
			return
		}
		result, err := Transform(input, templ)
		if err != nil {
			RenderErrorPage(w, http.StatusInternalServerError, "Internal Server Error", "An error occurred while processing the input.")
			return
		}
		data := TemplateData{Result: result}
		err = Temp.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			RenderErrorPage(w, http.StatusInternalServerError, "Internal Server Error", "There was an issue rendering the result.")
			return
		}
	} else {
		RenderErrorPage(w, http.StatusMethodNotAllowed, "Method Not Allowed", "Only POST method is allowed for this request.")
	}
}
